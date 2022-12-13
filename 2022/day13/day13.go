// https://adventofcode.com/2022/day/13

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	pairs := parseInput()
	// fmt.Println(pairs)
	res := comparePairs(pairs)
	fmt.Println(res)
}

func parseInput() (pairs [][]string) {
	scanner := bufio.NewScanner(os.Stdin)
	pair := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			pairs = append(pairs, pair)
			pair = make([]string, 0)
		} else {
			pair = append(pair, line)
		}
	}
	pairs = append(pairs, pair)
	return pairs
}

func comparePairs(pairs [][]string) int {
	var res int
	var corP []int
	for i, p := range pairs {
		ord := compareLists(p[0], p[1])
		fmt.Println("res:", ord, "\n")
		if ord == 1 {
			corP = append(corP, i+1)
			res += i + 1
		}
	}
	fmt.Println(corP)
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
			if open == 0 {
				return i
			} else {
				open--
			}
		}
	}
	return -1
}

func compareLists(p0, p1 string) int {
	fmt.Println("comparing", p0, p1)

	var i0, i1 int
	for i0 < len(p0) && i1 < len(p1) {

		if p0[i0] == ',' {
			i0++
			continue
		}
		if p1[i1] == ',' {
			i1++
			continue
		}

		// both lists
		if p0[i0] == '[' && p1[i1] == '[' {
			fmt.Println("both lists:")
			var endList0, endList1 int
			endList0 = listClosing(p0[i0+1:]) + 1
			endList1 = listClosing(p1[i1+1:]) + 1

			cmp := compareLists(p0[i0+1:i0+endList0], p1[i1+1:i1+endList1])
			if cmp != 0 {
				return cmp
			} else {
				i0 = endList0
				i1 = endList1
				continue
			}
		}

		// both int
		if (p0[i0] >= '0' && p0[i0] <= '9') && (p1[i1] >= '0' && p1[i1] <= '9') {

			var int0 int
			if contains(p0, ',') {
				int0, _ = strconv.Atoi(p0[i0 : i0+indexOf(p0[i0:], ',')])
			} else {
				int0, _ = strconv.Atoi(p0[i0:])
			}

			var int1 int
			if contains(p1, ',') {
				int1, _ = strconv.Atoi(p1[i1 : i1+indexOf(p1[i1:], ',')])
			} else {
				int1, _ = strconv.Atoi(p1[i1:])
			}

			fmt.Println("both ints:", int0, int1)
			if int0 < int1 {
				return 1
			}
			if int0 == int1 {
				i0 += len(p0[i0 : i0+indexOf(p0[i0:], ',')])
				i1 += len(p1[i1 : i1+indexOf(p1[i1:], ',')])
				continue
			}
			if int0 > int1 {
				return -1
			}
		}

		// one list
		if p0[i0] == '[' {
			endList0 := listClosing(p0[i0+1:])
			var integer string
			if contains(p1, ',') {
				integer = p1[i1 : i1+indexOf(p1[i1:], ',')]
			} else {
				integer = p1[i1:]
			}
			fmt.Println("first list:", p0[i0+1:i0+1+endList0], integer)
			cmp := compareLists(p0[i0+1:i0+1+endList0], integer)
			if cmp != 0 {
				return cmp
			} else {
				i0 = endList0
				i1 += len(integer)
				continue
			}
		}
		if p1[i1] == '[' {
			endList1 := listClosing(p1[i1+1:])
			var integer string
			if contains(p0, ',') {
				integer = p0[i0 : i0+indexOf(p0[i0:], ',')]
			} else {
				integer = p0[i0:]
			}
			fmt.Println("second list:", integer, p1[i1+1:i1+1+endList1])
			cmp := compareLists(integer, p1[i1+1:i1+1+endList1])
			if cmp != 0 {
				return cmp
			} else {
				i0 += len(integer)
				i1 = endList1
				continue
			}
		}

		i0++
		i1++

	}
	if len(p0) < len(p1) {
		fmt.Println("p0 less elms")
		return 1
	}
	if len(p0) > len(p1) {
		fmt.Println("p1 less elms")
		return -1
	}

	return 0
}

func contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

func indexOf(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return 1
}
