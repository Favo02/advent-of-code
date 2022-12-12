// https://adventofcode.com/2022/day/12
// using DepthFirstSearch

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	grid, destination := parseInput()
	StoE, minAtoE := getSandAtoE(grid, destination)

	fmt.Println("S to E (part1):\t\t", StoE)
	fmt.Println("min a to E (part2):\t", minAtoE)
}

type Point struct {
	x, y int
}

// REQUIRES: stdin is a valid challenge input
// MODIFIES: stdin
// EFFECTS: returns the map parsed as a grid and destination point coordinates
func parseInput() (grid []string, destination Point) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c == 'E' {
				destination = Point{i, len(grid)}
			}
		}
		grid = append(grid, line)
	}
	return grid, destination
}

// REQUIRES: grid is a valid map, curX, curY, targX, targY are valid grid coordinates
// EFFECTS: returns true if targX, targY is reachable da curX, curY, false otherwise
func checkHeightDifference(grid []string, curX, curY, targX, targY int) bool {
	cur := grid[curY][curX]
	target := grid[targY][targX]

	if cur == 'S' {
		cur = 'a'
	}
	if target == 'E' {
		target = 'z'
	}

	diff := getPointHeight(rune(target)) - getPointHeight(rune(cur))
	return diff <= 1
}

// EFFECTS: returns the height of a point (relative to 'a')
func getPointHeight(h rune) int {
	return int(h - 'a')
}

// REQUIRES: grid is valid map, cur is a point on grid
// EFFECTS: returns the amount of steps needed to reach every point (key of map) reachable on the grid
func depthFirstSearch(grid []string, cur Point) map[Point]int {
	queue := queue{nil}              // queue of points to be analyzed
	distances := make(map[Point]int) // distance from 'cur' to key of map
	reached := make(map[Point]bool)  // points already reached

	queue.enqueue(cur) // first point to analyze

	for !queue.isEmpty() { // while there are points to analyze
		u := queue.dequeue() // first point to analyze

		// scan each point reachable from u
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				// only move top, down, left or right, no diagonal
				if !(x == 0 || y == 0) {
					continue
				}
				// do not analyze self
				if x == 0 && y == 0 {
					continue
				}

				// out of grid bounds
				if (u.x+x) < 0 || (u.x+x) >= len(grid[0]) {
					continue
				}
				if (u.y+y) < 0 || (u.y+y) >= len(grid) {
					continue
				}

				// point not reachable from cur (height difference)
				if !checkHeightDifference(grid, u.x, u.y, u.x+x, u.y+y) {
					continue
				}

				v := Point{u.x + x, u.y + y} // point reached
				if !reached[v] {             // not reached yet
					distances[v] = distances[u] + 1 // distance to reach v
					reached[v] = true               // set as reached
					queue.enqueue(v)                // add to points to analyze
				}
			}
		}
	}
	return distances
}

// REQUIRES: destination is a valid point on grid
// EFFECTS: returns the minimum amount of steps needed to go from 'S' to destination and the minimum amount of steps needed to go from the closest 'a' and destination
func getSandAtoE(grid []string, destination Point) (int, int) {
	var StoE int = math.MaxInt
	var minAtoE int = math.MaxInt

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'S' {
				// get distances to each point reachable, extract only the one to "destination"
				StoE = depthFirstSearch(grid, Point{x, y})[destination]
			}
			if grid[y][x] == 'a' {
				// get distances to each point reachable, extract only the one to "destination", save only the minimun one
				AtoE, found := depthFirstSearch(grid, Point{x, y})[destination]
				if found && AtoE < minAtoE {
					minAtoE = AtoE
				}
			}
		}
	}
	return StoE, minAtoE
}

// QUEUE

type queue struct {
	head *queueNode
}

type queueNode struct {
	next  *queueNode
	point Point
}

func (q *queue) enqueue(p Point) {
	if q.head == nil {
		q.head = &queueNode{nil, p}
		return
	}
	node := q.head
	for node.next != nil {
		node = node.next
	}
	newNode := queueNode{nil, p}
	node.next = &newNode
}

func (q *queue) dequeue() Point {
	head := q.head
	q.head = q.head.next
	return head.point
}

func (q *queue) isEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}
