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

type Resources struct {
	oreR, clayR, obsidianR, geodeR int
	ore, clay, obsidian, geode     int
}

const DAYS int = 24

// indexes of various rocks
const ORE int = 0
const CLAY int = 1
const OBSIDIAN int = 2
const GEODE int = 3

func main() {
	bps := parseInput()
	res := initializeResources(len(bps))

	fmt.Println(bps, res)
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

func initializeResources(length int) (res []Resources) {
	for i := 0; i < length; i++ {
		res = append(res, Resources{1, 0, 0, 0, 0, 0, 0, 0})
	}
	return res
}

func recursiveBest(days int, blue Blueprint, res Resources) Resources {
	if days == 0 {
		return res
	}

	var currentRes Resources
	var maxRes Resources

	poss := generatedPossibleOperations(res, bp) // possible choice

	for k := range notVisited { // scan every possible generation (with NO generation)
		currentRes = recursiveBest(days-1, blue, res)
		if currentRes.geode > maxRes.geode {
			maxRes = currentRes
		}
	}

	return maxRes
}

func generatedPossibleOperations(r Resources, b Blueprint) []int {
	or := r.ore / b.oreCostOre

	cl := r.ore / b.clayCostOre

	ob1 := r.ore / b.obsidianCostOre
	ob2 := r.clay / b.obsidianCostClay
	ob := min(ob1, ob2)

	ge1 := r.ore / b.geodeCostOre
	ge2 := r.obsidian / b.geodeCostObsidian
	ge := min(ge1, ge2)

	return []int{or, cl, ob, ge}
}

// returns min value between a and b, a if equals
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
