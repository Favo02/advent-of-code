// https://adventofcode.com/2022/day/14

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

var sandStart Point = Point{500, 0}

func main() {
	rocks := parseInput()
	cave1 := fillRocks(rocks) // fill lines between rocks
	cave2 := copyMap(cave1)   // duplicate cave for part2

	cave1 = simulateSand(cave1, false) // simulate falling sand (without floor)
	cave2 = simulateSand(cave2, true)  // simulate falling sand (with floor)

	sand1 := countSand(cave1) // count sand stuck in cave
	sand2 := countSand(cave2) // count sand stuck in cave

	fmt.Println("sand standing on rocks (part1):\n\t", sand1)
	fmt.Println("sand standing on floor (part2):\n\t", sand2)
}

func parseInput() (rocks [][]Point) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " -> ") // get rocks of line

		// parse rocks ([]string) to []Point
		var rock []Point
		for _, t := range tokens { // scan each rock
			xy := strings.Split(t, ",")      // split x and y
			x, _ := strconv.Atoi(xy[0])      // parse x
			y, _ := strconv.Atoi(xy[1])      // parse y
			rock = append(rock, Point{x, y}) // add rock Point
		}
		rocks = append(rocks, rock) // add rocks line to rocks
	}
	return rocks
}

func fillRocks(rocks [][]Point) map[Point]int {
	cave := make(map[Point]int)
	for _, rock := range rocks { // scan each rocks line
		for i := 0; i < len(rock)-1; i++ { // scan each rock
			drawLine(cave, rock[i], rock[i+1]) // draw rock line in map between two rocks
		}
	}
	return cave
}

func drawLine(cave map[Point]int, source, destination Point) map[Point]int {
	diffX := source.x - destination.x
	diffY := source.y - destination.y

	// scan each point between rocks
	if diffX > 0 {
		for i := 0; i <= diffX; i++ {
			cave[Point{source.x - i, source.y}] = 1
		}
	} else {
		for i := 0; i <= -diffX; i++ {
			cave[Point{source.x + i, source.y}] = 1
		}
	}

	// scan each point between rocks
	if diffY > 0 {
		for i := 0; i <= diffY; i++ {
			cave[Point{source.x, source.y - i}] = 1
		}
	} else {
		for i := 0; i <= -diffY; i++ {
			cave[Point{source.x, source.y + i}] = 1
		}
	}

	return cave
}

func simulateSand(cave map[Point]int, floor bool) map[Point]int {
	var oldSand int
	for true {
		// sand spawn blocked
		if cave[sandStart] == -1 {
			break
		}

		// spawn new sand
		cave[sandStart] = -1

		// let sand spawned fall
		cave = sandFall(cave, sandStart, floor)

		// if sand spawned falled out of map
		if countSand(cave) == oldSand {
			break
		}
		oldSand = countSand(cave) // update number of sand pieces in the cave
	}
	return cave
}

func sandFall(cave map[Point]int, curSand Point, floor bool) map[Point]int {
	lowestRock := lowestRock(cave) // lowest rock
	var rest bool                  // is current sand piece at rest

	for !rest {

		if floor {
			// below lowest rock: stop at floor
			if curSand.y == lowestRock+1 {
				cave[curSand] = -1
				break
			}
		} else {
			// below lowest rock: falling out of the map
			if curSand.y == lowestRock {
				cave[curSand] = 0
				break
			}
		}

		// fall down one step
		if cave[Point{curSand.x, curSand.y + 1}] == 0 {
			cave[Point{curSand.x, curSand.y + 1}] = -1 // crate sand down
			cave[curSand] = 0                          // remove sand on old point
			curSand = Point{curSand.x, curSand.y + 1}  // update current sand
			continue
		}

		// fall one step down and to the left
		if cave[Point{curSand.x - 1, curSand.y + 1}] == 0 {
			cave[Point{curSand.x - 1, curSand.y + 1}] = -1 // create sand down left
			cave[curSand] = 0                              // remove sand on old point
			curSand = Point{curSand.x - 1, curSand.y + 1}  // update current sand
			continue
		}

		// one step down and to the right
		if cave[Point{curSand.x + 1, curSand.y + 1}] == 0 {
			cave[Point{curSand.x + 1, curSand.y + 1}] = -1 // create sand down right
			cave[curSand] = 0                              // remove sand on old point
			curSand = Point{curSand.x + 1, curSand.y + 1}  // update current sand
			continue
		}

		// rest
		rest = true
	}
	return cave
}

func countSand(cave map[Point]int) int {
	var sand int
	for _, p := range cave {
		if p == -1 {
			sand++
		}
	}
	return sand
}

func lowestRock(cave map[Point]int) int {
	low := math.MinInt
	for k, v := range cave {
		if v == 1 && k.y > low {
			low = k.y
		}
	}
	return low
}

func copyMap(cave map[Point]int) map[Point]int {
	newMap := make(map[Point]int)
	for k, v := range cave {
		newMap[k] = v
	}
	return newMap
}
