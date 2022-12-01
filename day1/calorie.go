// https://adventofcode.com/2022/day/1

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	top3 := maxElf()
	fmt.Println("max 3 elfs:", top3)
	fmt.Println("max 3 elfs total:", top3[0]+top3[1]+top3[2])
}

// EFFECTS: reads from stdin until EOF, calculating top 3 elfs with most calories
// MODIFIES: stdin
func maxElf() [3]int {
	var maxCals [3]int = [3]int{0, 0, 0}
	var curElf int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if line == "" {
			if curElf > maxCals[0] {
				maxCals[0] = curElf
				sort.Ints(maxCals[:])
			}
			curElf = 0
		} else {
			cal, _ := strconv.Atoi(line)
			curElf += cal
		}
	}
	return maxCals
}
