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
	var tot, totGroup int
	scanner := bufio.NewScanner(os.Stdin)

	var buffer [3]string
	var bufsize int
	for scanner.Scan() {
		line := scanner.Text()

		// part1
		part1 := line[:len(line)/2]
		part2 := line[len(line)/2:]
		fmt.Println(part1, part2)
		commonChar := compare(part1, part2)
		priority := charToVal(commonChar)
		fmt.Println(string(commonChar), priority)
		tot += priority

		// part2
		buffer[bufsize] = line
		bufsize++

		if bufsize == 3 {
			charGroup := compareBuffer(buffer)
			groupPriority := charToVal(charGroup)
			totGroup += groupPriority

			bufsize = 0
		}
	}
	fmt.Println(tot, totGroup)
}

func compare(str1, str2 string) rune {
	for _, v := range str1 {
		for _, w := range str2 {
			if v == w {
				return v
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
	} else {
		fmt.Println("????")
		return 0
	}
}

func compareBuffer(buf [3]string) rune {
	for _, v := range buf[0] {
		for _, w := range buf[1] {
			for _, z := range buf[2] {
				if v == w && w == z {
					return v
				}
			}
		}
	}
	return '0'
}
