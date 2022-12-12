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

func checkHeight(grid []string, curX, curY, targX, targY int) bool {
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

// returns the height of a point
func getPointHeight(h rune) int {
	return int(math.Abs(float64('a' - h)))
}

func depthFirstSearch(grid []string, cur Point) map[Point]int {
	queue := queue{nil}
	distances := make(map[Point]int)
	reached := make(map[Point]bool)
	queue.enqueue(cur)

	for !queue.isEmpty() {
		u := queue.dequeue()

		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if !(x == 0 || y == 0) {
					continue
				}
				if x == 0 && y == 0 {
					continue
				}

				if (u.x+x) < 0 || (u.x+x) >= len(grid[0]) {
					continue
				}
				if (u.y+y) < 0 || (u.y+y) >= len(grid) {
					continue
				}

				if !checkHeight(grid, u.x, u.y, u.x+x, u.y+y) {
					continue
				}

				v := Point{u.x + x, u.y + y}
				if !reached[v] {
					distances[v] = distances[u] + 1
					reached[v] = true
					queue.enqueue(v)
				}
			}
		}
	}
	return distances
}

func getSandAtoE(grid []string, destination Point) (int, int) {
	var StoE int = math.MaxInt
	var minAtoE int = math.MaxInt
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'S' {
				StoE = depthFirstSearch(grid, Point{x, y})[destination]
			}
			if grid[y][x] == 'a' {
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
