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
	time int
}

type Blizzard struct {
	direction rune
}

var timeGenerated int

var dirModifiers []Point = []Point{{0, -1, +1}, {+1, 0, +1}, {0, +1, +1}, {-1, 0, +1}}

func main() {
	valley, end := parseInput()
	start := Point{1, 0, 0}

	startToEnd := distanceInTime(valley, start, end)

	fmt.Println("start to end (part1):\n\t", startToEnd)

	endToStart := distanceInTime(valley, Point{end.x, end.y, startToEnd}, start)
	startToEnd2 := distanceInTime(valley, Point{start.x, start.y, startToEnd + endToStart}, end)
	part2 := startToEnd + endToStart + startToEnd2

	fmt.Println("start to end to start to end (part2):\n\t", part2)
}

// modifies valley placing the blizzard parsed from stdin
// modifies stdin
func parseInput() (map[Point][]Blizzard, Point) {
	valley := make(map[Point][]Blizzard)
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	lastPointX := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, v := range line {
			if v == '.' {
				valley[Point{x, y, 0}] = make([]Blizzard, 0)
				lastPointX = x
			}
			if v == '>' || v == '<' || v == '^' || v == 'v' {
				valley[Point{x, y, 0}] = make([]Blizzard, 1)
				valley[Point{x, y, 0}][0] = Blizzard{v}
			}
		}
		y++
	}
	return valley, Point{lastPointX, y - 1, 0}
}

// returns the minimum distance from start to end
// modifies valley (generating points in time)
func distanceInTime(valley map[Point][]Blizzard, start, end Point) int {
	dist := depthFirstSearch(valley, start, end) // distances from start to each point until start found
	for p, d := range dist {                     // search distance to end (unknown time)
		if p.x == end.x && p.y == end.y {
			return d
		}
	}
	return -1
}

// returns the distance from start to each point in valley (until end is found)
// modifies valley (generating points in time)
func depthFirstSearch(valley map[Point][]Blizzard, start Point, end Point) map[Point]int {
	queue := queue{nil}
	distances := make(map[Point]int)
	distances[start] = 0
	reached := make(map[Point]bool)
	reached[start] = true

	queue.enqueue(start)

	for !queue.isEmpty() {
		u := queue.dequeue()

		reach := reachable(valley, u) // points reachable from u
		for _, v := range reach {
			if !reached[v] {
				distances[v] = distances[u] + 1
				reached[v] = true
				queue.enqueue(v)
			}

			// end found, stop
			if v.x == end.x && v.y == end.y {
				return distances
			}
		}

	}
	return distances
}

// returns the points reachable from "u"
// modifies valley (generating points in time)
func reachable(valley map[Point][]Blizzard, u Point) (reac []Point) {

	// generate points at next time if not generated yet
	if timeGenerated <= u.time {
		generateNextMinute(valley, u.time)
		timeGenerated++
	}

	// scan each point reachable from current point (cur) using every direction modifier
	for _, dirMod := range dirModifiers {

		// point reached from cur
		v := Point{u.x + dirMod.x, u.y + dirMod.y, u.time + 1}

		blizzards, found := valley[v]

		// no blizzards on that point
		if found && len(blizzards) == 0 {
			reac = append(reac, v)
		}
	}

	// check if current point still safe (no blizzards)
	blizzards, found := valley[Point{u.x, u.y, u.time + 1}]
	if found && len(blizzards) == 0 {
		reac = append(reac, Point{u.x, u.y, u.time + 1})
	}

	return reac
}

// modifies valley generating the position of blizzards at next minute
func generateNextMinute(valley map[Point][]Blizzard, curTime int) {

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
				blizMod := getDirectionModifiers(bliz.direction) // get modifiers to reach next blizzard position
				newBlizPoint := Point{p.x + blizMod.x, p.y + blizMod.y, curTime + 1}

				_, valid := valley[newBlizPoint]
				if valid {
					valley[newBlizPoint] = append(valley[newBlizPoint], bliz)
				} else {
					// pacman effect
					pacman := pacmanEffect(valley, newBlizPoint, blizMod)
					valley[pacman] = append(valley[pacman], bliz)
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
func pacmanEffect(valley map[Point][]Blizzard, p, mod Point) Point {
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
