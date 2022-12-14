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

var sandSpawn Point = Point{500, 0}

const ROCK = 1
const SAND = -1
const VOID = 0

func main() {
	rocks := parseInput()
	cave1 := fillRocks(rocks) // fill lines between rocks
	cave2 := copyMap(cave1)   // duplicate cave for part2

	cave1 = fillSand(cave1, false) // simulate falling sand (without floor)
	cave2 = fillSand(cave2, true)  // simulate falling sand (with floor)

	sand1 := countSand(cave1) // count sand stuck in cave
	sand2 := countSand(cave2) // count sand stuck in cave

	fmt.Println("sand standing on rocks (part1):\n\t", sand1)
	fmt.Println("sand standing on floor (part2):\n\t", sand2)
}

// REQUIRES: stdin is a valid challenge input
// MODIFIES: stdin
// EFFECTS: returns each line of rocks
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

// REQUIRES: distance between two rocks on the same line must be one dimensional (only x or only y)
// EFFECTS: returns a map (cave) on which is represented each point that is a rock
func fillRocks(rocks [][]Point) map[Point]int {
	cave := make(map[Point]int)
	for _, rock := range rocks { // scan each rocks line
		for i := 0; i < len(rock)-1; i++ { // scan each rock
			drawLine(cave, rock[i], rock[i+1]) // set as rock on cave each point between two consecutive rocks (on the same line)
		}
	}
	return cave
}

// REQUIRES: distance between "source" and "destination" must be one dimensional (only x or only y)
// EFFECTS: returns cave on which has been set as rocks each point between source and destination
func drawLine(cave map[Point]int, source, destination Point) map[Point]int {
	diffX := source.x - destination.x
	diffY := source.y - destination.y

	// scan each point between rocks
	if diffX > 0 {
		for i := 0; i <= diffX; i++ {
			cave[Point{source.x - i, source.y}] = ROCK
		}
	} else {
		for i := 0; i <= -diffX; i++ {
			cave[Point{source.x + i, source.y}] = ROCK
		}
	}

	// scan each point between rocks
	if diffY > 0 {
		for i := 0; i <= diffY; i++ {
			cave[Point{source.x, source.y - i}] = ROCK
		}
	} else {
		for i := 0; i <= -diffY; i++ {
			cave[Point{source.x, source.y + i}] = ROCK
		}
	}

	return cave
}

// EFFECTS: returns the cave on which have been set as sand the points on which sand is stuck after falling from spawn
func fillSand(cave map[Point]int, floor bool) map[Point]int {
	var oldSand int
	for true {
		// sand spawn blocked
		if cave[sandSpawn] == SAND {
			break
		}

		// spawn new sand
		cave[sandSpawn] = SAND

		// let sand spawned fall
		cave = sandFall(cave, sandSpawn, floor)

		// if sand spawned falled out of cave
		if countSand(cave) == oldSand {
			break
		}
		oldSand = countSand(cave) // update number of sand pieces in the cave
	}
	return cave
}

// EFFECTS: return the cave on which a single point of sand ("curSand") have been moved to his final position
func sandFall(cave map[Point]int, curSand Point, floor bool) map[Point]int {
	lowestRock := lowestRock(cave) // lowest rock
	var rest bool                  // is current sand piece at rest

	for !rest {

		if floor {
			// below lowest rock: stop at floor
			if curSand.y == lowestRock+1 {
				cave[curSand] = SAND
				break
			}
		} else {
			// below lowest rock: falling out of the cave
			if curSand.y == lowestRock {
				cave[curSand] = VOID
				break
			}
		}

		// fall down one step
		if cave[Point{curSand.x, curSand.y + 1}] == VOID {
			cave[Point{curSand.x, curSand.y + 1}] = SAND // crate sand down
			cave[curSand] = VOID                         // remove sand on old point
			curSand = Point{curSand.x, curSand.y + 1}    // update current sand
			continue
		}

		// fall one step down and to the left
		if cave[Point{curSand.x - 1, curSand.y + 1}] == VOID {
			cave[Point{curSand.x - 1, curSand.y + 1}] = SAND // create sand down left
			cave[curSand] = VOID                             // remove sand on old point
			curSand = Point{curSand.x - 1, curSand.y + 1}    // update current sand
			continue
		}

		// one step down and to the right
		if cave[Point{curSand.x + 1, curSand.y + 1}] == VOID {
			cave[Point{curSand.x + 1, curSand.y + 1}] = SAND // create sand down right
			cave[curSand] = VOID                             // remove sand on old point
			curSand = Point{curSand.x + 1, curSand.y + 1}    // update current sand
			continue
		}

		// rest
		rest = true
	}
	return cave
}

// EFFECTS: returns the number of sand points in the cave
func countSand(cave map[Point]int) int {
	var sand int
	for _, p := range cave {
		if p == SAND {
			sand++
		}
	}
	return sand
}

// EFFECTS: return the y position of the lowest rock in the cave
func lowestRock(cave map[Point]int) int {
	low := math.MinInt
	for k, v := range cave {
		if v == ROCK && k.y > low {
			low = k.y
		}
	}
	return low
}

// EFFECTS: returns a new map copy of "oldMap"
func copyMap(oldMap map[Point]int) map[Point]int {
	newMap := make(map[Point]int)
	for k, v := range oldMap {
		newMap[k] = v
	}
	return newMap
}
