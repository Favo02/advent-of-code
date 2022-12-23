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

	for i := 0; true; i++ {
		someoneMoved := round(i)
		fmt.Println("round:", i+1)
		if !someoneMoved {
			break
		}
	}
}

// modifies elves map placing them
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

func round(nRound int) bool {

	newElves := make(map[Point]Elve)

	// generate proposed
	for p := range elves {

		// scan 8 positions around
		someoneClose := scanAround(p)

		// noone around, do nothing
		if !someoneClose {
			elves[p] = Elve{Point{}, false}
			continue
		}

		// direction to start looking

		// check direction
		i := 0
		dirFound := false
		var finalDirection int
		for i < 4 {
			if checkFreeDirection(p, directionToModifiers(directions[((nRound%4)+i)%4])) {
				finalDirection = ((nRound % 4) + i) % 4
				dirFound = true
				break
			}
			i++
		}

		// update elf
		if dirFound {
			nextPointX := p.x + directionToModifiers(directions[finalDirection])[1].x
			nextPointY := p.y + directionToModifiers(directions[finalDirection])[1].y
			elves[p] = Elve{Point{nextPointX, nextPointY}, true}
		} else {
			elves[p] = Elve{Point{}, false}
		}
	}

	var someoneMoved bool

	// move to proposed
	for p := range elves {

		if elves[p].toMove { // if point has point to move

			move := true
			// check unique
			for p2 := range elves {
				if p == p2 {
					continue
				}
				if elves[p2].toMove && elves[p].proposedMove == elves[p2].proposedMove {
					move = false
					break
				}
			}

			if move {

				if !(p == elves[p].proposedMove) {
					someoneMoved = true
				}

				newElves[elves[p].proposedMove] = elves[p]

			} else {
				newElves[p] = elves[p]
			}

		} else { // point has nowhere to move
			newElves[p] = elves[p]
		}
	}

	elves = newElves

	return someoneMoved
}

func scanAround(p Point) bool {
	for y := p.y - 1; y < p.y-1+3; y++ {
		for x := p.x - 1; x < p.x-1+3; x++ {
			if y == p.y && x == p.x {
				continue
			}
			_, found := elves[Point{x, y}]
			if found {
				return true
			}
		}
	}
	return false
}

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

func checkFreeDirection(cur Point, mod []Point) bool {
	for _, m := range mod {
		_, found := elves[Point{cur.x + m.x, cur.y + m.y}]
		if found {
			return false
		}
	}
	return true
}

func countEmpty() (count int) {
	minX, maxX, minY, maxY := getGridBorders()
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			_, found := elves[Point{x, y}]
			if !found {
				count++
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	return count
}

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
