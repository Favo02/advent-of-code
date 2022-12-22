// https://adventofcode.com/2022/day/22
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

type Point struct {
	x, y int
}

var board map[Point]bool
var moveQty []int
var moveDir []rune
var dirs []rune = []rune{'R', 'D', 'L', 'U'}

func main() {
	board = make(map[Point]bool)
	parseInput()
	start := getStartingPoint()
	end, dir := executeInstructions(start)
	fmt.Println(((end.y + 1) * 1000) + ((end.x + 1) * 4) + dir)
}

func parseInput() {
	scanner := bufio.NewScanner(os.Stdin)

	var tempInstructions []rune

	var row int
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 && !unicode.IsDigit(rune(line[0])) {
			for column, c := range line {
				if c == '.' { // valid point
					board[Point{column, row}] = true
				}
				if c == '#' { // wall
					board[Point{column, row}] = false
				}
			}
			row++
		} else {
			for _, inst := range line {
				tempInstructions = append(tempInstructions, inst)
			}
		}

	}
	moveQty, moveDir = fixInstructions(tempInstructions)
}

func fixInstructions(tempInstructions []rune) (move []int, dir []rune) {

	var digitBuffer string

	for _, v := range tempInstructions {
		if unicode.IsDigit(v) {
			digitBuffer = fmt.Sprint(digitBuffer, string(v))
			continue
		} else {
			digit, _ := strconv.Atoi(digitBuffer)
			digitBuffer = ""
			move = append(move, digit)

			dir = append(dir, v)
		}
	}
	digit, _ := strconv.Atoi(digitBuffer)
	digitBuffer = ""
	move = append(move, digit)

	return move, dir
}

func getStartingPoint() Point {
	minX := math.MaxInt
	for k := range board {
		if k.y == 0 && k.x < minX {
			minX = k.x
		}
	}
	return Point{minX, 0}
}

func executeInstructions(start Point) (Point, int) {
	cur := start
	dir := 0 // starting direction

	for i, qty := range moveQty {
		cur = move(cur, qty, dirs[dir])

		if i < len(moveDir) {
			switch moveDir[i] {
			case 'R':
				dir = (dir + 1) % 4
			case 'L':
				dir = ((dir-1)%4 + 4) % 4
			}
		}
	}

	return cur, dir
}

func move(cur Point, qty int, dir rune) Point {
	var modX, modY int
	switch dir {
	case 'R':
		modX = +1
		modY = 0
	case 'D':
		modX = 0
		modY = +1
	case 'L':
		modX = -1
		modY = 0
	case 'U':
		modX = 0
		modY = -1
	}

	for i := 0; i < qty; i++ {
		newP := Point{cur.x + modX, cur.y + modY}
		valid, found := board[newP]

		// out of range, pacman effect
		if !found {
			newP = pacManEffect(newP, modX, modY)
			valid, found = board[newP]
		}

		// newP is an empty space, move there
		if found && valid {
			cur = newP
			continue
		}

		// newP is a wall, stop moving
		if found && !valid {
			break
		}

	}

	return cur
}

func pacManEffect(cur Point, modX, modY int) Point {
	for true {
		next := Point{cur.x - modX, cur.y - modY}
		_, found := board[next]
		if !found {
			return cur
		} else {
			cur = next
		}
	}
	return cur
}
