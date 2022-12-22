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
	immediate   int
	operation   rune
	factor1     string
	factor2     string
}

var root1, root2 int = 1, 0

func main() {
	monkeys, rootIndex := parseInput()

	humnValue := binarySearch(monkeys, rootIndex)

	fmt.Println(humnValue)
}

func parseInput() (monkeys []Monkey, rootIndex int) {
	scanner := bufio.NewScanner(os.Stdin)
	var index int
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		var name, factor1, factor2 string
		var isImmediate bool
		var immediate int
		var operation rune

		name = tokens[0][:4]
		if len(tokens) == 2 { // immediate monkey
			isImmediate = true
			immediate, _ = strconv.Atoi(tokens[1])
		} else { // operation monkey
			factor1 = tokens[1]
			factor2 = tokens[3]
			operation = rune(tokens[2][0])
		}

		if name == "root" {
			rootIndex = index
		}

		m := Monkey{name, isImmediate, immediate, operation, factor1, factor2}
		monkeys = append(monkeys, m)

		index++
	}
	return monkeys, rootIndex
}

func calculateMonkey(monkeys []Monkey, m Monkey) int {
	if m.isImmediate {
		return m.immediate
	}

	var res1, res2 int

	i1 := findMonkey(monkeys, m.factor1)
	i2 := findMonkey(monkeys, m.factor2)
	res1 = calculateMonkey(monkeys, monkeys[i1])
	res2 = calculateMonkey(monkeys, monkeys[i2])

	if m.name == "root" {
		root1 = res1
		root2 = res2
	}

	switch m.operation {
	case '+':
		return res1 + res2
	case '-':
		return res1 - res2
	case '*':
		return res1 * res2
	case '/':
		return res1 / res2
	}
	return 0
}

func findMonkey(monkeys []Monkey, name string) int {
	for i, m := range monkeys {
		if m.name == name {
			return i
		}
	}
	return -1
}

func binarySearch(monkeys []Monkey, rootIndex int) int {
	var min, max int = 0, 10000000000000000

	var found bool
	var test int
	for !found {
		test = (min + max) / 2

		monkeys[findMonkey(monkeys, "humn")] = Monkey{"humn", true, test, 'a', "", ""}
		calculateMonkey(monkeys, monkeys[rootIndex])

		if root1 == root2 {
			found = true
		} else if root1-root2 > 0 {
			min = test + 1
		} else {
			max = test - 1
		}

	}

	return test
}
