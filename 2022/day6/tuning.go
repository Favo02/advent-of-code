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
	fmt.Println(input)
	for i := 0; i < len(input)-length; i++ {
		fmt.Println(input[i : i+length])
		if !hasDuplicateChar(input[i : i+length]) {
			return i + length // index of last char added to substring
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
