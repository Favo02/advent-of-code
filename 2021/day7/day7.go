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
	fmt.Println(calculateLessFuel1(crabs, maxCrab, minCrab))
	fmt.Println(calculateLessFuel2(crabs, maxCrab, minCrab))
}

func parseInput() ([]int, int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	var crabsStr []string
	for scanner.Scan() {
		crabsStr = strings.Split(scanner.Text(), ",")
	}

	var crabs []int
	var maxCrab, minCrab int = math.MinInt, math.MaxInt
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

func calculateLessFuel1(crabs []int, max, min int) int {
	var minFuel int = math.MaxInt
	for goal := min; goal <= max; goal++ {
		var fuel int
		for _, crab := range crabs {
			fuel += int(math.Abs(float64(goal - crab)))
		}
		if fuel < minFuel {
			minFuel = fuel
		}
		fuel = 0
	}
	return minFuel
}

func calculateLessFuel2(crabs []int, max, min int) int {
	var minFuel int = math.MaxInt
	for goal := min; goal <= max; goal++ {
		var fuel int
		for _, crab := range crabs {
			fuel += calculateFuelSteps(int(math.Abs(float64(goal - crab))))
		}
		if fuel < minFuel {
			minFuel = fuel
		}
		fuel = 0
	}
	return minFuel
}

func calculateFuelSteps(diff int) int {
	var sum int
	for i := 0; i <= diff; i++ {
		sum += i
	}
	return sum
}
