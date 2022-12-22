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
	oreCost       int
	clayCost      int
	obsidianCost1 int
	obsidianCost2 int
	geodeCost1    int
	geodeCost2    int
}

type Resources struct {
	oreR, clayR, obsidianR, geodeR int
	ore, clay, obsidian, geode     int
}

func main() {
	bps := parseInput()
	res := initializeResources(bps)

	max := maxGeodes(25, res[bps[0]], bps[0])
	fmt.Println(max)
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

func initializeResources(bps []Blueprint) map[Blueprint]Resources {
	res := make(map[Blueprint]Resources)
	for _, b := range bps {
		res[b] = Resources{1, 0, 0, 0, 0, 0, 0, 0}
	}
	return res
}

func maxGeodes(time int, res Resources, bp Blueprint) Resources {
	fmt.Println(time, res, bp)

	if time <= 0 {
		return res
	}

	// modifies the resources
	res = updateResources(res)

	// get what is possible to create with current resource
	poss := generatedPossibleOperations(res, bp)
	fmt.Println(poss)

	var globalMax Resources
	var maxRes Resources

	if checkAllZero(poss) {
		maxRes = maxGeodes(time-1, res, bp)
		if maxRes.geode > globalMax.geode {
			globalMax = maxRes
		}
	} else {
		for i := range poss {
			// update resource for what poss selected
			res = removeResources(res, bp, i)
			// pass updated res to maxGeodes
			// pass time decreased by 1
			maxRes = maxGeodes(time-1, res, bp)
			// if geodes generated max then save it
			if maxRes.geode > globalMax.geode {
				globalMax = maxRes
			}
		}
	}

	return globalMax
}

// returns the current resources adding what robots have collected
func updateResources(r Resources) Resources {
	return Resources{r.oreR, r.clayR, r.obsidianR, r.geodeR, r.ore + r.oreR, r.clay + r.clayR, r.obsidian + r.obsidianR, r.geode + r.geodeR}
}

func generatedPossibleOperations(r Resources, b Blueprint) []int {
	or := r.ore / b.oreCost

	cl := r.ore / b.clayCost

	ob1 := r.ore / b.obsidianCost1
	ob2 := r.clay / b.obsidianCost2
	ob := min(ob1, ob2)

	ge1 := r.ore / b.geodeCost1
	ge2 := r.obsidian / b.geodeCost2
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

// remove resources for "selected" robot created
func removeResources(r Resources, b Blueprint, selected int) Resources {
	var or, cl, ob, ge int
	var subor, subcl, subob int
	switch selected {
	case 0: // ore
		or = r.ore / b.oreCost
		subor = or * b.oreCost
	case 1: // clay
		cl = r.ore / b.clayCost
		subor = cl * b.clayCost
	case 2: // obsidian
		ob1 := r.ore / b.obsidianCost1
		ob2 := r.clay / b.obsidianCost2
		ob = min(ob1, ob2)
		subor = ob * b.obsidianCost1
		subcl = ob * b.obsidianCost2
	case 3: // geode
		ge1 := r.ore / b.geodeCost1
		ge2 := r.obsidian / b.geodeCost2
		ge = min(ge1, ge2)
		subor = ge * b.geodeCost1
		subob = ge * b.geodeCost2
	}
	return Resources{r.oreR + or, r.clayR + cl, r.obsidianR + ob, r.geodeR + ge, r.ore - subor, r.clay - subcl, r.obsidian - subob, r.geode}
}

func checkAllZero(sl []int) bool {
	for _, v := range sl {
		if v != 0 {
			return false
		}
	}
	return true
}

func fastestWay() {

}
