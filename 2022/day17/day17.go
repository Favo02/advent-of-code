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

func main() {
	gas := parseInput()

	heightPart1, heightPart2 := spawnRocks(gas, 1_000_000_000_000)
	fmt.Println("tower height at", PART1, "rocks spawned:\n\t", heightPart1)
	fmt.Println("tower height at", PART2, "rocks spawned:\n\t", heightPart2)
}

// returns the gas flows read from stdin
func parseInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	gas := scanner.Text()
	return gas
}

// modifies cave placing nRocks rocks,
// returns the cave height at PART1 time and at PART2 time
func spawnRocks(gas string, nRocks int) (int, int) {

	// cave: true = rock fixed, cave: false = rock not fixed
	cave := make(map[Point]bool)

	var gasIndex int // index of gas

	var part1res int // result at PART1 rocks spawned

	// cycle
	var foundCycle bool
	var cycleHeight, steps int
	states := make(map[State]StateValue) // states already encountered

	for nRocksSpawned := 0; nRocksSpawned <= nRocks; nRocksSpawned++ {
		// save current state
		curState := getCurrentState(cave, gas, gasIndex, nRocksSpawned%5)

		// save result for part1
		if nRocksSpawned == PART1 {
			part1res = highestRock(cave)
		}

		spawnRock(cave, nRocksSpawned%5)         // spawn new rock
		gasIndex = fallRock(cave, gas, gasIndex) // let it fall (and update gas index)

		// skip cycle check if before part1 result
		if nRocksSpawned < PART1 {
			continue
		}

		// check for cycle
		oldStatus, found := states[curState]
		if found && !foundCycle { // cycle found
			foundCycle = true

			cycle := nRocksSpawned - oldStatus.rocks           // number of rocks in the cycle
			cycleHeight = highestRock(cave) - oldStatus.height // height of a single cycle

			steps = (PART2 - oldStatus.rocks) / cycle // steps needed to go the closest possible to PART2 rocks

			nRocksSpawned = oldStatus.rocks + (cycle * steps) // skip all rocks spawned it cycle updating loop counter

		} else {
			// save state
			states[curState] = StateValue{highestRock(cave), nRocksSpawned}
		}
	}
	return part1res, (highestRock(cave) + ((steps - 1) * cycleHeight)) - 1
}

// modifies the cave, placing a rock (of shape of index t) on grid (at top of cave)
func spawnRock(cave map[Point]bool, t int) {
	rock := ROCKS[t] // rock shape to spawn

	max := -highestRock(cave) // height of highest rock already in cave

	spawn := Point{WALLEFT + 3, max - 4} // spawn point

	rockHeight := strings.Split(rock, "\n") // height of new rock
	spawn.y -= (len(rockHeight) - 1)        // update spawn point with rock height

	for _, rockLine := range rockHeight { // scan each rock line
		x := spawn.x
		for _, rockPoint := range rockLine { // scan each line point
			if rockPoint == '#' { // place only if rock point is #
				cave[Point{x, spawn.y}] = false
			}
			x++
		}
		spawn.y++
	}
}

// returns the highest point of a rock
func highestRock(cave map[Point]bool) int {
	min := 0 // highest rock has the lowest y
	for rock := range cave {
		if rock.y < min {
			min = rock.y
		}
	}
	return -(min)
}

// modifies cave, letting the rocks fall to definite position
func fallRock(cave map[Point]bool, gas string, gasIndex int) int {
	var gasTurn bool = true // alternate gravity and gas effects

	var isFixed bool // is rock overlapping another rock

	for !isFixed { // keep falling until overlap

		// calculate movement direction
		var xMult, yMult int // x and y movement

		if gasTurn { // gas turn
			yMult = 0 // no vertical movement

			gasDir := gas[gasIndex%len(gas)]
			if gasDir == '<' {
				xMult = -1 // to left movement
			} else if gasDir == '>' {
				xMult = 1 // to right movement
			}

			gasIndex++

		} else { // gravity turn
			xMult = 0  // no horizontal movement
			yMult = +1 // to bottom movement
		}

		// move rock
		var currentRock, toMove []Point
		for k, fixed := range cave { // scan each rock point
			// skip already fixed points
			if fixed {
				continue
			}

			currentRock = append(currentRock, k)
			toMove = append(toMove, Point{k.x + xMult, k.y + yMult}) // move to movement
		}

		var overlap, invalid bool // new rock is overlapping or invalid position

		// check if any new rock point is out of cave range or overlaps an existing rock
		for _, r := range toMove {
			if cave[r] { // overlap a rock
				overlap = true
				break
			}
			if r.y >= 0 { // floor (for 1st rock)
				overlap = true
				break
			}
			if !(r.x > WALLEFT && r.x < WALLRIGHT) { // out of cave range
				invalid = true
			}
		}

		if overlap && !gasTurn { // overlap going down: rock is fixed
			for _, cr := range currentRock {
				cave[cr] = true
			}
			isFixed = true
		}

		if !overlap && !invalid { // new rock is valid
			for _, r := range currentRock { // remove current position
				delete(cave, r)
			}
			for _, r := range toMove { // add new position
				cave[r] = false
			}
		}

		gasTurn = !gasTurn // alternate turn
		// visualizeMap(cave)
	}
	return gasIndex
}

// modifies stdout printing a visualization of the cave
func visualizeMap(cave map[Point]bool) {
	for i := -highestRock(cave); i < 0; i++ { // start from highest rock
		for j := WALLEFT; j < WALLRIGHT; j++ { // scan cave width
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

	// generate top 40 rocks state
	var topRocks string
	topRockIndex := -highestRock(cave)
	for i := topRockIndex; i < topRockIndex+40; i++ {
		for j := 1; j <= 7; j++ {
			if cave[Point{j, i}] {
				topRocks += "#"
			} else {
				topRocks += "."
			}
		}
	}

	// create state object
	return State{rock, gasIndex % len(gas), topRocks}
}
