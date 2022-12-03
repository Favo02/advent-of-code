package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1, part2 := parseInput()
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func parseInput() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)

	var horizontal, depth int // part1
	var aim, depth2 int       // part2

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		movement := tokens[0]
		amount, _ := strconv.Atoi(tokens[1])

		switch movement {
		case "forward":
			horizontal += amount
			depth2 += amount * aim
		case "down":
			depth += amount
			aim += amount
		case "up":
			depth -= amount
			aim -= amount
		}
	}

	part1 := horizontal * depth
	part2 := horizontal * depth2

	return part1, part2
}
