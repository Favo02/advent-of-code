// https://adventofcode.com/2022/day/13

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

func main() {
	pairs, packets := parseInput()

	part1 := comparePairs(pairs)

	packets = append(packets, "[[2]]")
	packets = append(packets, "[[6]]")
	packets = sortP(packets)
	part2 := findP(packets, "[[2]]") * findP(packets, "[[6]]")
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func parseInput() (pairs [][]string, packets []string) {
	scanner := bufio.NewScanner(os.Stdin)
	pair := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			pairs = append(pairs, pair)
			pair = make([]string, 0)
		} else {
			packets = append(packets, line)
			pair = append(pair, line)
		}
	}
	pairs = append(pairs, pair)
	return pairs, packets
}

func comparePairs(pairs [][]string) int {
	var res int
	for i, p := range pairs {
		ord := compareLists(p[0], p[1])
		// fmt.Println("cmp:", p[0], p[1], ord)
		// fmt.Println()
		if ord == 1 {
			res += i + 1
		}
	}
	return res
}

func listClosing(s1 string) int {
	var open int
	for i, c := range s1 {
		if c == '[' {
			open++
			continue
		}
		if c == ']' {
			if open == 1 {
				return i
			} else {
				open--
			}
		}
	}
	return -1
}

func indexOf(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}

func contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

func compareLists(p0, p1 string) int {
	if p0 == p1 {
		// fmt.Println("equal")
		return 0
	}

	var i0, i1 int // index to analyze
	for i0 < len(p0) && i1 < len(p1) {
		// fmt.Println("comparing", p0[i0:], p1[i1:])

		// skip ,
		if p0[i0] == ',' {
			i0++
			continue
		}
		if p1[i1] == ',' {
			i1++
			continue
		}

		// ------ BOTH LIST --------
		// recursive until int comparison
		if p0[i0] == '[' && p1[i1] == '[' {
			// fmt.Println("both lists: recursion")

			// calculate the dimension of list
			var end0, end1 int
			end0 = listClosing(p0[i0:])
			end1 = listClosing(p1[i1:])

			cmp := compareLists(p0[i0+1:i0+end0], p1[i1+1:i1+end1])
			if cmp != 0 {
				return cmp
			} else {
				i0 = i0 + end0 + 1
				i1 = i1 + end1 + 1
				continue
			}
		}

		// ------ BOTH INT --------
		if unicode.IsDigit(rune(p0[i0])) && unicode.IsDigit(rune(p1[i1])) {

			var int0 int
			if contains(p0[i0:], ',') {
				int0, _ = strconv.Atoi(p0[i0 : i0+indexOf(p0[i0:], ',')])
			} else {
				int0, _ = strconv.Atoi(p0[i0:])
			}

			var int1 int
			if contains(p1[i1:], ',') {
				int1, _ = strconv.Atoi(p1[i1 : i1+indexOf(p1[i1:], ',')])
			} else {
				int1, _ = strconv.Atoi(p1[i1:])
			}

			// fmt.Println("both ints:", int0, int1)

			if int0 < int1 {
				return 1
			}
			if int0 > int1 {
				return -1
			}

			if int0 == int1 {
				i0 += len(strconv.Itoa(int0))
				i1 += len(strconv.Itoa(int1))
				// fmt.Println("same int: continue")
				continue
			}
		}

		// ------ ONE LIST AND ONE INT --------
		if p0[i0] == '[' {
			end0 := listClosing(p0[i0:])

			var int1 int
			if contains(p1[i1:], ',') {
				int1, _ = strconv.Atoi(p1[i1 : i1+indexOf(p1[i1:], ',')])
			} else {
				int1, _ = strconv.Atoi(p1[i1:])
			}

			// fmt.Println("first list:", p0[i0+1:i0+end0], int1)
			cmp := compareLists(p0[i0+1:i0+end0], strconv.Itoa(int1))

			if cmp != 0 {
				return cmp
			} else {
				i0 = i0 + end0 + 1
				i1 += len(strconv.Itoa(int1))
				continue
			}
		}

		if p1[i1] == '[' {
			end1 := listClosing(p1[i1:])

			var int0 int
			if contains(p0[i0:], ',') {
				int0, _ = strconv.Atoi(p0[i0 : i0+indexOf(p0[i0:], ',')])
			} else {
				int0, _ = strconv.Atoi(p0[i0:])
			}

			// fmt.Println("first list:", p1[i1+1:i1+end1], int0)
			cmp := compareLists(strconv.Itoa(int0), p1[i1+1:i1+end1])

			if cmp != 0 {
				return cmp
			} else {
				i1 = i1 + end1 + 1
				i0 += len(strconv.Itoa(int0))
				continue
			}
		}

		i0++
		i1++
	}

	// ------ ONE EMPTY OR EQUAL --------
	if len(p0) < len(p1) {
		// fmt.Println("p0 less elms")
		return 1
	}
	if len(p0) > len(p1) {
		// fmt.Println("p1 less elms")
		return -1
	}

	return 0
}

func findP(p []string, s string) int {
	for i, v := range p {
		if v == s {
			return i + 1
		}
	}
	return -1
}

func sortP(p []string) []string {
	sort.Slice(p, func(i, j int) bool {
		return boolCompareLists(p[i], p[j])
	})
	return p
}

func boolCompareLists(i, j string) bool {
	res := compareLists(i, j)
	if res == 1 {
		return true
	}
	return false
}
