package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ElfCouple struct {
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

// REQUIRES: stdin is a valid challenge input
// MODIFIES: stdin
// EFFECTS: returns the number of full and partial overlaps in elfs couples
func parseInput() (int, int) {
	var fullCount, partCount int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, ",")

		elf1 := strings.Split(tokens[0], "-")
		elf2 := strings.Split(tokens[1], "-")

		min1, _ := strconv.Atoi(elf1[0])
		max1, _ := strconv.Atoi(elf1[1])

		min2, _ := strconv.Atoi(elf2[0])
		max2, _ := strconv.Atoi(elf2[1])

		elf := ElfCouple{min1, max1, min2, max2}
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

// REQUIRES: e != nil
// EFFECTS: returns (true, true) if the elf couple has a complete overlap,
// returns (false, true) if the elf couple has a partial overlap,
// returns (false, false) if the elf couple has not overlap
func checkOverlap(e ElfCouple) (bool, bool) {

	maxMin := max(e.min1, e.min2) // biggest min
	minMax := min(e.max1, e.max2) // smallest max

	// partial overlap
	if maxMin <= minMax {

		// full overlap
		if e.min1 <= e.min2 && e.max1 >= e.max2 {
			return true, true
		}
		if e.min2 <= e.min1 && e.max2 >= e.max1 {
			return true, true
		}

		return false, true
	}

	// no overlap
	return false, false
}

// EFFECTS: returns smallest number between a and b, a if a = b
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// EFFECTS: returns the biggest number between a and b, a if a = b
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
