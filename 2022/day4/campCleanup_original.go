package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	min1 int
	max1 int
	min2 int
	max2 int
}

func main() {
	elfs := parseInput()
	fmt.Println(elfs)
	fully, over := countOverlap(elfs)
	fmt.Println(fully)
	fmt.Println(over)
}

func parseInput() []Elf {
	var elfs []Elf

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, ",")

		part1 := strings.Split(tokens[0], "-")
		min1, _ := strconv.Atoi(part1[0])
		max1, _ := strconv.Atoi(part1[1])

		part2 := strings.Split(tokens[1], "-")
		min2, _ := strconv.Atoi(part2[0])
		max2, _ := strconv.Atoi(part2[1])

		elf := Elf{min1, max1, min2, max2}
		elfs = append(elfs, elf)
	}

	return elfs
}

func countOverlap(elfs []Elf) (int, int) {
	var fullyOverlap, overlap int
	for _, e := range elfs {
		if e.min1 <= e.min2 && e.max1 >= e.max2 {
			fmt.Println(e)
			fullyOverlap++
			overlap++
			continue
		}
		if e.min2 <= e.min1 && e.max2 >= e.max1 {
			fmt.Println(e)
			fullyOverlap++
			overlap++
			continue
		}

		if e.max1 >= e.min2 && e.min1 < e.max2 {
			fmt.Println("-", e)
			overlap++
			continue
		}
		if e.max2 >= e.min1 && e.min2 < e.max1 {
			fmt.Println("-", e)
			overlap++
			continue
		}
	}
	return fullyOverlap, overlap
}
