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
	fmt.Println(nums)
	fmt.Println(boards)
	fmt.Println(getBingoScore(nums, boards))
}

func parseInput() ([]int, [][][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	// numbers
	scanner.Scan()
	numsStr := strings.Split(scanner.Text(), ",")
	nums := strToNum(numsStr)

	// boards
	var i int = -1
	var boards [][][]int
	for scanner.Scan() {
		line := scanner.Text()

		// end of board
		if line == "" {
			boards = append(boards, make([][]int, 0))
			i++
			continue
		}

		// board lines
		numStr := strings.Split(line, " ")
		num := strToNum(numStr)
		boardLine := make([]int, 0)
		boardLine = append(boardLine, num...)
		boards[i] = append(boards[i], boardLine)
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
	foundLines := make([][]int, len(boards))
	for i := range foundLines {
		foundLines[i] = make([]int, len(boards[i]))
	}

	foundColumns := make([][]int, len(boards))
	for i := range foundLines {
		foundColumns[i] = make([]int, len(boards[i]))
	}
	foundSum := make([]int, len(boards))

	winBoard := make([]bool, len(boards))
	hasFistWon := false

	var firstWin, lastWin int

	for _, num := range nums {

		// scan boards
		for i := 0; i < len(boards); i++ {

			// scan board lines
			for j := 0; j < len(boards[i]); j++ {

				// scan board line numbers
				for k := 0; k < len(boards[i][j]); k++ {
					if num == boards[i][j][k] {
						foundLines[i][j]++
						foundColumns[i][k]++
						foundSum[i] += boards[i][j][k]
					}

					if (foundLines[i][j] == 5 || foundColumns[i][k] == 5) && !winBoard[i] {
						winBoard[i] = true
						sumBoard := sumBoard(boards[i])
						if !hasFistWon {
							firstWin = (sumBoard - foundSum[i]) * num
							hasFistWon = true
						}
						lastWin = (sumBoard - foundSum[i]) * num
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

func yetToWinBoards(winBoards []bool) int {
	var count int
	for _, b := range winBoards {
		if b == false {
			count++
		}
	}
	return count
}
