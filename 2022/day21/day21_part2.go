// https://adventofcode.com/2022/day/21
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	name        string
	isImmediate bool
	immediate   float64
	operation   rune
	factor1     string
	factor2     string
}

type BranchVal struct {
	value float64
	valid bool
}

var branches map[string]BranchVal

func main() {
	branches = make(map[string]BranchVal)
	monkeys, rootIndex := parseInput()
	branches["humn"] = BranchVal{0, false}

	calcHumn(monkeys, monkeys[rootIndex])

	fmt.Println(branches["humn"].value)
}

func parseInput() (monkeys []Monkey, rootIndex int) {
	scanner := bufio.NewScanner(os.Stdin)
	var index int
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		var name, factor1, factor2 string
		var isImmediate bool
		var immediate float64
		var operation rune

		name = tokens[0][:4]
		if len(tokens) == 2 { // immediate monkey
			isImmediate = true
			immediate, _ = strconv.ParseFloat(tokens[1], 64)
			branches[name] = BranchVal{immediate, true}
		} else { // operation monkey
			factor1 = tokens[1]
			factor2 = tokens[3]
			operation = rune(tokens[2][0])
		}

		if name == "root" {
			rootIndex = index
			operation = '='
		}

		m := Monkey{name, isImmediate, immediate, operation, factor1, factor2}
		monkeys = append(monkeys, m)

		index++
	}
	return monkeys, rootIndex
}

func calculateValidBranches(monkeys []Monkey, m Monkey) (float64, bool) {
	if m.factor1 == "humn" {
		return 0, false
	}
	if m.factor2 == "humn" {
		return 0, false
	}

	if m.isImmediate {
		return m.immediate, true
	}

	i1 := findMonkey(monkeys, m.factor1)
	i2 := findMonkey(monkeys, m.factor2)
	res1, valid1 := calculateValidBranches(monkeys, monkeys[i1])
	res2, valid2 := calculateValidBranches(monkeys, monkeys[i2])

	if !valid1 {
		branches[m.name] = BranchVal{res2, false}
		return res2, false
	}
	if !valid2 {
		branches[m.name] = BranchVal{res1, false}
		return res1, false
	}

	switch m.operation {
	case '+':
		branches[m.name] = BranchVal{res1 + res2, true}
		return res1 + res2, true
	case '-':
		branches[m.name] = BranchVal{res1 - res2, true}
		return res1 - res2, true
	case '*':
		branches[m.name] = BranchVal{res1 * res2, true}
		return res1 * res2, true
	case '/':
		branches[m.name] = BranchVal{res1 / res2, true}
		return res1 / res2, true
	}
	return 0, true
}

func findMonkey(monkeys []Monkey, name string) int {
	for i, m := range monkeys {
		if m.name == name {
			return i
		}
	}
	return -1
}

func calcHumn(monkeys []Monkey, root Monkey) {

	var rootVal float64
	var validBranchVal float64
	var branchToDet string

	if !branches[root.factor1].valid && !branches[root.factor2].valid {
		calculateValidBranches(monkeys, monkeys[findMonkey(monkeys, root.factor1)])
		calculateValidBranches(monkeys, monkeys[findMonkey(monkeys, root.factor2)])
	}

	if branches[root.factor1].valid {
		validBranchVal = branches[root.factor1].value
		branchToDet = root.factor2
	}
	if branches[root.factor2].valid {
		validBranchVal = branches[root.factor2].value
		branchToDet = root.factor1
	}

	rootVal = branches[root.name].value

	switch root.operation {
	case '=':
		branches[branchToDet] = BranchVal{validBranchVal, true}

	case '+':
		branches[branchToDet] = BranchVal{rootVal - validBranchVal, true}

	case '-':
		branches[branchToDet] = BranchVal{rootVal + validBranchVal, true}

	case '*':
		branches[branchToDet] = BranchVal{rootVal / validBranchVal, true}

	case '/':
		branches[branchToDet] = BranchVal{rootVal * validBranchVal, true}

	}

	fmt.Println(branches[branchToDet].value, string(root.operation), validBranchVal, "=", branches[root.name])

	next := monkeys[findMonkey(monkeys, branchToDet)]
	if !next.isImmediate {
		calcHumn(monkeys, next)
	}

}
