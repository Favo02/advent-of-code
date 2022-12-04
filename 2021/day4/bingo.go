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
	getBingoScore(nums, boards)
}

func parseInput() ([]int, [][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	// numbers
	scanner.Scan()
	numsStr := strings.Split(scanner.Text(), ",")
	nums := strToNum(numsStr)

	// boards
	var i int = -1
	var boards [][]int
	for scanner.Scan() {
		line := scanner.Text()

		// end of board
		if line == "" {
			boards = append(boards, make([]int, 0))
			i++
			continue
		}

		// board lines
		numStr := strings.Split(line, " ")
		num := strToNum(numStr)
		boards[i] = append(boards[i], num...)
		boards[i] = append(boards[i], -1)
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

func getBingoScore(nums []int, boards [][]int) int {
	for _, num := range nums {
		for _, v := range v {

		}
	}
}
