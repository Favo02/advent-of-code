// https://adventofcode.com/2022/day/6

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := parseInput()
	part1 := checkSubstrings(input, 4)
	part2 := checkSubstrings(input, 14)
	fmt.Println(" 4 characters (part1):", part1)
	fmt.Println("14 characters (part2):", part2)
}

// REQUIRES: stdin is a valid challenge input
// MODIFIES: stdin
// EFFECTS: returns a string read from stdin terminated by EOF
func parseInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var str string
	for scanner.Scan() {
		line := scanner.Text()
		str += line
	}
	return str
}

// REQUIRES: length >= 0
// EFFECTS: returns the index of index of the last character added to a string of "length" length that has no duplicate characters
func checkSubstrings(input string, length int) int {
	for i := 0; i < len(input)-length; i++ {
		if !hasDuplicateChar(input[i : i+length]) {
			return i + length // index of last char added to substring
		}
	}
	return -1
}

// EFFECTS: returns true if "str" has at least one duplicate characted, false otherwise
func hasDuplicateChar(str string) bool {
	for i, v := range str {
		for j, w := range str {
			if i != j && v == w {
				return true
			}
		}
	}
	return false
}
