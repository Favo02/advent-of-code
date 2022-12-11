package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	modulos        []map[int]int // instead of actual values just save the modulos (of each possible test divider)
	operation      rune          // operation to perform (*, +)
	square         bool          // operation to perform is a square (^2)
	coefficient    int           // number to perform the operation on (except for square)
	test           int           // number to test the modulo
	true_          int           // monkey to move the value if test is valid
	false_         int           // monkey to move the value if test is invalid
	itemsInspected int           // number of items inspected
}

func main() {
	var rounds int = 10000

	monkeys, tempObj, dividers := parseInput()
	monkeys = calulateModulos(monkeys, tempObj, dividers)
	monkeys = simulateRounds(monkeys, dividers, rounds)
	businessLevel := getBusinessLevel(monkeys)

	fmt.Println("business level after", rounds, "rounds:", businessLevel)
}

// REQUIRES: stdin is a valid challenge input
// MODIFIES: stdin
// EFFECTS: returns the monkeys, the objects (monkey items) and the dividers (all dividers of each monkey)
func parseInput() (monkeys []Monkey, tempObj [][]int, dividers []int) {
	scanner := bufio.NewScanner(os.Stdin)

	var curMonkey Monkey
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		// new monkey start: initialize object
		if line[:6] == "Monkey" {
			curMonkey = Monkey{}
		}

		// starting items
		if line[:6] == "  Star" {
			tokens := strings.Split(line[18:], ", ")
			var starting []int
			for _, v := range tokens {
				n, _ := strconv.Atoi(v)
				starting = append(starting, n)
			}

			// save items not in monkeys, but in temporary slice of slice
			// because in monkeys the modulos for all possible dividers are saved and all dividers are still unknown
			tempObj = append(tempObj, starting)
		}

		// operation to perform
		if line[:6] == "  Oper" {
			// special case: square
			if len(line) == 28 && line == "  Operation: new = old * old" {
				curMonkey.square = true
				curMonkey.operation = '*'
				curMonkey.coefficient = 1
			} else { // normal operation (*, +)
				curMonkey.square = false
				curMonkey.operation = rune(line[23])
				coefficient, _ := strconv.Atoi(line[25:])
				curMonkey.coefficient = coefficient
			}
		}

		// test to perform
		if line[:6] == "  Test" {
			test, _ := strconv.Atoi(line[21:])
			curMonkey.test = test

			// save test divider, to calculate modulos
			dividers = append(dividers, test)
		}

		// monkey to move if test true
		if line[:8] == "    If t" {
			tr, _ := strconv.Atoi(line[29:])
			curMonkey.true_ = tr
		}

		// monkey to move if test false
		if line[:8] == "    If f" {
			fa, _ := strconv.Atoi(line[30:])
			curMonkey.false_ = fa

			// monkey is completed, append to slice
			monkeys = append(monkeys, curMonkey)
		}

	}
	return monkeys, tempObj, dividers
}

// REQUIRES: len(monkeys) = len(tempObj) = len(dividers), monkeys, tempObj, dividers not nil
// EFFECTS: returns the monkeys updated with the modulos calculated on the objects (items) divided by each divider
func calulateModulos(monkeys []Monkey, tempObj [][]int, dividers []int) []Monkey {
	// scan each monkey
	for i := 0; i < len(monkeys); i++ {

		// slice to collect modulos for each item for each divider
		objects := make([]map[int]int, 0)

		// scan each item
		for j := 0; j < len(tempObj[i]); j++ {

			// map to collect modulo for each divider
			modulos := make(map[int]int)

			// scan dividers
			for k := 0; k < len(dividers); k++ {
				// calculate modulo for each item for each divider
				mod := tempObj[i][j] % dividers[k]
				modulos[dividers[k]] = mod
			}

			objects = append(objects, modulos)
		}
		monkeys[i].modulos = objects
	}

	return monkeys
}

// REQUIRES: each monkey of "monkeys" not nil, dividors not nil
// EFFECTS: returns the monkeys after simulating "rounds" rounds
func simulateRounds(monkeys []Monkey, dividors []int, rounds int) []Monkey {
	// simulate "rounds" rounds
	for round := 1; round <= rounds; round++ {

		// scan each monkey
		for monkeyIndex, monkey := range monkeys {

			// scan every modulo
			for i := 0; i < len(monkey.modulos); i++ {
				monkeys[monkeyIndex].itemsInspected++

				// refresh modulo for each operation

				// +: mod = (modulo + coefficient) % div
				if monkey.operation == '+' {
					for _, div := range dividors {
						mod := (monkey.modulos[i][div] + monkey.coefficient) % div
						monkeys[monkeyIndex].modulos[i][div] = mod
					}

				} else if monkey.operation == '*' {

					for _, div := range dividors {

						var mod int
						// ^2: mod = (modulo) * (modulo)
						if monkey.square {
							mod = monkey.modulos[i][div] * monkey.modulos[i][div]
						} else { // *: (modulo) * (multiplier % div)
							mod = monkey.modulos[i][div] * (monkey.coefficient % div)
						}

						monkeys[monkeyIndex].modulos[i][div] = mod
					}

				}

				test := monkey.test         // number to divide by to test
				trueM := monkey.true_       // monkey index to move if test true
				falseM := monkey.false_     // movey index to move if test false
				object := monkey.modulos[i] // object to move

				if monkey.modulos[i][test] == 0 { // test true: move to trueM
					monkeys[trueM].modulos = append(monkeys[trueM].modulos, object)
				} else { // test false: move to falseM
					monkeys[falseM].modulos = append(monkeys[falseM].modulos, object)
				}
			}

			// remove object scanned from "original" monkey
			// all object have been moved to "trueM" or "falseM" monkeys
			monkeys[monkeyIndex].modulos = make([]map[int]int, 0)
		}
	}
	return monkeys
}

// REQUIRES: each monkey of "monkeys" not nil
// EFFECTS: returns the product of the 2 monkeys with most items inspected
func getBusinessLevel(monkeys []Monkey) int {
	// top 2 monkeys with most itesmInspected
	var max1, max2 int

	// scan each monkey
	for _, m := range monkeys {

		// save top2
		if m.itemsInspected > max1 { // bigger than the biggest saved
			max2 = max1             // move old biggest to 2nd biggest
			max1 = m.itemsInspected // save new biggest
		} else if m.itemsInspected > max2 { // bigger than 2nd biggest saved (but not biggest)
			max2 = m.itemsInspected // save new 2nd biggest
		}
	}
	return max1 * max2
}
