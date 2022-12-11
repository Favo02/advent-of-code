package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	objects        []uint64 // save the values of the monkey
	operation      rune     // operation to perform (*, +)
	square         bool     // operation to perform is a square (^2)
	coefficient    uint64   // number to perform the operation on
	test           uint64   // number to test the modulo
	true_          uint64   // monkey to move the valueif test is valid
	false_         uint64   // monkey to move the value if test is invalid
	itemsInspected uint64   // number of items inspected
}

func main() {
	var rounds int = 20

	monkeys := parseInput()
	monkeys = simulateRounds(monkeys, rounds)
	businessLevel := getBusinessLevel(monkeys)

	fmt.Println("business level after", rounds, "rounds:", businessLevel)
}

func parseInput() (monkeys []Monkey) {
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

		// starting objects
		if line[:6] == "  Star" {
			tokens := strings.Split(line[18:], ", ") // save objects

			// converts objects (strings) to ints
			var starting []uint64
			for _, v := range tokens {
				n, _ := strconv.Atoi(v)
				starting = append(starting, uint64(n))
			}

			curMonkey.objects = starting
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
				curMonkey.coefficient = uint64(coefficient)
			}
		}

		// test to perform
		if line[:6] == "  Test" {
			test, _ := strconv.Atoi(line[21:])
			curMonkey.test = uint64(test)
		}

		// monkey to move if test true
		if line[:8] == "    If t" {
			tr, _ := strconv.Atoi(line[29:])
			curMonkey.true_ = uint64(tr)
		}

		// monkey to move if test false
		if line[:8] == "    If f" {
			fa, _ := strconv.Atoi(line[30:])
			curMonkey.false_ = uint64(fa)

			// monkey is completed, append to slice
			monkeys = append(monkeys, curMonkey)
		}

	}
	return monkeys
}

func simulateRounds(monkeys []Monkey, rounds int) []Monkey {
	// simulate "rounds" rounds
	for round := 1; round <= rounds; round++ {

		// scan each monkey
		for mi, m := range monkeys {

			// scan each object (number) of the monkey
			for i := 0; i < len(m.objects); i++ {
				monkeys[mi].itemsInspected++

				// determine operation to perform
				if m.square { // ^2
					m.objects[i] = m.objects[i] * m.objects[i]
				} else {
					switch m.operation {
					case '*': // *
						m.objects[i] = m.objects[i] * m.coefficient
					case '+': // +
						m.objects[i] = m.objects[i] + m.coefficient
					}
				}

				// divide each object by 3
				m.objects[i] = m.objects[i] / 3

				// perform test (divisible by "test")
				if m.objects[i]%m.test == 0 { // true -> move object to monkey "true_"
					monkeys[m.true_].objects = append(monkeys[m.true_].objects, m.objects[i])
				} else { // false -> move object to monkey "false_"
					monkeys[m.false_].objects = append(monkeys[m.false_].objects, m.objects[i])
				}
			}

			// remove object scanned from "original" monkey
			// all object have been moved to "true_" or "false_" monkeys
			monkeys[mi].objects = make([]uint64, 0)
		}
	}
	return monkeys
}

func getBusinessLevel(monkeys []Monkey) uint64 {
	// top 2 monkeys with most itesmInspected
	var max1, max2 uint64

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
