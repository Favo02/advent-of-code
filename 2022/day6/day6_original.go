package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(parseInput())
}

func parseInput() int {
	scanner := bufio.NewScanner(os.Stdin)

	var buffer string
	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range line {
			if len(buffer) < 14 { // 14 part2, 4 part1
				buffer = buffer + string(v)
				continue
			}
			buffer = buffer[1:] + string(v)
			fmt.Println(buffer)
			if !hasDuplicateChar(buffer) {
				return i + 1
			}
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
