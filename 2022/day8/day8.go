// https://adventofcode.com/2022/day/8

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	grid := parseInput()
	visibleTrees := countVisibleTrees(grid)
	bestVisibility := maxVisibility(grid)
	fmt.Println("number of visible trees (part1):\n\t", visibleTrees)
	fmt.Println("max visibility from a tree (part2):\n\t", bestVisibility)
}

// REQUIRES: stdin is a valid challenge input
// MODIFIES: stdin
// EFFECTS: returns a bidimensional array representing the tree grid
func parseInput() [][]int {
	var grid [][]int

	var index int = 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, make([]int, 0)) // add new line to grid
		for _, n := range line {
			num, _ := strconv.Atoi(string(n))      // parse num
			grid[index] = append(grid[index], num) // insert num into grid
		}

		index++ // move to next line
	}
	return grid
}

// EFFECTS: returns the number of visible tree (trees without any taller tree on each side of them)
func countVisibleTrees(grid [][]int) int {
	// count initialized with tree on the edge (always visible)
	var count int = len(grid)*2 + len(grid[0])*2 - 4 // each side * 2 - 4 corners

	// scan each tree
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {

			// check visibility
			if checkVisible(grid, i, j) {
				count++
			}

		}
	}

	// return number of trees with visibility
	return count
}

// REQUIRES: i, j are valid indexes in grid, i, j are not on the edges of grid (i,j != 0; i,j != len(grid))
// EFFECTS: returns true if the tree in position i, j is visible, false otherwise
func checkVisible(grid [][]int, i, j int) bool {
	var top, bottom, left, right bool = true, true, true, true

	// ij = x,y of the tree we are checking visibility

	// scan every tree above ij
	for ii := i - 1; ii >= 0; ii-- {
		// if a tree is >= of ij height, ij has no visbility
		if grid[i][j] <= grid[ii][j] {
			top = false
		}
	}

	// scan every tree below ij
	for ii := i + 1; ii < len(grid); ii++ {
		// if a tree is >= of ij height, ij has no visbility
		if grid[i][j] <= grid[ii][j] {
			bottom = false
		}
	}

	// scan every tree to the left of ij
	for jj := j - 1; jj >= 0; jj-- {
		// if a tree is >= of ij height, ij has no visbility
		if grid[i][j] <= grid[i][jj] {
			left = false
		}
	}

	// scan every tree to the rigth of ij
	for jj := j + 1; jj < len(grid[0]); jj++ {
		// if a tree is >= of ij height, ij has no visbility
		if grid[i][j] <= grid[i][jj] {
			right = false
		}
	}

	// if ij tree has visibility to at least one side (or between every side)
	return top || bottom || left || right
}

// EFFECTS: returns the visibility (how many trees are visible from that tree, on every side) from the tree with the most visibility
func maxVisibility(grid [][]int) int {
	var maxVis = 0

	// scan each tree
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {

			// calculate visibility from tree
			vis := countVisibility(grid, i, j)

			// if visibility is the highest yet
			if vis > maxVis {
				maxVis = vis
			}
		}
	}

	// return visibility of tree with max visibility
	return maxVis
}

// REQUIRES: i, j are valid indexes in grid, i, j are not on the edges of grid (i,j != 0; i,j != len(grid))
// EFFECTS: returns the visibility (number of trees visible from ij multiplied for each side) from the tree in position i, j
func countVisibility(grid [][]int, i, j int) int {
	var top, bottom, left, right int

	// scan trees above ij
	for ii := i - 1; ii >= 0; ii-- {
		// found a tree taller:
		// add taller tree to visibility and stop scanning
		if grid[i][j] <= grid[ii][j] {
			top++
			break
		}
		// found shorter tree:
		// add shorter tree to visibility
		top++
	}

	// scan trees below ij
	for ii := i + 1; ii < len(grid); ii++ {
		// found a tree taller:
		// add taller tree to visibility and stop scanning
		if grid[i][j] <= grid[ii][j] {
			bottom++
			break
		}
		// found shorter tree:
		// add shorter tree to visibility
		bottom++
	}

	// scan trees to the left of ij
	for jj := j - 1; jj >= 0; jj-- {
		// found a tree taller:
		// add taller tree to visibility and stop scanning
		if grid[i][j] <= grid[i][jj] {
			left++
			break
		}
		// found shorter tree:
		// add shorter tree to visibility
		left++
	}

	// scan trees to the right of ij
	for jj := j + 1; jj < len(grid[0]); jj++ {
		// found a tree taller:
		// add taller tree to visibility and stop scanning
		if grid[i][j] <= grid[i][jj] {
			right++
			break
		}
		// found shorter tree:
		// add shorter tree to visibility
		right++
	}

	// return visibility score
	return top * bottom * left * right
}
