package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	objectsOffsets []map[int]int // every object to what is divisible by
	operation      rune
	square         bool
	multiplier     int
	test           int
	true_          int
	false_         int
	itemsInspected int
}

func main() {
	monkeys, tempObj, conditions := parseInput()
	monkeys = generateObjects(monkeys, tempObj, conditions)
	monkeys = start(monkeys, conditions, 10000)
	fmt.Println(getBusinessLevel(monkeys))
}

func parseInput() (monkeys []Monkey, tempObj [][]int, conditions []int) {
	scanner := bufio.NewScanner(os.Stdin)

	var curMonkey Monkey
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if line[:6] == "Monkey" {
			curMonkey = Monkey{}
		}

		if line[:6] == "  Star" {
			tokens := strings.Split(line[18:], ", ")
			var starting []int
			for _, v := range tokens {
				n, _ := strconv.Atoi(v)
				starting = append(starting, n)
			}
			tempObj = append(tempObj, starting)
		}

		if line[:6] == "  Oper" {
			if len(line) == 28 && line == "  Operation: new = old * old" {
				curMonkey.square = true
				curMonkey.operation = '*'
				curMonkey.multiplier = 1
			} else {
				curMonkey.square = false
				curMonkey.operation = rune(line[23])
				multiplier, _ := strconv.Atoi(line[25:])
				curMonkey.multiplier = multiplier
			}
		}

		if line[:6] == "  Test" {
			test, _ := strconv.Atoi(line[21:])
			conditions = append(conditions, test)
			curMonkey.test = test
		}

		if line[:8] == "    If t" {
			tr, _ := strconv.Atoi(line[29:])
			curMonkey.true_ = tr
		}

		if line[:8] == "    If f" {
			fa, _ := strconv.Atoi(line[30:])
			curMonkey.false_ = fa

			monkeys = append(monkeys, curMonkey)
		}

	}
	return monkeys, tempObj, conditions
}

func generateObjects(monkeys []Monkey, tempObj [][]int, conditions []int) []Monkey {
	for i := 0; i < len(monkeys); i++ {

		objects := make([]map[int]int, 0)

		for j := 0; j < len(tempObj[i]); j++ {

			offsets := make(map[int]int)

			for k := 0; k < len(conditions); k++ {

				mod := tempObj[i][j] % conditions[k]
				if mod == 0 {
					offsets[conditions[k]] = 0
				} else {
					offsets[conditions[k]] = mod
				}
			}

			objects = append(objects, offsets)
		}
		monkeys[i].objectsOffsets = objects
	}

	return monkeys
}

func start(monkeys []Monkey, conditions []int, rounds int) []Monkey {
	for round := 1; round <= rounds; round++ {

		for monkeyIndex := 0; monkeyIndex < len(monkeys); monkeyIndex++ {

			// scan every object
			for i := 0; i < len(monkeys[monkeyIndex].objectsOffsets); i++ {
				monkeys[monkeyIndex].itemsInspected++

				// refresh divisibility for each operation

				if monkeys[monkeyIndex].square {

					// refresh every divisibility
					for _, cond := range conditions {

						// (n % m) * (k % m)
						mod := monkeys[monkeyIndex].objectsOffsets[i][cond] * monkeys[monkeyIndex].objectsOffsets[i][cond]

						monkeys[monkeyIndex].objectsOffsets[i][cond] = mod
					}

				} else if monkeys[monkeyIndex].operation == '+' {

					// refresh every divisibility
					for _, cond := range conditions {
						mod := (monkeys[monkeyIndex].multiplier + monkeys[monkeyIndex].objectsOffsets[i][cond]) % cond
						monkeys[monkeyIndex].objectsOffsets[i][cond] = mod
					}

				} else if monkeys[monkeyIndex].operation == '*' {

					// refresh every divisibility
					for _, cond := range conditions {

						// (n % m) * (k % m)
						mod := monkeys[monkeyIndex].objectsOffsets[i][cond] * (monkeys[monkeyIndex].multiplier % cond)

						monkeys[monkeyIndex].objectsOffsets[i][cond] = mod

					}

				}

				test := monkeys[monkeyIndex].test
				trueM := monkeys[monkeyIndex].true_
				falseM := monkeys[monkeyIndex].false_

				objectCopy := make(map[int]int)
				for k, v := range monkeys[monkeyIndex].objectsOffsets[i] {
					objectCopy[k] = v
				}

				if monkeys[monkeyIndex].objectsOffsets[i][test] == 0 {
					monkeys[trueM].objectsOffsets = append(monkeys[trueM].objectsOffsets, objectCopy)
				} else {
					monkeys[falseM].objectsOffsets = append(monkeys[falseM].objectsOffsets, objectCopy)
				}
			}

			monkeys[monkeyIndex].objectsOffsets = make([]map[int]int, 0)
		}
	}
	printMonkeys(monkeys)
	return monkeys
}

func getBusinessLevel(monkeys []Monkey) int {
	var max1, max2 int
	for _, m := range monkeys {
		if m.itemsInspected > max1 {
			max2 = max1
			max1 = m.itemsInspected
		} else if m.itemsInspected > max2 {
			max2 = m.itemsInspected
		}
	}
	fmt.Println(max1, ",", max2)
	return max1 * max2
}

func printMonkeys(monkeys []Monkey) {
	for i, m := range monkeys {
		fmt.Print("M", i, ":", len(m.objectsOffsets), "-", m.itemsInspected, " ")
	}
	fmt.Println()
}
