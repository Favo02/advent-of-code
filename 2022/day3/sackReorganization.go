// https://adventofcode.com/2022/day/3

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tot, totGroup := parseInput()
	fmt.Println(tot, totGroup)
}

// REQUIRES: stdin is a valid puzzle input
// MODIFIES: stdin
// EFFECTS: returns the priority for every elf (part1) and the priority for every group of elfs (part2)
func parseInput() (int, int) {
	var tot, totGroup int
	scanner := bufio.NewScanner(os.Stdin)

	var buffer [3]string // buffer to save last 3 sacks
	var bufsize int      // number of sacks saved in buffer
	for scanner.Scan() {
		line := scanner.Text()

		// --- part1 --- //
		substr1 := line[:len(line)/2]
		substr2 := line[len(line)/2:]
		commonChar := compare2Strings(substr1, substr2)
		priority := charToVal(commonChar)
		tot += priority

		// --- part2 --- //
		buffer[bufsize] = line
		bufsize++

		if bufsize == 3 { // each group of elfs
			charGroup := compare3Strings(buffer)
			groupPriority := charToVal(charGroup)
			totGroup += groupPriority

			bufsize = 0 // reset buffer to start new elfs group
		}
	}
	return tot, totGroup
}

// EFFECTS: returns the common character between str1 and str2, '0' if there are no common characters
func compare2Strings(str1, str2 string) rune {
	for _, v := range str1 {
		for _, w := range str2 {
			if v == w {
				return v
			}
		}
	}
	return '0'
}

// EFFECTS: returns the common character between the 3 strings contained in strings, '0' if there are no common characters
func compare3Strings(strings [3]string) rune {
	for _, v := range strings[0] {
		for _, w := range strings[1] {
			// run 3rd loop only if a match between first two strings is already found
			if v == w {
				for _, z := range strings[2] {
					if w == z {
						return v
					}
				}
			}
		}
	}
	return '0'
}

// EFFECTS: returns the priority of the character c
func charToVal(c rune) int {
	if c >= 'a' && c <= 'z' {
		var base rune = 'a'
		return int(c-base) + 1
	} else if c >= 'A' && c <= 'Z' {
		var base rune = 'A'
		return int(c-base) + 27
	}
	return 0
}
