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
var faces map[Point]int

var moveQty []int
var moveDir []rune
var dirs []rune = []rune{'R', 'D', 'L', 'U'}

const FACE_SIZE int = 50

func main() {
	board = make(map[Point]bool)
	faces = make(map[Point]int)
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
				var face int

				if row < 50 {
					if column >= 50 && column < 100 {
						face = 1
					} else if column >= 100 {
						face = 2
					}
				} else if row >= 50 && row < 100 {
					face = 3
				} else if row >= 100 && row < 150 {
					if column < 50 {
						face = 4
					} else if column >= 50 {
						face = 5
					}
				} else if row >= 150 {
					face = 6
				}

				if c == '.' { // valid point
					board[Point{column, row}] = true
					faces[Point{column, row}] = face
				}
				if c == '#' { // wall
					board[Point{column, row}] = false
					faces[Point{column, row}] = face
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
	var runeDir rune

	for i, qty := range moveQty {
		// fmt.Println(cur, qty, string(dirs[dir]))
		cur, runeDir = move(cur, qty, dirs[dir])
		dir = runeDirToIndex(runeDir)
		// fmt.Println(cur)
		// fmt.Println()

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

func move(cur Point, qty int, dir rune) (Point, rune) {
	var modX, modY int

	for i := 0; i < qty; i++ {

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

		newP := Point{cur.x + modX, cur.y + modY}
		valid, found := board[newP]

		// fmt.Println(cur.x+1, cur.y+1, string(dir))

		var newDir rune

		// out of range, pacman effect
		if !found {
			// fmt.Println("out at:", newP)
			newP, newDir = pacManEffect(cur, dir)
			// fmt.Println("in at:", newP)
			// fmt.Println()
			valid, found = board[newP]
			if !(found && !valid) {
				dir = newDir
			}
		}

		// newP is an empty space, move there
		if found && valid {
			// fmt.Println("valid")
			cur = newP
			continue
		}

		// newP is a wall, stop moving
		if found && !valid {
			// fmt.Println("wall")
			break
		}

	}

	return cur, dir
}

func runeDirToIndex(runeDir rune) int {
	switch runeDir {
	case 'R':
		return 0
	case 'D':
		return 1
	case 'L':
		return 2
	case 'U':
		return 3
	}
	return -1
}

func pacManEffect(cur Point, dir rune) (Point, rune) {
	curFace := faces[cur]
	// fmt.Println("curface:", curFace, "dir:", string(dir))
	relX := cur.x - getBordersByFace(curFace).minX
	relY := cur.y - getBordersByFace(curFace).minY

	switch faces[cur] {
	case 1:
		switch dir {
		case 'U':
			return Point{getBordersByFace(6).minX, getBordersByFace(6).minY + relX}, 'R'
		case 'L':
			return Point{getBordersByFace(4).minX, getBordersByFace(4).minY + (49 - relY)}, 'R'
		}

	case 2:
		switch dir {
		case 'U':
			return Point{getBordersByFace(6).minX + relX, getBordersByFace(6).maxY}, 'U'
		case 'R':
			return Point{getBordersByFace(5).maxX, getBordersByFace(5).minY + (49 - relY)}, 'L'
		case 'D':
			return Point{getBordersByFace(3).maxX, getBordersByFace(3).minY + relX}, 'L'
		}

	case 3:
		switch dir {
		case 'L':
			return Point{getBordersByFace(4).minX + relY, getBordersByFace(4).minY}, 'D'
		case 'R':
			return Point{getBordersByFace(2).minX + relY, getBordersByFace(2).maxY}, 'U'
		}

	case 4:
		switch dir {
		case 'U':
			return Point{getBordersByFace(3).minX, getBordersByFace(3).minY + relX}, 'R'
		case 'L':
			return Point{getBordersByFace(1).minX, getBordersByFace(1).minY + (49 - relY)}, 'R'
		}

	case 5:
		switch dir {
		case 'R':
			return Point{getBordersByFace(2).maxX, getBordersByFace(2).minY + (49 - relY)}, 'L'
		case 'D':
			return Point{getBordersByFace(6).maxX, getBordersByFace(6).minY + relX}, 'L'
		}

	case 6:
		switch dir {
		case 'L':
			return Point{getBordersByFace(1).minX + relY, getBordersByFace(1).minY}, 'D'
		case 'R':
			return Point{getBordersByFace(5).minX + relY, getBordersByFace(5).maxY}, 'U'
		case 'D':
			return Point{getBordersByFace(2).minX + relX, getBordersByFace(2).minY}, 'D'
		}

	}

	return cur, 'Z'
}

type facePosition struct {
	minX, maxX, minY, maxY int
}

// returns face borders: minX, maxX, minY, maxY
func getBordersByFace(face int) facePosition {
	switch face {
	case 1:
		return facePosition{50, 99, 0, 49}
	case 2:
		return facePosition{100, 149, 0, 49}
	case 3:
		return facePosition{50, 99, 50, 99}
	case 4:
		return facePosition{0, 49, 100, 149}
	case 5:
		return facePosition{50, 99, 100, 149}
	case 6:
		return facePosition{0, 49, 150, 199}
	}
	return facePosition{0, 0, 0, 0}
}
