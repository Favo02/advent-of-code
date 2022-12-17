// https://adventofcode.com/2022/day/17
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

var rocks [5]string = [5]string{"####", ".#.\n###\n.#.", "..#\n..#\n###", "#\n#\n#\n#", "##\n##"}

var cave map[Point]bool = make(map[Point]bool)
var fixed map[Point]bool = make(map[Point]bool)

const WALLEFT, WALLRIGHT int = 0, 8

var turn, gasIndex int

var gas string

func main() {
	gas = parseInput()

	res := spawnRocks(1_000_000_000_000)
	fmt.Println(res)

}

func reset() {
	cave = make(map[Point]bool)
	fixed = make(map[Point]bool)
	turn = 0
	gasIndex = 0
	statuses = make(map[State]Payload)
}

func parseInput() (gas string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	gas = line
	return gas
}

// place rock on grid
func spawnRock(t int) {
	rock := rocks[t]

	max := -highestRock()

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

func highestRock() int {
	min := 0
	for rock := range cave {
		if rock.y < min {
			min = rock.y
		}
	}
	return -(min)
}

func spawnRocks(nRocks int) int {
	tot := 1_000_000_000_000
	var foundCycle bool
	var cycleHeight, steps int

	for rocksSpawned := 0; rocksSpawned <= nRocks; rocksSpawned++ {
		spawnRock(rocksSpawned % 5)
		turn = 0
		fallRock()

		curState := getCurrentState(rocksSpawned%5 + 1)

		oldStatus, found := statuses[curState]
		if found && !foundCycle {
			foundCycle = true
			fmt.Println("---found cycle---")

			cycle := rocksSpawned - oldStatus.rocks // number of rocks in the cycle
			cycleHeight = highestRock() - oldStatus.height

			steps = (tot - oldStatus.rocks) / cycle // steps needed to reach tot rocks

			// useless - only debug
			stepsRim := (tot - oldStatus.rocks) % cycle // remaining of steps
			fmt.Println("rim:", stepsRim)

			rocksSpawned = oldStatus.rocks + (cycle * steps)

		} else {
			statuses[curState] = Payload{highestRock(), rocksSpawned}
		}
	}
	return (highestRock() + ((steps - 1) * cycleHeight)) - 1
}

func fallRock() {
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

		// fmt.Println("x:", xMult, "y:", yMult)

		var currentRock, toMove []Point
		for k := range cave {
			// skip already fixed points
			if fixed[k] {
				continue
			}

			currentRock = append(currentRock, k)
			toMove = append(toMove, Point{k.x + xMult, k.y + yMult})
		}

		// fmt.Println("cur", currentRock)
		// fmt.Println("tom", toMove)

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
}

func visualizeMap() {
	for i := -highestRock(); i < 0; i++ {
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

var statuses map[State]Payload = make(map[State]Payload)

type Payload struct {
	height int
	rocks  int
}

type State struct {
	nextRockIndex int
	nextGasIndex  int
	top10rocks    string
}

func getCurrentState(rock int) State {
	curState := State{rock, gasIndex % len(gas), getTopRocks()}
	return curState
}

func getTopRocks() (res string) {
	topRockIndex := -highestRock()
	for i := topRockIndex; i < topRockIndex+20; i++ {
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
