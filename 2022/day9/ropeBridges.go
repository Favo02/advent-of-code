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

	var rope []Point
	var crossed map[Point]bool

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

		crossed[Point{0, 0}] = true

		for _, instruction := range instructions {
			moveHead(instruction, rope, crossed, tail)
		}
		fmt.Print("Points crossed by the tail (of length ", tail, "): ", len(crossed), "\n")
	}
}

func parseInput() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func moveHead(instruction string, rope []Point, crossed map[Point]bool, tailLength int) ([]Point, map[Point]bool) {
	tokens := strings.Split(instruction, " ")
	direction := tokens[0]
	amount, _ := strconv.Atoi(tokens[1])
	for i := 0; i < amount; i++ {
		moveStep(direction, rope)
		crossed[rope[tailLength-1]] = true
	}

	return rope, crossed
}

func moveStep(dir string, rope []Point) []Point {
	switch dir {
	case "U":
		rope[0].y--
	case "D":
		rope[0].y++
	case "L":
		rope[0].x--
	case "R":
		rope[0].x++
	}

	for i := 1; i < len(rope); i++ {
		if !isTailClose(rope[i], rope[i-1]) {
			rope = moveTailPoint(rope, i, i-1)
		}
	}

	return rope
}

func isTailClose(currentPoint, targetPoint Point) bool {
	for i := targetPoint.y - 1; i < targetPoint.y+2; i++ {
		for j := targetPoint.x - 1; j < targetPoint.x+2; j++ {
			if i == currentPoint.y && j == currentPoint.x {
				return true
			}
		}
	}
	return false
}

func moveTailPoint(rope []Point, currentIndex, targetIndex int) []Point {
	// same column
	if rope[targetIndex].x == rope[currentIndex].x {
		if rope[targetIndex].y > rope[currentIndex].y {
			rope[currentIndex].y++
		} else {
			rope[currentIndex].y--
		}
		return rope
	}

	// same row
	if rope[targetIndex].y == rope[currentIndex].y {
		if rope[targetIndex].x > rope[currentIndex].x {
			rope[currentIndex].x++
		} else {
			rope[currentIndex].x--
		}
		return rope
	}

	// diagonal
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
