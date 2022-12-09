// https://adventofcode.com/2022/day/9

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TAIL int = 10 // 2 for part1, 10 for part2

type Point struct {
	x int
	y int
}

func main() {
	instructions := parseInput()

	// points are relative to starting point, can be negative
	rope := make([]Point, TAIL)     // current rope points positions
	crossed := make(map[Point]bool) // points crossed by the tail of the rope

	crossed[Point{0, 0}] = true

	for _, instruction := range instructions {
		moveHead(instruction, rope, crossed)
	}
	fmt.Println("Points crossed by the tail:", len(crossed))
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

func moveHead(instruction string, rope []Point, crossed map[Point]bool) ([]Point, map[Point]bool) {
	tokens := strings.Split(instruction, " ")
	direction := tokens[0]
	amount, _ := strconv.Atoi(tokens[1])
	for i := 0; i < amount; i++ {
		moveStep(direction, rope)
		crossed[rope[TAIL-1]] = true
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
