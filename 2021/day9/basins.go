package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	grid := parseInput()
	// printGrid(grid)
	var lowPoints int = 0
	var lowPointsVal int = 0
	var basins [3]int

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			// top
			if (x != 0) && grid[x-1][y] <= grid[x][y] {
				continue
			}
			// bottom
			if (x != len(grid)-1) && grid[x+1][y] <= grid[x][y] {
				continue
			}
			// left
			if (y != 0) && grid[x][y-1] <= grid[x][y] {
				continue
			}
			// right
			if (y != len(grid[x])-1) && grid[x][y+1] <= grid[x][y] {
				continue
			}

			lowPoints++
			lowPointsVal += 1 + grid[x][y]
			basin := 1 + exploreBasin(grid, x, y, 0)

			// save the basin only if is one of the top 3 biggest
			if basin > basins[0] { // first element of basins will always be the smaller
				basins[0] = basin
				sort.Ints(basins[:]) // sort to make the first el the smaller
			}
		}
	}

	Println("number of lowpoints:", lowPoints)
	Println("value of lowpoints:", lowPointsVal)
	Println("top 3 basins:", basins)
	Println("value of basins:", basins[0]*basins[1]*basins[2])
}

func exploreBasin(grid [][]int, x, y, count int) int {
	var valXY int = grid[x][y]
	grid[x][y] = 9
	// top
	if (x != 0) && grid[x-1][y] >= valXY && grid[x-1][y] != 9 {
		count = exploreBasin(grid, x-1, y, count+1)
	}
	// bottom
	if (x != len(grid)-1) && grid[x+1][y] >= valXY && grid[x+1][y] != 9 {
		count = exploreBasin(grid, x+1, y, count+1)
	}
	// left
	if (y != 0) && grid[x][y-1] >= valXY && grid[x][y-1] != 9 {
		count = exploreBasin(grid, x, y-1, count+1)
	}
	// right
	if (y != len(grid[x])-1) && grid[x][y+1] >= valXY && grid[x][y+1] != 9 {
		count = exploreBasin(grid, x, y+1, count+1)
	}
	return count
}

func parseInput() [][]int {
	var grid [][]int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var lineNum []int

		for _, v := range line {
			n, _ := strconv.Atoi(string(v))
			lineNum = append(lineNum, n)
		}
		grid = append(grid, lineNum)
	}

	return grid
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			Print(grid[i][j])
		}
		Println()
	}
}
