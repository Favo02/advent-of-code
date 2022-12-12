// https://adventofcode.com/2022/day/13

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := parseInput()
	fmt.Println(lines)
}

func parseInput() (lines []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
