package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	grid := parseInput()
	fmt.Println(countVisible(grid))
	fmt.Println(countMaxVisibility(grid))
}

func parseInput() [][]int {
	var grid [][]int

	var index int = -1
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, make([]int, 0))
		index++
		for _, n := range line {
			num, _ := strconv.Atoi(string(n))
			grid[index] = append(grid[index], num)
		}
	}
	return grid
}

func countVisible(grid [][]int) int {
	var count int = len(grid)*2 + len(grid[0])*2 - 4

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			// fmt.Println("check", i, j, "tree")
			if checkVisible(grid, i, j) > 0 {
				count++
			}
			// count += checkVisible(grid, i, j)
		}
	}
	return count
}

func checkVisible(grid [][]int, i, j int) int {
	var top, bottom, left, right int = 1, 1, 1, 1
	for ii := i; ii >= 0; ii-- {
		if ii != i && grid[i][j] <= grid[ii][j] {
			// fmt.Println(i, j, "<", ii, j)
			top = 0
		}
	}
	for ii := i; ii < len(grid); ii++ {
		if ii != i && grid[i][j] <= grid[ii][j] {
			bottom = 0
		}
	}
	for jj := j; jj >= 0; jj-- {
		if jj != j && grid[i][j] <= grid[i][jj] {
			left = 0
		}
	}
	for jj := j; jj < len(grid[0]); jj++ {
		if jj != j && grid[i][j] <= grid[i][jj] {
			right = 0
		}
	}
	// fmt.Println(top, bottom, left, right)
	return top + bottom + left + right
}

func countMaxVisibility(grid [][]int) int {
	var maxVis = 0

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			vis := checkVisible2(grid, i, j)
			// fmt.Println("check", i, j, "tree", vis)
			if vis > maxVis {
				maxVis = vis
			}
		}
	}
	return maxVis
}

func checkVisible2(grid [][]int, i, j int) int {
	var top, bottom, left, right int
	for ii := i; ii >= 0; ii-- {
		if ii != i && grid[i][j] <= grid[ii][j] {
			top++
			break
		}
		top++
	}
	top--
	for ii := i; ii < len(grid); ii++ {
		if ii != i && grid[i][j] <= grid[ii][j] {
			bottom++
			break
		}
		bottom++
	}
	bottom--
	for jj := j; jj >= 0; jj-- {
		if jj != j && grid[i][j] <= grid[i][jj] {
			left++
			break
		}
		left++
	}
	left--
	for jj := j; jj < len(grid[0]); jj++ {
		if jj != j && grid[i][j] <= grid[i][jj] {
			right++
			break
		}
		right++
	}
	right--
	// fmt.Println(top, bottom, left, right)
	return top * bottom * left * right
}
