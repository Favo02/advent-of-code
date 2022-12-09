// https://adventofcode.com/2022/day/9

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	const TAIL_LENGTH_p1 int = 2  // tail for part1 (length 2)
	const TAIL_LENGTH_p2 int = 10 // tail for part2 (length 10)

	instructions := parseInput()

	var rope []Point           // current points of the rope
	var crossed map[Point]bool // points crossed by the tail (last point) of the rope

	// execute 2 times, one for part1, one for part2
	for i := 0; i < 2; i++ {
		var tail int
		if i == 0 {
			tail = TAIL_LENGTH_p1
		} else {
			tail = TAIL_LENGTH_p2
		}

		// points are relative to starting point, can be negative
		rope = make([]Point, tail)     // current rope points positions
		crossed = make(map[Point]bool) // points crossed by the tail of the rope

		crossed[Point{0, 0}] = true // starting point is crossed by every point of the rope

		for _, instruction := range instructions {
			moveHead(instruction, rope, crossed, tail) // move head in the direction of instruction (the tail of the rope will follow the head)
		}
		fmt.Print("Points crossed by the tail (of length ", tail, "):\n\t", len(crossed), "\n")
	}
}

// REQUIRES: stdin is a valid challenge input
// MODIFIES: stdin
// EFFECTS: returns the instruction list
func parseInput() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

// REQUIRES: instruction is a valid instruction format (<direction> <amount>)
// EFFECTS: returns the updated rope updated (moved by the instruction) and the updated points crossed by the tail (moved by the instruction)
func moveHead(instruction string, rope []Point, crossed map[Point]bool, tailLength int) ([]Point, map[Point]bool) {
	tokens := strings.Split(instruction, " ")
	direction := tokens[0]
	amount, _ := strconv.Atoi(tokens[1])
	for i := 0; i < amount; i++ { // move the head the amount of times the instruction requires
		moveStep(direction, rope)          // moves the head in the direction 1 time
		crossed[rope[tailLength-1]] = true // mark as crossed the point the rope tail is currently
	}

	return rope, crossed
}

// REQUIRES: dir is a valid direction ("U", "D", "L", "R"), rope has at least one Point
// EFFECTS: returns the rope moved in the direction "dir" (updates both head of rope and all tail points)
func moveStep(dir string, rope []Point) []Point {
	switch dir {
	case "U":
		rope[0].y-- // move head up
	case "D":
		rope[0].y++ // move head down
	case "L":
		rope[0].x-- // move head left
	case "R":
		rope[0].x++ // move head right
	}

	// make tail points follow the head
	for i := 1; i < len(rope); i++ { // scan every point of the rope (except the head)
		if !isTailClose(rope[i], rope[i-1]) { // if the point is not close to his predecessor
			rope = moveTailPoint(rope, i, i-1) // move the point close to his predecessor
		}
	}

	return rope
}

// REQUIRES: currentPoint, targetPoint not nil
// EFFECTS: returns true if target point is close (if it is tha same point or if it is in one of the 8 cells surrounding the point), false otherwise
func isTailClose(currentPoint, targetPoint Point) bool {
	// scan each of the 8 (+ the same point) surrounding the targetPoint
	for i := targetPoint.y - 1; i < targetPoint.y+2; i++ {
		for j := targetPoint.x - 1; j < targetPoint.x+2; j++ {
			// if the current point is in this 9 points then it is close to targetPoint
			if i == currentPoint.y && j == currentPoint.x {
				return true
			}
		}
	}
	return false
}

// REQUIRES: rope is not nil, currentIndex and targetIndex are valid indexes of rope
// EFFECTS: returns the rope updated with the point at "currentIndex" moved if it is not close to "targetIndex" point
func moveTailPoint(rope []Point, currentIndex, targetIndex int) []Point {
	// same column, so move point only horizontally
	if rope[targetIndex].x == rope[currentIndex].x {
		if rope[targetIndex].y > rope[currentIndex].y {
			rope[currentIndex].y++
		} else {
			rope[currentIndex].y--
		}
		return rope
	}

	// same row, so move point only vertically
	if rope[targetIndex].y == rope[currentIndex].y {
		if rope[targetIndex].x > rope[currentIndex].x {
			rope[currentIndex].x++
		} else {
			rope[currentIndex].x--
		}
		return rope
	}

	// not same row or same column, so move point diagonally
	if rope[currentIndex].x < rope[targetIndex].x && rope[currentIndex].y < rope[targetIndex].y {
		rope[currentIndex].x++
		rope[currentIndex].y++
		return rope
	}
	if rope[currentIndex].x > rope[targetIndex].x && rope[currentIndex].y < rope[targetIndex].y {
		rope[currentIndex].x--
		rope[currentIndex].y++
		return rope
	}
	if rope[currentIndex].x < rope[targetIndex].x && rope[currentIndex].y > rope[targetIndex].y {
		rope[currentIndex].x++
		rope[currentIndex].y--
		return rope
	}
	if rope[currentIndex].x > rope[targetIndex].x && rope[currentIndex].y > rope[targetIndex].y {
		rope[currentIndex].x--
		rope[currentIndex].y--
		return rope
	}
	return rope
}
