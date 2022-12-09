// https://adventofcode.com/2022/day/9

// ------------------------------
//
// 			NOT REFACTORED YET
//
// ------------------------------

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TAIL int = 10 // 2 for part1, 10 for part2

// Point are relative to starting point, can be negative
var tailCrossed map[Point]bool = make(map[Point]bool)
var curRope []Point = make([]Point, TAIL)

type Point struct {
	x int
	y int
}

func main() {
	lines := parseInput()
	tailCrossed[Point{0, 0}] = true
	for _, line := range lines {
		moveHead(line)
	}
	count := countTail()
	fmt.Println(count)
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

func moveHead(line string) {
	tokens := strings.Split(line, " ")
	dir := tokens[0]
	amnt, _ := strconv.Atoi(tokens[1])
	for i := 0; i < amnt; i++ {
		moveStep(dir)
	}
}

func moveStep(dir string) {
	switch dir {
	case "U":
		curRope[0].y--
	case "D":
		curRope[0].y++
	case "L":
		curRope[0].x--
	case "R":
		curRope[0].x++
	}

	for i := 1; i < len(curRope); i++ {
		if !isTailClose(i, curRope[i-1]) {
			moveTailPoint(i, curRope[i-1])
		}
	}

	tailCrossed[curRope[TAIL-1]] = true
}

func isTailClose(ii int, targetPoint Point) bool {
	if targetPoint.x == curRope[ii].x && targetPoint.y == curRope[ii].y {
		return true
	}
	for i := targetPoint.y - 1; i < targetPoint.y+2; i++ {
		for j := targetPoint.x - 1; j < targetPoint.x+2; j++ {
			if i == curRope[ii].y && j == curRope[ii].x {
				return true
			}
		}
	}
	return false
}

func moveTailPoint(i int, targetPoint Point) {
	if targetPoint.x == curRope[i].x {
		if targetPoint.y > curRope[i].y {
			curRope[i].y++
		} else {
			curRope[i].y--
		}
		return
	}

	if targetPoint.y == curRope[i].y {
		if targetPoint.x > curRope[i].x {
			curRope[i].x++
		} else {
			curRope[i].x--
		}
		return
	}

	if curRope[i].x < targetPoint.x && curRope[i].y < targetPoint.y {
		curRope[i].x++
		curRope[i].y++
		return
	}
	if curRope[i].x > targetPoint.x && curRope[i].y < targetPoint.y {
		curRope[i].x--
		curRope[i].y++
		return
	}
	if curRope[i].x < targetPoint.x && curRope[i].y > targetPoint.y {
		curRope[i].x++
		curRope[i].y--
		return
	}
	if curRope[i].x > targetPoint.x && curRope[i].y > targetPoint.y {
		curRope[i].x--
		curRope[i].y--
		return
	}
}

func countTail() int {
	return len(tailCrossed)
}
