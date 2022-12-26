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
	time int
}

type Blizzard struct {
	direction rune
}

var valley map[Point][]Blizzard

var dirModifiers []Point = []Point{{0, -1, +1}, {+1, 0, +1}, {0, +1, +1}, {-1, 0, +1}}

func main() {
	valley = make(map[Point][]Blizzard)
	parseInput()
	for i := 0; i < 500; i++ {
		generateNextMinute(i)
	}

	dist := depthFirstSearch(Point{1, 0, 0})

	minDist := math.MaxInt

	for p, d := range dist {
		if p.x == 120 && p.y == 26 {
			if d < minDist {
				minDist = d
			}
		}
	}
	fmt.Println(minDist)
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
				valley[Point{x, y, 0}] = make([]Blizzard, 0)
			}
			if v == '>' || v == '<' || v == '^' || v == 'v' {
				valley[Point{x, y, 0}] = make([]Blizzard, 1)
				valley[Point{x, y, 0}][0] = Blizzard{v}
			}
		}
		y++
	}
}

// returns the points reachable from "u"
func reachable(u Point) (reac []Point) {
	// scan each point reachable from current point (cur) using every direction modifier
	for _, dirMod := range dirModifiers {

		// point reached from cur
		v := Point{u.x + dirMod.x, u.y + dirMod.y, u.time + 1}

		blizzards, found := valley[v]

		if found && len(blizzards) == 0 {
			reac = append(reac, v)
		}
	}

	// check if current point still safe
	blizzards, found := valley[Point{u.x, u.y, u.time + 1}]
	if found && len(blizzards) == 0 {
		reac = append(reac, Point{u.x, u.y, u.time + 1})
	}
	return reac
}

// modifies valley moving the blizzards to next minute
func generateNextMinute(curTime int) {

	// initialize points empty at time+1
	for p := range valley {
		if p.time == curTime {
			valley[Point{p.x, p.y, curTime + 1}] = make([]Blizzard, 0)
		}
	}

	// place blizzards
	for p, blizzards := range valley {
		if p.time == curTime {
			for _, bliz := range blizzards {
				blizMod := getDirectionModifiers(bliz.direction)
				newBlizPoint := Point{p.x + blizMod.x, p.y + blizMod.y, curTime + 1}

				_, valid := valley[newBlizPoint]
				if valid {
					valley[newBlizPoint] = append(valley[newBlizPoint], bliz)
				} else {
					// pacman effect
					valley[pacmanEffect(newBlizPoint, blizMod)] = append(valley[pacmanEffect(newBlizPoint, blizMod)], bliz)
				}
			}
		}

	}
}

// returns the modifiers to reach "dir" direction
func getDirectionModifiers(dir rune) Point {
	switch dir {
	case '^':
		return dirModifiers[0]
	case '>':
		return dirModifiers[1]
	case 'v':
		return dirModifiers[2]
	case '<':
		return dirModifiers[3]
	}
	fmt.Println("err")
	return Point{0, 0, 0}
}

// returns the point of the blizzard applying the pacman effect
func pacmanEffect(p, mod Point) Point {
	for true {
		newP := Point{p.x - mod.x, p.y - mod.y, p.time}
		_, valid := valley[newP]
		if !valid {
			return p
		}
		p = newP
	}
	fmt.Println("err")
	return Point{0, 0, 0}
}

// modifies stdout printint the valley at "time" time
func printValley(time int) {
	for y := 0; y < 6; y++ {
		for x := 0; x < 8; x++ {
			bliz, found := valley[Point{x, y, time}]
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

func depthFirstSearch(start Point) map[Point]int {
	queue := queue{nil}
	distances := make(map[Point]int)
	distances[start] = 0
	reached := make(map[Point]bool)
	reached[start] = true

	queue.enqueue(start)

	for !queue.isEmpty() {
		u := queue.dequeue()

		reach := reachable(u)
		for _, v := range reach {
			if !reached[v] {
				distances[v] = distances[u] + 1
				reached[v] = true
				queue.enqueue(v)
			}
		}

	}
	return distances
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
