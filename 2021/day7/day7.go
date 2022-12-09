package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	crabs, maxCrab, minCrab := parseInput()

	part1, part2 := calculateLessFuel(crabs, maxCrab, minCrab)
	fmt.Println("constant fuel cost (part1):", part1)
	fmt.Println("increasing fuel cost (part2):", part2)
}

// returns crabs list, max crab and min crab
func parseInput() ([]int, int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	var crabsStr []string
	for scanner.Scan() {
		crabsStr = strings.Split(scanner.Text(), ",") // list of crab (str)
	}

	// convert crabs from str to int
	var crabs []int
	// while converting every crab, save max and min
	var maxCrab int = math.MinInt
	var minCrab int = math.MaxInt

	for _, crabStr := range crabsStr {
		crab, _ := strconv.Atoi(crabStr)
		crabs = append(crabs, crab)
		if crab > maxCrab {
			maxCrab = crab
		}
		if crab < minCrab {
			minCrab = crab
		}
	}

	return crabs, maxCrab, minCrab
}

// returns the least fuel possible to align every crab, both usign constant fuel cost (part1) and increasing fuel cost (part2)
func calculateLessFuel(crabs []int, max, min int) (int, int) {
	var minFuelConst int = math.MaxInt
	var minFuelIncr int = math.MaxInt

	// scan every possible point to align crabs
	for goal := min; goal <= max; goal++ {

		// calculate fuel needed to move all crabs
		var fuelConst, fuelIncr int
		for _, crab := range crabs {
			fuelConst += int(math.Abs(float64(goal - crab)))
			fuelIncr += calculateFuelSteps(int(math.Abs(float64(goal - crab))))
		}

		// save min fuel cost possible
		if fuelConst < minFuelConst {
			minFuelConst = fuelConst
		}
		if fuelIncr < minFuelIncr {
			minFuelIncr = fuelIncr
		}
		fuelConst, fuelIncr = 0, 0
	}

	return minFuelConst, minFuelIncr
}

// returns fuel increased fuel cost (for part2)
func calculateFuelSteps(diff int) int {
	return ((diff * (diff - 1)) / 2) + diff
}
