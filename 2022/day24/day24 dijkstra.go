// https://adventofcode.com/2022/day/24
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x, y int
}

type Blizzard struct {
	direction rune
}

var valley map[Point][]Blizzard

var dirModifiers []Point = []Point{{0, -1}, {+1, 0}, {0, +1}, {-1, 0}}

func main() {
	valley = make(map[Point][]Blizzard)
	parseInput()
	printValley(Point{1, 0})
	dist := dijkstra(Point{1, 0})
	fmt.Println("dist:", dist)
	fmt.Println("dist:", dist[Point{6, 5}])
}

// modifies valley placing the blizzard parsed from stdin
// modifies stdin
func parseInput() {
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, v := range line {
			if v == '.' {
				valley[Point{x, y}] = make([]Blizzard, 0)
			}
			if v == '>' || v == '<' || v == '^' || v == 'v' {
				valley[Point{x, y}] = make([]Blizzard, 1)
				valley[Point{x, y}][0] = Blizzard{v}
			}
		}
		y++
	}
}

func dijkstra(cur Point) map[Point]int {

	nextMinute()
	// create map of size number of nodes
	distances := make(map[Point]int)
	queue := make(map[Point]int)

	// initialize each point to inf
	for p := range valley {
		distances[p] = math.MaxInt
		queue[p] = math.MaxInt
	}

	// starting point to 0
	distances[cur] = 0
	queue[cur] = 0

	// while c not empty
	for len(queue) > 0 {

		// cur = min v
		var min int = math.MaxInt
		var cur Point
		for k := range queue {
			if distances[k] < min {
				min = distances[k]
				cur = k
			}
		}

		// remove minPoint from c
		delete(queue, cur)

		fmt.Println(cur)
		fmt.Println(reachable(cur))
		printValley(cur)

		if len(reachable(cur)) == 1 {

		}

		// scan each point reachable from current point (cur)
		for _, v := range reachable(cur) {
			// point reached from cur

			weight := 1 // weight of move from cur to v

			// if this path to reach v is less than old path, save this one
			if (distances[cur] + weight) < distances[v] {
				distances[v] = distances[cur] + weight
			}

		}
		nextMinute()

		// ignore unreachable nodes
		// unreachable := true
		// for cLeft := range queue {
		// 	if distances[cLeft] != math.MaxInt {
		// 		unreachable = false
		// 	}
		// }
		// if unreachable {
		// 	break
		// }
	}
	return distances
}

// returns the points reachable from "u"
func reachable(u Point) (reac []Point) {
	// fmt.Println("reachable")
	// scan each point reachable from current point (cur)
	for _, dirMod := range dirModifiers {

		// point reached from cur
		v := Point{u.x + dirMod.x, u.y + dirMod.y}

		blizzards, found := valley[v]

		if found && len(blizzards) == 0 {
			reac = append(reac, v)
		}
	}
	return append(reac, u)
}

// modifies valley moving the blizzards to next minute
func nextMinute() {
	newValley := make(map[Point][]Blizzard)

	// initialize empty
	for p := range valley {
		newValley[p] = make([]Blizzard, 0)
	}

	// place blizzards
	for p, blizzards := range valley {

		for _, bliz := range blizzards {
			blizMod := getDirectionModifiers(bliz.direction)
			newBlizPoint := Point{p.x + blizMod.x, p.y + blizMod.y}

			_, valid := valley[newBlizPoint]
			if valid {
				newValley[newBlizPoint] = append(newValley[newBlizPoint], bliz)
			} else {
				// pacman effect
				newValley[pacmanEffect(newBlizPoint, blizMod)] = append(newValley[pacmanEffect(newBlizPoint, blizMod)], bliz)
			}
		}
	}
	valley = newValley
}

func getDirectionModifiers(dir rune) Point {
	switch dir {
	case '<':
		return Point{-1, 0}
	case '>':
		return Point{+1, 0}
	case '^':
		return Point{0, -1}
	case 'v':
		return Point{0, +1}
	}
	fmt.Println("err")
	return Point{0, 0}
}

func pacmanEffect(p, mod Point) Point {
	// fmt.Println("pacman")
	for true {
		newP := Point{p.x - mod.x, p.y - mod.y}
		_, valid := valley[newP]
		if !valid {
			return p
		}
		p = newP
	}
	fmt.Println("err")
	return Point{0, 0}
}

func printValley(cur Point) {
	for y := 0; y < 6; y++ {
		for x := 0; x < 8; x++ {
			if x == cur.x && y == cur.y {
				fmt.Print("E")
				continue
			}
			bliz, found := valley[Point{x, y}]
			if !found {
				fmt.Print("#")
			} else if len(bliz) == 0 {
				fmt.Print(".")
			} else if len(bliz) == 1 {
				fmt.Print(string(bliz[0].direction))
			} else {
				fmt.Print(len(bliz))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// QUEUE

type queue struct {
	head *queueNode
}

type queueNode struct {
	next    *queueNode
	payload Point
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
	return head.payload
}

func (q *queue) isEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}
