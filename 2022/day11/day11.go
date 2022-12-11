package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	monkeys := parseInput()
	monkeys = start(monkeys)
	fmt.Println(getBusinessLevel(monkeys))
}

type Monkey struct {
	objects        []uint64
	operation      rune
	square         bool
	multiplier     uint64
	test           uint64
	true_          uint64
	false_         uint64
	itemsInspected uint64
}

func parseInput() (monkeys []Monkey) {
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
			var starting []uint64
			for _, v := range tokens {
				n, _ := strconv.Atoi(v)
				starting = append(starting, uint64(n))
			}
			curMonkey.objects = starting
		}

		if line[:6] == "  Oper" {
			if len(line) == 28 && line == "  Operation: new = old * old" {
				curMonkey.square = true
				continue
			}
			op := line[23]
			mult, _ := strconv.Atoi(line[25:])
			curMonkey.operation = rune(op)
			curMonkey.multiplier = uint64(mult)
		}

		if line[:6] == "  Test" {
			test, _ := strconv.Atoi(line[21:])
			curMonkey.test = uint64(test)
		}

		if line[:8] == "    If t" {
			tr, _ := strconv.Atoi(line[29:])
			curMonkey.true_ = uint64(tr)
		}

		if line[:8] == "    If f" {
			fa, _ := strconv.Atoi(line[30:])
			curMonkey.false_ = uint64(fa)

			monkeys = append(monkeys, curMonkey)
		}

	}
	return monkeys
}

func start(monkeys []Monkey) []Monkey {
	for round := 1; round <= 1000; round++ {

		for mi, m := range monkeys {

			for i := 0; i < len(m.objects); i++ {
				monkeys[mi].itemsInspected++

				if m.square {
					m.objects[i] = m.objects[i] * m.objects[i]
				} else {
					switch m.operation {
					case '*':
						m.objects[i] = m.objects[i] * m.multiplier
					case '+':
						m.objects[i] = m.objects[i] + m.multiplier
					}
				}

				// m.objects[i] = m.objects[i] / 3

				if m.objects[i]%m.test == 0 {
					monkeys[m.true_].objects = append(monkeys[m.true_].objects, m.objects[i])
				} else {
					monkeys[m.false_].objects = append(monkeys[m.false_].objects, m.objects[i])
				}

				// fmt.Println("---")
				// fmt.Println(monkeys)

			}
			monkeys[mi].objects = make([]uint64, 0)

		}
		if round == 1000 { // 1000 = 27019168
			fmt.Println("---")
			fmt.Println("1000 =", getBusinessLevel(monkeys))
		}

	}

	return monkeys
}

func getBusinessLevel(monkeys []Monkey) uint64 {
	var max1, max2 uint64
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
