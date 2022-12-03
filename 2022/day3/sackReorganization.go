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

func parseInput() (int, int) {
	var tot, totGroup int
	scanner := bufio.NewScanner(os.Stdin)

	var buffer [3]string // buffer to save last 3 sacks
	var bufsize int      // number of sacks saved in buffer
	for scanner.Scan() {
		line := scanner.Text()

		// part1
		substr1 := line[:len(line)/2]
		substr2 := line[len(line)/2:]
		commonChar := compareStrings(substr1, substr2)
		priority := charToVal(commonChar)
		// fmt.Println(string(commonChar), priority)
		tot += priority

		// part2
		buffer[bufsize] = line
		bufsize++

		if bufsize == 3 { // each group of elfs
			charGroup := compareBuffer(buffer)
			groupPriority := charToVal(charGroup)
			totGroup += groupPriority

			bufsize = 0
		}
	}
	return tot, totGroup
}

func compareStrings(str1, str2 string) rune {
	for _, v := range str1 {
		for _, w := range str2 {
			if v == w {
				return v
			}
		}
	}
	return '0'
}

func compareBuffer(buf [3]string) rune {
	for _, v := range buf[0] {
		for _, w := range buf[1] {
			// run 3rd loop only if a match between first two strings is already found
			if v == w {
				for _, z := range buf[2] {
					if w == z {
						return v
					}
				}
			}
		}
	}
	return '0'
}

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
