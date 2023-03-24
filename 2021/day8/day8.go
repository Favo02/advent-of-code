// https://adventofcode.com/2021/day/8
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	patterns, digits := parseInput()

	part1 := countUniqueValues(digits)
	fmt.Println("part1", part1)

	part2 := calculateDigits(patterns, digits)
	fmt.Println("part2", part2)
}

// returns list of patters (part before |) and list of digits (part after |)
func parseInput() (patterns, digits []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " | ")
		patterns = append(patterns, tokens[0])
		digits = append(digits, tokens[1])
	}
	return patterns, digits
}

// returns the number of unique digits: 1, 4, 7, 8 (ONLY DIGITS, PART AFTER |)
func countUniqueValues(digits []string) int {
	count := 0

	for _, line := range digits {
		tokens := strings.Split(line, " ")

		for _, d := range tokens {
			switch len(d) {
			case 2, 4, 3, 7:
				count++
			}
		}

	}

	return count
}

// returns the sum of digits
func calculateDigits(patterns, digits []string) int {
	var sum int
	for i, line := range patterns {
		key := decodePattern(line)

		digitsTokens := strings.Split(digits[i], " ")
		sum += calcultateLine(key, digitsTokens)

	}
	return sum
}

// decodes the signals in pattern part (very messy, awful code)
// returns a map: signal --> number
func decodePattern(pattern string) map[string]int {

	key := make(map[int]string)

	tokens := strings.Split(pattern, " ")

	var fiveSegments []string
	var sixSegments []string

	for _, p := range tokens {
		switch len(p) {
		// unique (1, 4, 7, 8)
		case 2: // 1
			key[1] = p
		case 4: // 4
			key[4] = p
		case 3: // 7
			key[7] = p
		case 7: // 8
			key[8] = p

		// 5 segments (2, 3, 5)
		case 5:
			fiveSegments = append(fiveSegments, p)

		// 6 segments (6, 9, 0)
		case 6:
			sixSegments = append(sixSegments, p)
		}
	}

	// 3
	for i, p := range fiveSegments {
		if strings.ContainsRune(p, rune(key[1][0])) && strings.ContainsRune(p, rune(key[1][1])) {
			key[3] = p
			fiveSegments = remove(fiveSegments, i)
			break
		}
	}

	// 6
	for i, p := range sixSegments {
		if !(strings.ContainsRune(p, rune(key[1][0])) && strings.ContainsRune(p, rune(key[1][1]))) {
			key[6] = p
			sixSegments = remove(sixSegments, i)
			break
		}
	}

	// decode top (c) and bottom (f) 1 segments
	var c rune
	if strings.ContainsRune(key[6], rune(key[1][0])) {
		c = rune(key[1][1])
		// f = rune(key[1][0]) // useless
	} else {
		c = rune(key[1][0])
	}

	// 2, 5
	for _, p := range fiveSegments {
		if strings.ContainsRune(p, c) {
			key[2] = p
		} else {
			key[5] = p
		}
	}

	// 9, 0
	for _, p := range sixSegments {
		var found bool
		// contains full 4 --> 9
		// not contains full 4 --> 0
		for _, p2 := range key[4] {
			if !(strings.ContainsRune(p, p2)) {
				key[0] = p
				found = true
			}
		}
		if !found {
			key[9] = p
		}

	}

	return invertMap(key)
}

// removes an element from a slice (slice gets shuffled!)
func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// inverts the decoded map: number (int) --> signal (string) becomes signal --> number
func invertMap(m map[int]string) map[string]int {
	res := make(map[string]int)

	for k, v := range m {
		res[v] = k
	}

	return res
}

// returns the value of digits of a line (concatenated, not added up)
func calcultateLine(key map[string]int, digits []string) int {
	fmt.Println(key)

	var res int
	for _, d := range digits {

		for k, v := range key {
			if compareSets(k, d) {

				res *= 10
				res += v

				break
			}
		}
	}
	fmt.Println(digits, res)
	return res
}

// returns true if two strings contains the same elements (compare two sets)
func compareSets(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, va := range a {
		if !strings.ContainsRune(b, va) {
			return false
		}
	}

	return true
}
