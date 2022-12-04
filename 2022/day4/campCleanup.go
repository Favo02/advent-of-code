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
	full, partial := parseInput()
	fmt.Println("Full overlap (part1):\t", full)
	fmt.Println("Partial overlap (part2):", partial)
}

func parseInput() (int, int) {
	var fullCount, partCount int

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
		fullOverlap, partialOverlap := checkOverlap(elf)

		if fullOverlap {
			fullCount++
		}
		if partialOverlap {
			partCount++
		}
	}

	return fullCount, partCount
}

func checkOverlap(e Elf) (bool, bool) {
	// complete overlap
	if e.min1 <= e.min2 && e.max1 >= e.max2 {
		return true, true
	}
	if e.min2 <= e.min1 && e.max2 >= e.max1 {
		return true, true
	}

	// partial overlap
	if e.max1 >= e.min2 && e.min1 < e.max2 {
		return false, true
	}
	if e.max2 >= e.min1 && e.min2 < e.max1 {
		return false, true
	}

	// no overlap
	return false, false
}