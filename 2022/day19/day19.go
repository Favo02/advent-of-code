// https://adventofcode.com/2022/day/19
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Blueprint struct {
	oreCostOre        int
	clayCostOre       int
	obsidianCostOre   int
	obsidianCostClay  int
	geodeCostOre      int
	geodeCostObsidian int

	maxOreCost      int
	maxClayCost     int
	maxObsidianCost int
}

type State struct {
	time                                               int
	oreRobots, clayRobots, obsidianRobots, geodeRobots int
	ore, clay, obsidian, geode                         int
}

// max geodes generated
var maxGeodeStatus State

func main() {
	blueprints := parseInput() // parse blueprints

	var sum int
	for id, blue := range blueprints {
		fmt.Println("---- blueprint", id+1, "----")
		simulateStates(State{23, 1, 0, 0, 0, 2, 0, 0, 0}, blue)
		sum += ((id + 1) * maxGeodeStatus.geode)
		fmt.Println(maxGeodeStatus.geode)
		fmt.Println("---------------")
		maxGeodeStatus = State{}
	}

	fmt.Println(sum)
}

// returns blueprints parsed from stdin
// modifies stdin
func parseInput() (blueprints []Blueprint) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		ore, _ := strconv.Atoi(tokens[6])
		clay, _ := strconv.Atoi(tokens[12])
		obs1, _ := strconv.Atoi(tokens[18])
		obs2, _ := strconv.Atoi(tokens[21])
		geo1, _ := strconv.Atoi(tokens[27])
		geo2, _ := strconv.Atoi(tokens[30])
		blue := Blueprint{ore, clay, obs1, obs2, geo1, geo2, 0, 0, 0}

		// max resources cost
		blue.maxOreCost = max(max(max(blue.oreCostOre, blue.clayCostOre), blue.obsidianCostOre), blue.geodeCostOre)
		blue.maxClayCost = blue.obsidianCostClay
		blue.maxObsidianCost = blue.geodeCostObsidian

		blueprints = append(blueprints, blue)
	}
	return blueprints
}

func simulateStates(cur State, blue Blueprint) {

	// end of time reached, stop
	if cur.time <= 0 {
		return
	}

	// optimization
	if cur.geode < maxGeodeStatus.geode && cur.time <= maxGeodeStatus.time {
		return
	}

	// number of robots of resource > max amount needed of resource
	if cur.oreRobots > blue.maxOreCost {
		return
	}
	if cur.clayRobots > blue.maxClayCost {
		return
	}
	if cur.obsidianRobots > blue.maxObsidianCost {
		return
	}

	// max geodes check
	if cur.geode > maxGeodeStatus.geode {
		maxGeodeStatus = cur
		fmt.Println(cur.time, cur.geode, cur)
	}

	// generate possible operations (robots to craft)
	possibilities := generatePossibleOperations(cur, blue)

	// scan each possible operation
	for indexRobotToAdd, poss := range possibilities {
		if poss {
			robotsToAdd := make([]int, 4)
			robotsToAdd[indexRobotToAdd] = 1

			// update resources edit robots precalculated
			simulateStates(updateResources(cur, indexRobotToAdd, blue), blue)
		}
	}
	simulateStates(updateResources(cur, -1, blue), blue)
}

// returns the possible operations to perform with "cur" resources
// [0] = ore robot
// [1] = clay robot
// [2] = obsidian robot
// [3] = geode robot
func generatePossibleOperations(cur State, b Blueprint) []bool {
	res := []bool{false, false, false, false}

	or := cur.ore / b.oreCostOre
	if or > 0 {
		res[0] = true
	}

	cl := cur.ore / b.clayCostOre
	if cl > 0 {
		res[1] = true
	}

	ob1 := cur.ore / b.obsidianCostOre
	ob2 := cur.clay / b.obsidianCostClay
	if ob1 > 0 && ob2 > 0 {
		res[2] = true
	}

	ge1 := cur.ore / b.geodeCostOre
	ge2 := cur.obsidian / b.geodeCostObsidian
	if ge1 > 0 && ge2 > 0 {
		res[3] = true
	}

	return res
}

// returns the updated resources, increasing time and adding what robots have collected
func updateResources(cur State, indexRobotToAdd int, blue Blueprint) State {
	var oreR, clayR, obsR, geoR int
	var ore, clay, obsidian, geode int

	ore = cur.ore + cur.oreRobots
	clay = cur.clay + cur.clayRobots
	obsidian = cur.obsidian + cur.obsidianRobots
	geode = cur.geode + cur.geodeRobots

	oreR = cur.oreRobots
	clayR = cur.clayRobots
	obsR = cur.obsidianRobots
	geoR = cur.geodeRobots

	switch indexRobotToAdd {
	case -1: // no robots to add

	case 0: // ore robot
		oreR++
		ore -= blue.oreCostOre

	case 1: // clay robot
		clayR++
		ore -= blue.clayCostOre

	case 2: // obsidian robot
		obsR++
		ore -= blue.obsidianCostOre
		clay -= blue.obsidianCostClay

	case 3: // geode robot
		geoR++
		ore -= blue.geodeCostOre
		obsidian -= blue.geodeCostObsidian
	}

	return State{
		cur.time - 1,
		oreR, clayR, obsR, geoR,
		ore, clay, obsidian, geode,
	}
}

// returns the maximum cost of ore
func getMaxOreCost(blue Blueprint) int {
	return max(max(max(blue.oreCostOre, blue.clayCostOre), blue.obsidianCostOre), blue.geodeCostOre)
}

// returns the max value between a and b
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
