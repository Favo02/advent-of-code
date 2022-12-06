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

func parseInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var str string
	for scanner.Scan() {
		line := scanner.Text()
		str += line
	}
	return str
}

func checkSubstrings(input string, length int) int {
	var buffer string
	for i, char := range input {
		if len(buffer) < length {
			buffer = buffer + string(char)
			continue
		}
		buffer = buffer[1:] + string(char)
		if !hasDuplicateChar(buffer) {
			return i + 1
		}
	}
	return -1
}

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
