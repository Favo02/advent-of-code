// https://adventofcode.com/2022/day/17
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// point in the cave
type Point struct {
	x, y int
}

// state at a current time
type State struct {
	nextRockIndex int
	nextGasIndex  int
	topRocks      string
}

// values at a current time
type StateValue struct {
	height int
	rocks  int
}

// rocks shapes
var ROCKS [5]string = [5]string{"####", ".#.\n###\n.#.", "..#\n..#\n###", "#\n#\n#\n#", "##\n##"}

// walls of cave
const WALLEFT int = 0
const WALLRIGHT int = 8

// part1 and part2 rocks number
const PART1 int = 2022
const PART2 int = 1_000_000_000_000

// index of gas affecting the rocks

func main() {
	gas := parseInput()

	p1, p2 := spawnRocks(gas, 1_000_000_000_000)
	fmt.Println(p1, p2)
}

// returns the gas flows read from stdin
func parseInput() (gas string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	gas = line
	return gas
}

// modifies cave placing nRocks rocks,
// returns the cave height at PART1 time and at PART2 time
func spawnRocks(gas string, nRocks int) (int, int) {

	cave := make(map[Point]bool)
	fixed := make(map[Point]bool)

	var part1res int
	var foundCycle bool
	var cycleHeight, steps int
	var gasIndex int

	statuses := make(map[State]StateValue)

	for rocksSpawned := 0; rocksSpawned <= nRocks; rocksSpawned++ {
		if rocksSpawned == PART1 {
			part1res = highestRock(cave)
		}

		spawnRock(cave, rocksSpawned%5)
		gasIndex = fallRock(cave, fixed, gas, gasIndex)

		curState := getCurrentState(cave, gas, gasIndex, rocksSpawned%5+1)

		oldStatus, found := statuses[curState]
		if rocksSpawned >= PART1 && found && !foundCycle {
			foundCycle = true

			cycle := rocksSpawned - oldStatus.rocks // number of rocks in the cycle
			cycleHeight = highestRock(cave) - oldStatus.height

			steps = (PART2 - oldStatus.rocks) / cycle // steps needed to reach tot rocks

			rocksSpawned = oldStatus.rocks + (cycle * steps)

		} else {
			statuses[curState] = StateValue{highestRock(cave), rocksSpawned}
		}
	}
	return part1res, (highestRock(cave) + ((steps - 1) * cycleHeight)) - 1
}

// modifies the cave, placing a rock (of shape of index t) on grid (at top of cave)
func spawnRock(cave map[Point]bool, t int) {
	rock := ROCKS[t]

	max := -highestRock(cave)

	spawn := Point{WALLEFT + 3, max - 4}

	levels := strings.Split(rock, "\n")
	spawn.y -= (len(levels) - 1)

	for _, lev := range levels {
		x := spawn.x
		for _, ch := range lev {
			if ch == '#' {
				cave[Point{x, spawn.y}] = true
			}
			x++
		}
		spawn.y++
	}
}

// returns the highest point of a rock
func highestRock(cave map[Point]bool) int {
	min := 0
	for rock := range cave {
		if rock.y < min {
			min = rock.y
		}
	}
	return -(min)
}

// modifies cave, letting the rocks fall to definite position
func fallRock(cave, fixed map[Point]bool, gas string, gasIndex int) int {
	var turn int = 0
	var overlap bool
	for !overlap {

		var xMult, yMult int

		if turn%2 == 0 { // GAS
			yMult = 0

			gasDir := gas[gasIndex%len(gas)]
			if gasDir == '<' {
				xMult = -1
			} else if gasDir == '>' {
				xMult = 1
			}

			gasIndex++

		} else { // GRAVITY
			xMult = 0
			yMult = +1
		}

		var currentRock, toMove []Point
		for k := range cave {
			// skip already fixed points
			if fixed[k] {
				continue
			}

			currentRock = append(currentRock, k)
			toMove = append(toMove, Point{k.x + xMult, k.y + yMult})
		}

		var invalid bool
		// check if any point of rock overlaps an existing rock
		for _, r := range toMove {
			if fixed[r] {
				overlap = true
				break
			}
			if r.y >= 0 {
				overlap = true
				break
			}
			if !(r.x > WALLEFT && r.x < WALLRIGHT) {
				invalid = true
			}
		}

		if !overlap {
			if !invalid {
				for _, r := range currentRock {
					delete(cave, r)
				}
				for _, nr := range toMove {
					cave[nr] = true
				}
			}
		} else {
			if turn%2 == 1 {
				for _, cr := range currentRock {
					fixed[cr] = true
				}
			} else {
				overlap = false

			}
		}

		turn++
		// visualizeMap()
	}
	return gasIndex
}

// modifies stdout printing a visualization of the cave
func visualizeMap(cave map[Point]bool) {
	for i := -highestRock(cave); i < 0; i++ {
		for j := 1; j < 8; j++ {
			if cave[Point{j, i}] {
				fmt.Print("█")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// returns the current state (generated on 40 top rocks levels)
func getCurrentState(cave map[Point]bool, gas string, gasIndex, rock int) State {
	curState := State{rock, gasIndex % len(gas), getTopRocks(cave)}
	return curState
}

// returns the top 40 rocks
func getTopRocks(cave map[Point]bool) (res string) {
	topRockIndex := -highestRock(cave)
	for i := topRockIndex; i < topRockIndex+40; i++ {
		for j := 1; j <= 7; j++ {
			if cave[Point{j, i}] {
				res += "#"
			} else {
				res += "."
			}
		}
	}
	return res
}
