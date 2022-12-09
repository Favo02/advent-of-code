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

const TAIL int = 1 // 1 for part1, 9 for part2

var grid [][]bool
var curX, curY int // head
var curXs []int = make([]int, TAIL)
var curYs []int = make([]int, TAIL)

func main() {
	lines := parseInput()
	grid = append(grid, make([]bool, 1))
	grid[0][0] = true
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
	// fmt.Println("move:", dir)
	switch dir {
	case "U":
		curY--
		if curY < 0 {
			editGrid('U')
			curY++
			for i := 0; i < len(curYs); i++ {
				curYs[i]++
			}
		}
	case "D":
		curY++
		if curY >= len(grid) {
			editGrid('D')
		}
	case "L":
		curX--
		if curX < 0 {
			editGrid('L')
			curX++
			for i := 0; i < len(curXs); i++ {
				curXs[i]++
			}
		}
	case "R":
		curX++
		if curX >= len(grid) {
			editGrid('R')
		}
	}

	// fmt.Println(grid)

	for i := 0; i < len(curXs); i++ {
		if i == 0 {
			if !isTailClose(i, curX, curY) {
				moveTail(i, curX, curY)
			}
		} else {
			if !isTailClose(i, curXs[i-1], curYs[i-1]) {
				moveTail(i, curXs[i-1], curYs[i-1])
			}
		}
	}
	// fmt.Println("tail:", curXs[8], curYs[8])

	// // fmt.Println("tail:", curXtail, curYtail)
	grid[curYs[TAIL-1]][curXs[TAIL-1]] = true

	// // fmt.Println(grid)

	// fmt.Println("\n---\n")
}

func editGrid(dir rune) {
	switch dir {
	case 'U':
		grid = append(grid, make([]bool, len(grid[0])))
		for i := len(grid) - 1; i > 0; i-- {
			for j := 0; j < len(grid[i]); j++ {
				grid[i][j] = grid[i-1][j]
			}
		}
		for i := 0; i < len(grid[0]); i++ { //reset 1st line
			grid[0][i] = false
		}

	case 'D':
		grid = append(grid, make([]bool, len(grid[0])))

	case 'L':
		for i := 0; i < len(grid); i++ {
			grid[i] = append(grid[i], false)
			for j := len(grid[i]) - 1; j > 0; j-- {
				grid[i][j] = grid[i][j-1]
			}
		}
		// reset 1st column
		for i := 0; i < len(grid); i++ {
			grid[i][0] = false
		}

	case 'R':
		for i := 0; i < len(grid); i++ {
			grid[i] = append(grid[i], false)
		}
	}
}

func isTailClose(ii int, targetX, targetY int) bool {
	if targetX == curXs[ii] && targetY == curYs[ii] {
		return true
	}
	for i := targetY - 1; i < targetY+2; i++ {
		for j := targetX - 1; j < targetX+2; j++ {
			if !(i >= 0 && i < len(grid)) {
				continue
			}
			if !(j >= 0 && j < len(grid[i])) {
				continue
			}
			if i == curYs[ii] && j == curXs[ii] {
				return true
			}
		}
	}
	return false
}

func moveTail(i int, targetX, targetY int) {
	if targetX == curXs[i] {
		if targetY > curYs[i] {
			curYs[i]++
		} else {
			curYs[i]--
		}
		return
	}

	if targetY == curYs[i] {
		if targetX > curXs[i] {
			curXs[i]++
		} else {
			curXs[i]--
		}
		return
	}

	if curXs[i] < targetX && curYs[i] < targetY {
		curXs[i]++
		curYs[i]++
		return
	}
	if curXs[i] > targetX && curYs[i] < targetY {
		curXs[i]--
		curYs[i]++
		return
	}
	if curXs[i] < targetX && curYs[i] > targetY {
		curXs[i]++
		curYs[i]--
		return
	}
	if curXs[i] > targetX && curYs[i] > targetY {
		curXs[i]--
		curYs[i]--
		return
	}
}

func countTail() int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == true {
				count++
			}
		}
	}
	return count
}
