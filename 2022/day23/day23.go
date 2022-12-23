// https://adventofcode.com/2022/day/23
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x, y int
}

type Elve struct {
	proposedMove Point
	toMove       bool
}

var elves map[Point]Elve
var directions []rune = []rune{'N', 'S', 'W', 'E'}

func main() {
	elves = make(map[Point]Elve)
	parseInput()

	var round10empty, finalRound int

	for round := 1; true; round++ {
		someoneMoved := simulateRound(round - 1) // simulate a round

		if round == 10 { // save empty spaces at round 10
			round10empty = countEmpty()
		}

		if !someoneMoved {
			finalRound = round
			break
		}
	}

	fmt.Println("empty spaces after round 10 (part1):\n\t", round10empty)
	fmt.Println("first rounds where no elf move (part2):\n\t", finalRound)
}

// modifies elfes map parsing stdin input
// modifies stdin
func parseInput() {
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, v := range line {
			if v == '#' {
				elves[Point{x, y}] = Elve{Point{}, false}
			}
		}
		y++
	}
}

// returns true if at least one elf moved, false otherwise
// modifies elfes map simulating the movements of a round
func simulateRound(nRound int) (someoneMoved bool) {

	generateProposedMoves(nRound) // generate proposed moves

	elves, someoneMoved = moveToProposed() // moves elves to proposed position (if possible)

	return someoneMoved // return if someone moved
}

// modifies elfs assigning them the proposed move and the will to move
func generateProposedMoves(nRound int) {
	// generate proposed moves
	for p := range elves {

		// scan 8 positions around the elve
		someoneClose := scanAround(p)

		// noone around, do nothing
		if !someoneClose {
			elves[p] = Elve{Point{}, false}
			continue
		}

		// find direction to move
		var i int                // numbers of directions checked
		var dirFound bool        // found direction to move
		var direction int        // direction to move
		var dirModifiers []Point // modifiers to reach points of selected direction

		for i < 4 { // try every possible direction
			direction = (nRound + i) % 4 // starting direction to look (changes every round)
			dirModifiers = directionToModifiers(directions[direction])
			if checkFreeDirection(p, dirModifiers) { // check if direction is free
				dirFound = true // direction found
				break           // stop looking for direction
			}
			i++
		}

		// update elf setting his proposed move
		if dirFound { // move found
			// apply direction modifiers to current point x and y
			proposedMoveX := p.x + dirModifiers[1].x
			proposedMoveY := p.y + dirModifiers[1].y
			// assign proposed move and will to move to elf
			elves[p] = Elve{Point{proposedMoveX, proposedMoveY}, true}
		} else {
			elves[p] = Elve{Point{}, false} // no will to move
		}
	}
}

// returns true if there is an elf in the 8 positions around "p", false otherwise
func scanAround(p Point) bool {
	for y := p.y - 1; y < p.y-1+3; y++ {
		for x := p.x - 1; x < p.x-1+3; x++ {
			if y == p.y && x == p.x { // skip self
				continue
			}
			_, found := elves[Point{x, y}]
			if found { // elf found
				return true
			}
		}
	}
	return false
}

// returns the modifiers to apply to current point to reach direction "dir"
func directionToModifiers(dir rune) []Point {
	switch dir {
	case 'N':
		return []Point{{-1, -1}, {0, -1}, {+1, -1}}
	case 'S':
		return []Point{{-1, +1}, {0, +1}, {+1, +1}}
	case 'W':
		return []Point{{-1, -1}, {-1, 0}, {-1, +1}}
	case 'E':
		return []Point{{+1, -1}, {+1, 0}, {+1, +1}}
	}
	return []Point{}
}

// returns true if points reached by modifying "cur" current point with "mod" modifies are empty, false otherwise
func checkFreeDirection(cur Point, mod []Point) bool {
	for _, m := range mod {
		_, found := elves[Point{cur.x + m.x, cur.y + m.y}]
		if found { // elf found
			return false
		}
	}
	return true
}

// returns a new map of elves with elves moved to proposed position if possible
// returns true if at least one elf moved, false otherwise
func moveToProposed() (map[Point]Elve, bool) {

	newElves := make(map[Point]Elve) // new elves map

	var someoneMoved bool // at least one elf moved

	for p := range elves { // scan each elf

		if elves[p].toMove { // if elf precalculated a valid position to move

			// check position to move is unique (no other elf wants to move there)
			uniqueMove := true
			for p2 := range elves {
				if p == p2 { // skip self
					continue
				}
				if elves[p2].toMove && elves[p].proposedMove == elves[p2].proposedMove { // another elf wants to move there
					uniqueMove = false // shouldnt move
					break
				}
			}

			if uniqueMove {
				newElves[elves[p].proposedMove] = elves[p] // move elf to proposed move
				someoneMoved = true
			} else {
				newElves[p] = elves[p] // elf stays in same position
			}

		} else { // elf didnt want to move
			newElves[p] = elves[p] // elf stays in same position
		}
	}
	return newElves, someoneMoved
}

// returns the number of empty spaces in the area of elves
func countEmpty() (count int) {
	minX, maxX, minY, maxY := getGridBorders() // get dimensions of area to look in

	// scan each position
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {

			_, found := elves[Point{x, y}]
			if !found { // position empty
				count++
			}
		}
	}
	return count
}

// returns the maximum and minimun x and y coordinate
func getGridBorders() (int, int, int, int) {
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for p := range elves {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	return minX, maxX, minY, maxY
}
