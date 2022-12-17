// https://adventofcode.com/2022/day/12
// usign Dijkstra algorithm

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var sizeX, sizeY int
var map_ []string
var steps int

func main() {
	map__, cur, dest := parseInput()
	map_ = map__
	sizeX = len(map_[0])
	sizeY = len(map_)
	fmt.Println("S:", cur.x, cur.y)
	fmt.Println("E:", dest.x, dest.y)
	fmt.Println("sx:", sizeX, "sy:", sizeY)

	var minDist int = math.MaxInt
	aCount := 0
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			if map__[y][x] == 'a' {
				fmt.Println("from", x, y, "- a n°", aCount)
				aCount++
				distances := dijkstra(Point{x, y})
				distFromE := distances[dest]
				fmt.Println(distFromE)
				fmt.Println()
				if distFromE < minDist {
					minDist = distFromE
				}
			}
		}
	}
	fmt.Println("res:", minDist)
}

type Point struct {
	x, y int
}

func parseInput() (lines []string, cur, dest Point) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c == 'S' {
				cur = Point{i, len(lines)}
			}
			if c == 'E' {
				dest = Point{i, len(lines)}
			}
		}
		lines = append(lines, line)
	}
	return lines, cur, dest
}

func dijkstra(cur Point) map[Point]int {

	// create map of size number of nodes
	distances := make(map[Point]int)
	c := make(map[Point]int)

	// initialize each point to inf
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			distances[Point{x, y}] = math.MaxInt
			c[Point{x, y}] = math.MaxInt

			// starting point to 0
			if x == cur.x && y == cur.y {
				distances[Point{x, y}] = 0
				c[Point{x, y}] = 0
			}
		}
	}

	// while c not empty
	for len(c) > 0 {

		// fmt.Println(len(c))

		// u = min v
		var min int = math.MaxInt
		var u Point
		for k := range c {
			if distances[k] < min {
				min = distances[k]
				u = k
			}
		}

		// remove minPoint from c
		// fmt.Println(u, min)
		delete(c, u)

		// FOR EACH (u,v) ∈ E DO
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if !(x == 0 || y == 0) {
					continue
				}
				if x == 0 && y == 0 {
					continue
				}

				if (u.x+x) < 0 || (u.x+x) >= sizeX {
					continue
				}
				if (u.y+y) < 0 || (u.y+y) >= sizeY {
					continue
				}

				// u = u
				// v = u.x + x, u.y + y
				if !checkHeight(u.x, u.y, u.x+x, u.y+y) {
					continue
				}

				// IF d[u] + w(u,v) < d[v]
				weight := 1
				if (distances[u] + weight) < distances[Point{u.x + x, u.y + y}] {
					distances[Point{u.x + x, u.y + y}] = distances[u] + weight
				}

			}
		}

		unreachable := true
		for cLeft := range c {
			if distances[cLeft] != math.MaxInt {
				unreachable = false
			}
		}

		if unreachable {
			break
		}
	}
	return distances

}

func checkHeight(curX, curY, targX, targY int) bool {
	cur := map_[curY][curX]
	target := map_[targY][targX]

	if cur == 'S' {
		cur = 'a'
	}
	if (cur == 'y' || cur == 'z') && target == 'E' {
		return true
	}

	diff := getPointHeight(rune(target)) - getPointHeight(rune(cur))

	if diff <= 1 {
		return true
	}
	return false
}

// returns the height of a point
func getPointHeight(h rune) int {
	return int(math.Abs(float64('a' - h)))
}
