// https://adventofcode.com/2022/day/20
// https://github.com/Favo02/advent-of-code

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
