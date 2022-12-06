package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums, boards := parseInput()
	first, last := getBingoScore(nums, boards)
	fmt.Println("first board to win (part1):", first)
	fmt.Println("last board to win (part2):", last)
}

func parseInput() ([]int, [][][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	// numbers
	scanner.Scan()
	numsStr := strings.Split(scanner.Text(), ",")
	nums := strToNum(numsStr)

	// boards
	var i int = -1
	var boards [][][]int // array of boards, every board is a bidimensional array
	for scanner.Scan() {
		line := scanner.Text()

		// end of board, append new board
		if line == "" {
			boards = append(boards, make([][]int, 0))
			i++
			continue
		}

		// board lines
		numStr := strings.Split(line, " ")
		num := strToNum(numStr) // line numbers
		boardLine := make([]int, 0)
		boardLine = append(boardLine, num...)
		boards[i] = append(boards[i], boardLine) // append line to board
	}

	return nums, boards
}

func strToNum(strs []string) []int {
	var res []int
	for _, s := range strs {
		if s != "" && s != " " { // double spaces
			n, _ := strconv.Atoi(s)
			res = append(res, n)
		}
	}
	return res
}

func getBingoScore(nums []int, boards [][][]int) (int, int) {
	// how many numbers have been found on each line of each board
	foundLines := make([][]int, len(boards))
	for i := range foundLines { // initialize
		foundLines[i] = make([]int, len(boards[i]))
	}

	// how many numbers have been found on each column of each board
	foundColumns := make([][]int, len(boards))
	for i := range foundLines { // initialize
		foundColumns[i] = make([]int, len(boards[i]))
	}

	// sum of numbers found on each board
	foundSum := make([]int, len(boards))

	// boards that have already won
	winBoard := make([]bool, len(boards))

	// save first board to win
	var firstWin int
	hasFistWon := false

	// last board to win
	var lastWin int

	// for every number
	for _, num := range nums {

		// scan boards
		for i := 0; i < len(boards); i++ {

			// scan board lines
			for j := 0; j < len(boards[i]); j++ {

				// scan board line numbers
				for k := 0; k < len(boards[i][j]); k++ {

					// if number on that board
					if num == boards[i][j][k] {
						foundLines[i][j]++             // increment numbers found on that line
						foundColumns[i][k]++           // increment numbers found on that columns
						foundSum[i] += boards[i][j][k] // add number found to found sum
					}

					// if line or column have found all numbers and that board didnt win yet
					if (foundLines[i][j] == len(boards[i]) || foundColumns[i][k] == len(boards[k])) && !winBoard[i] {
						winBoard[i] = true                           // this board won
						sumBoard := sumBoard(boards[i])              // sum of numbers on that board
						numbersNotDrawnSum := sumBoard - foundSum[i] // total number of that board - sum of numbers drawn

						if !hasFistWon { // if no board won yet save first win
							firstWin = numbersNotDrawnSum * num
							hasFistWon = true
						}
						lastWin = numbersNotDrawnSum * num // save this as last win
					}
				}
			}
		}
	}
	return firstWin, lastWin
}

func sumBoard(board [][]int) int {
	var sum int
	for i := range board {
		for j := range board[i] {
			sum += board[i][j]
		}
	}
	return sum
}
