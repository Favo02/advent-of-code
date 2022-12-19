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
}

type State struct {
	oreRobots, clayRobots, obsidianRobots, geodeRobots int
	ore, clay, obsidian, geode                         int
}

func main() {
	bps := parseInput()
	res := initializeResources(len(bps))

	var maxGeodes []State

	// findGeodeRobot(res[0], bps[0], 24)
	// maxGeodes = append(maxGeodes, maxGeodeStatus)
	// maxGeodeStatus.geode = 0 // reset

	// fmt.Println("---")

	findGeodeRobot(res[1], bps[1], 24)
	maxGeodes = append(maxGeodes, maxGeodeStatus)
	maxGeodeStatus.geode = 0 // reset

	fmt.Println(maxGeodes)
}

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
		blueprint := Blueprint{ore, clay, obs1, obs2, geo1, geo2}
		blueprints = append(blueprints, blueprint)
	}
	return blueprints
}

func initializeResources(length int) (res []State) {
	for i := 0; i < length; i++ {
		res = append(res, State{1, 0, 0, 0, 1, 0, 0, 0})
	}
	return res
}

var maxGeodeStatus State
var maxGeodeStatusDay int = 24

func findGeodeRobot(cur State, blue Blueprint, day int) {
	if day <= 0 {
		return
	}

	if cur.geode > maxGeodeStatus.geode {
		fmt.Println(day, cur)
		maxGeodeStatus = cur
		maxGeodeStatusDay = day
	}

	possibilities := generatedPossibleOperations(cur, blue)
	for i := len(possibilities) - 1; i >= 0; i-- {
		poss := possibilities[i]

		if poss {
			// calculate (no edit!) how many robots to create
			robotsToAdd := precalculateTrasfrormState(cur, blue, i)
			// update resources edit robots precalculated
			findGeodeRobot(trasformState(updateResources(cur), blue, robotsToAdd), blue, day-1)
		}
	}
	findGeodeRobot(updateResources(cur), blue, day-1)

}

func generatedPossibleOperations(cur State, b Blueprint) []bool {
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

// returns the current resources adding what robots have collected
func updateResources(r State) State {
	return State{
		r.oreRobots,
		r.clayRobots,
		r.obsidianRobots,
		r.geodeRobots,

		r.ore + r.oreRobots,
		r.clay + r.clayRobots,
		r.obsidian + r.obsidianRobots,
		r.geode + r.geodeRobots,
	}
}

func precalculateTrasfrormState(r State, b Blueprint, index int) []int {
	var or, cl, ob, ge int
	switch index {
	case 0: // ore
		or = r.ore / b.oreCostOre
	case 1: // clay
		cl = r.ore / b.clayCostOre
	case 2: // obsidian
		ob1 := r.ore / b.obsidianCostOre
		ob2 := r.clay / b.obsidianCostClay
		ob = min(ob1, ob2)
	case 3: // geode
		ge1 := r.ore / b.geodeCostOre
		ge2 := r.obsidian / b.geodeCostObsidian
		ge = min(ge1, ge2)
	}
	return []int{or, cl, ob, ge}
}

func trasformState(r State, b Blueprint, pretrasform []int) State {

	// new robots precalculated
	ore := pretrasform[0]
	clay := pretrasform[1]
	obsidian := pretrasform[2]
	geode := pretrasform[3]

	subOre := (ore * b.oreCostOre) + (clay * b.clayCostOre) + (obsidian * b.obsidianCostOre) + (geode * b.geodeCostOre)
	subClay := obsidian * b.obsidianCostClay
	subObsidian := geode * b.geodeCostObsidian

	return State{
		r.oreRobots + ore,
		r.clayRobots + clay,
		r.obsidianRobots + obsidian,
		r.geodeRobots + geode,

		r.ore - subOre,
		r.clay - subClay,
		r.obsidian - subObsidian,
		r.geode,
	}
}

// returns min value between a and b, a if equals
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
