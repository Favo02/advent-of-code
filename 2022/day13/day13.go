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
	packets := parseInput()

	part1 := comparePairs(packets) // compare packet pairs

	packets = append(packets, "[[2]]")
	packets = append(packets, "[[6]]")
	packets = sortPackets(packets)
	part2 := findPacketIndex(packets, "[[2]]") * findPacketIndex(packets, "[[6]]")

	fmt.Println("sum of index of correct order packet pairs")
	fmt.Println("\tpart1:", part1)
	fmt.Println("product of indexes of packet [[2]] and [[6]]")
	fmt.Println("\tpart2:", part2)
}

// returns the packets
func parseInput() (packets []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			packets = append(packets, line)
		}
	}
	return packets
}

// returns the sum of indexes (starting from 1) of pairs of packets that are in the correct order
func comparePairs(packets []string) int {
	var res int
	var index int // packet pair index
	for i := 0; i < len(packets); i += 2 {
		ord := compareLists(packets[i], packets[i+1]) // compare pair

		if ord == 1 { // pair is in order
			res += index + 1 // sum pair index
		}
		index++
	}
	return res
}

// compare two lists: returns 1 if p0 < p1, 0 if p0 == p1, -1 if p0 > p1
func compareLists(p0, p1 string) int {
	// same lists, return 0
	if p0 == p1 {
		return 0
	}

	var i0, i1 int // index to analyze

	// scanning the whole list
	for i0 < len(p0) && i1 < len(p1) {

		// skip comma (,)
		if p0[i0] == ',' {
			i0++
			continue
		}
		if p1[i1] == ',' {
			i1++
			continue
		}

		// compare two lists: recursive until int comparison
		if p0[i0] == '[' && p1[i1] == '[' {

			// calculate the dimension of list
			var end0, end1 int
			end0 = listClosing(p0[i0:])
			end1 = listClosing(p1[i1:])

			// recursive on content of list
			cmp := compareLists(p0[i0+1:i0+end0], p1[i1+1:i1+end1])

			if cmp != 0 { // result found
				return cmp
			} else { // same list, keep scanning (jump the end of list)
				i0 = i0 + end0 + 1
				i1 = i1 + end1 + 1
				continue
			}
		}

		// compare two ints
		if unicode.IsDigit(rune(p0[i0])) && unicode.IsDigit(rune(p1[i1])) {

			// parse first int
			var int0 int
			// parse longer than 1 digit numbers
			if indexOf(p0[i0:], ',') != -1 { // if contains comma (,)
				int0, _ = strconv.Atoi(p0[i0 : i0+indexOf(p0[i0:], ',')])
			} else { // no comma: int = whole string
				int0, _ = strconv.Atoi(p0[i0:])
			}

			// parse second int
			var int1 int
			// parse longer than 1 digit numbers
			if indexOf(p1[i1:], ',') != -1 { // if contains comma (,)
				int1, _ = strconv.Atoi(p1[i1 : i1+indexOf(p1[i1:], ',')])
			} else { // no comma: int = whole string
				int1, _ = strconv.Atoi(p1[i1:])
			}

			// compare ints
			if int0 < int1 {
				return 1
			}
			if int0 > int1 {
				return -1
			}

			// same int, keep scanning (jump to end of ints)
			if int0 == int1 {
				i0 += len(strconv.Itoa(int0))
				i1 += len(strconv.Itoa(int1))
				continue
			}
		}

		// one list and one int: parse int to list

		if p0[i0] == '[' { // first el is a list
			end0 := listClosing(p0[i0:]) // find end of list

			// parse int (second element)
			var int1 int
			if indexOf(p1[i1:], ',') != -1 { // if contains comma (,)
				int1, _ = strconv.Atoi(p1[i1 : i1+indexOf(p1[i1:], ',')])
			} else { // no comma: int = whole string
				int1, _ = strconv.Atoi(p1[i1:])
			}

			// compare list content and int
			cmp := compareLists(p0[i0+1:i0+end0], strconv.Itoa(int1))

			if cmp != 0 { // result found
				return cmp
			} else { // same list and int (converted to list), keep scanning (jump to end of int and end of list)
				i0 = i0 + end0 + 1
				i1 += len(strconv.Itoa(int1))
				continue
			}
		}

		if p1[i1] == '[' { // second el is a list
			end1 := listClosing(p1[i1:]) // find end of list

			// parse int (first element)
			var int0 int
			if indexOf(p0[i0:], ',') != -1 { // if contains comma (,)
				int0, _ = strconv.Atoi(p0[i0 : i0+indexOf(p0[i0:], ',')])
			} else { // no comma: int = whole string
				int0, _ = strconv.Atoi(p0[i0:])
			}

			// compare int and list content
			cmp := compareLists(strconv.Itoa(int0), p1[i1+1:i1+end1])

			if cmp != 0 { // result found
				return cmp
			} else { // same int (converted to list) and list, keep scanning (jump to end of list and end of int)
				i1 = i1 + end1 + 1
				i0 += len(strconv.Itoa(int0))
				continue
			}
		}
	}

	// end of list, one list is shorted tha the otherone
	if len(p0) < len(p1) { // first list shorter
		return 1
	} else { // second list shorter
		return -1
	}
}

// return index of end of list (closing ] for first [)
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

// return the index of runr "r" in string "s", -1 if not found
func indexOf(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}

// returns packets "p" sorted
func sortPackets(p []string) []string {
	sort.Slice(p, func(i, j int) bool {
		return compareLists(p[i], p[j]) == 1
	})
	return p
}

// returns the index of pachet "s" in packets "p"
func findPacketIndex(p []string, s string) int {
	for i, v := range p {
		if v == s {
			return i + 1
		}
	}
	return -1
}
