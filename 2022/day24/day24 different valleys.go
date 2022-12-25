// https://adventofcode.com/2022/day/24
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

type Blizzard struct {
	direction rune
}

var dirModifiers []Point = []Point{{0, -1}, {+1, 0}, {0, +1}, {-1, 0}}

func main() {
	valley := parseInput()
	printValley(valley)
	dist := depthFirstSearch(valley, Point{1, 0})
	fmt.Println("dist:", dist)
	fmt.Println("dist:", dist[Point{6, 5}])
}

// modifies valley placing the blizzard parsed from stdin
// modifies stdin
func parseInput() map[Point][]Blizzard {
	valley := make(map[Point][]Blizzard)
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
	return valley
}

var time int

func depthFirstSearch(valley map[Point][]Blizzard, start Point) map[Point]int {
	queue := queue{nil}
	distances := make(map[Point]int)
	distances[start] = 0
	reached := make(map[Point]bool)
	reached[start] = true

	queue.enqueue(start)

	for !queue.isEmpty() {
		u := queue.dequeue()

		valley = nextMinute(valley)
		printValley(valley)

		reach := reachable(valley, u)

		// nodes should be visitable more than once

		for _, v := range reach {
			if !reached[v] {
				distances[v] = distances[u] + 1
				reached[v] = true
				queue.enqueue(v)
			}
		}

		if len(reach) == 0 {
			queue.enqueue(u)
		}

	}
	return distances
}

// returns the points reachable from "u"
func reachable(valley map[Point][]Blizzard, u Point) (reac []Point) {
	// scan each point reachable from current point (cur)
	for _, dirMod := range dirModifiers {

		// point reached from cur
		v := Point{u.x + dirMod.x, u.y + dirMod.y}

		blizzards, found := valley[v]

		if found && len(blizzards) == 0 {
			reac = append(reac, v)
		}
	}

	// check if current point still safe
	blizzards, found := valley[u]
	if found && len(blizzards) == 0 {
		reac = append(reac, u)
	}

	return reac
}

// modifies valley moving the blizzards to next minute
func nextMinute(valley map[Point][]Blizzard) map[Point][]Blizzard {
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
				pacman := pacmanEffect(valley, newBlizPoint, blizMod)
				newValley[pacman] = append(newValley[pacman], bliz)
			}
		}
	}
	return newValley
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

func pacmanEffect(valley map[Point][]Blizzard, p, mod Point) Point {
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

func printValley(valley map[Point][]Blizzard) {
	for y := 0; y < 6; y++ {
		for x := 0; x < 8; x++ {
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
