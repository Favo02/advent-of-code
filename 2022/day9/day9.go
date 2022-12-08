// https://adventofcode.com/2022/day/9

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	parseInput()
}

func parseInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
