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
	fmt.Println("max elf:", top3[2])
	fmt.Println("max 3 elfs total:", top3[0]+top3[1]+top3[2])
}

// REQUIRES: stdin is a valid puzzle input
// EFFECTS: reads from stdin until EOF, calculating top 3 elfs carrying most calories
// MODIFIES: stdin
func maxElf() [3]int {
	var maxCals [3]int = [3]int{0, 0, 0}
	var curElf int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		cal, _ := strconv.Atoi(line)
		curElf += cal

		if line == "" { // current elf ended
			if curElf > maxCals[0] { // check if current elf in top3
				maxCals[0] = curElf
				sort.Ints(maxCals[:])
			}
			curElf = 0 // reset current elf
		}
	}

	// check also last elf (no empty line at EOF)
	if curElf > maxCals[0] { // check if current elf in top3
		maxCals[0] = curElf
		sort.Ints(maxCals[:])
	}
	curElf = 0 // reset current elf

	return maxCals
}
