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

// REQUIRES: stdin is a valid challenge input
// MODIFIES: stdint
// EFFECTS: returns the packets parsed from stdin
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

// REQUIRES: each packet of "packets" is valid
// EFFECTS: returns the sum of indexes (starting from 1) of pairs of packets that are in correct order
func comparePairs(packets []string) int {
	var res int
	var index int // packet pair index
	for i := 0; i < len(packets); i += 2 {
		ord := comparePackets(packets[i], packets[i+1]) // compare pair

		if ord == 1 { // pair is in order
			res += index + 1 // sum pair index
		}
		index++
	}
	return res
}

// REQUIRES: pack0 and pack1 are valid packets
// EFFECTS: compare two packets: returns 1 if pack0 < pack1, 0 if pack0 == pack1, -1 if pack0 > pack1
func comparePackets(pack0, pack1 string) int {
	// same lists, return 0
	if pack0 == pack1 {
		return 0
	}

	var index0, index1 int // index to analyze

	// scanning the whole list
	for index0 < len(pack0) && index1 < len(pack1) {

		// skip comma (,)
		if pack0[index0] == ',' {
			index0++
			continue
		}
		if pack1[index1] == ',' {
			index1++
			continue
		}

		// compare two lists: recursive until int comparison
		if pack0[index0] == '[' && pack1[index1] == '[' {

			// calculate the dimension of list
			var end0, end1 int
			end0 = listClosing(pack0[index0:])
			end1 = listClosing(pack1[index1:])

			// recursive on content of list
			cmp := comparePackets(pack0[index0+1:index0+end0], pack1[index1+1:index1+end1])

			if cmp != 0 { // result found
				return cmp
			} else { // same list, keep scanning (jump the end of list)
				index0 = index0 + end0 + 1
				index1 = index1 + end1 + 1
				continue
			}
		}

		// compare two ints
		if unicode.IsDigit(rune(pack0[index0])) && unicode.IsDigit(rune(pack1[index1])) {

			// parse first int
			var int0 int
			// parse longer than 1 digit numbers
			if indexOf(pack0[index0:], ',') != -1 { // if contains comma (,)
				int0, _ = strconv.Atoi(pack0[index0 : index0+indexOf(pack0[index0:], ',')])
			} else { // no comma: int = whole string
				int0, _ = strconv.Atoi(pack0[index0:])
			}

			// parse second int
			var int1 int
			// parse longer than 1 digit numbers
			if indexOf(pack1[index1:], ',') != -1 { // if contains comma (,)
				int1, _ = strconv.Atoi(pack1[index1 : index1+indexOf(pack1[index1:], ',')])
			} else { // no comma: int = whole string
				int1, _ = strconv.Atoi(pack1[index1:])
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
				index0 += len(strconv.Itoa(int0))
				index1 += len(strconv.Itoa(int1))
				continue
			}
		}

		// one list and one int: parse int to list

		if pack0[index0] == '[' { // first el is a list
			end0 := listClosing(pack0[index0:]) // find end of list

			// parse int (second element)
			var int1 int
			if indexOf(pack1[index1:], ',') != -1 { // if contains comma (,)
				int1, _ = strconv.Atoi(pack1[index1 : index1+indexOf(pack1[index1:], ',')])
			} else { // no comma: int = whole string
				int1, _ = strconv.Atoi(pack1[index1:])
			}

			// compare list content and int
			cmp := comparePackets(pack0[index0+1:index0+end0], strconv.Itoa(int1))

			if cmp != 0 { // result found
				return cmp
			} else { // same list and int (converted to list), keep scanning (jump to end of int and end of list)
				index0 = index0 + end0 + 1
				index1 += len(strconv.Itoa(int1))
				continue
			}
		}

		if pack1[index1] == '[' { // second el is a list
			end1 := listClosing(pack1[index1:]) // find end of list

			// parse int (first element)
			var int0 int
			if indexOf(pack0[index0:], ',') != -1 { // if contains comma (,)
				int0, _ = strconv.Atoi(pack0[index0 : index0+indexOf(pack0[index0:], ',')])
			} else { // no comma: int = whole string
				int0, _ = strconv.Atoi(pack0[index0:])
			}

			// compare int and list content
			cmp := comparePackets(strconv.Itoa(int0), pack1[index1+1:index1+end1])

			if cmp != 0 { // result found
				return cmp
			} else { // same int (converted to list) and list, keep scanning (jump to end of list and end of int)
				index1 = index1 + end1 + 1
				index0 += len(strconv.Itoa(int0))
				continue
			}
		}
	}

	// end of list, one list is shorted tha the otherone
	if len(pack0) < len(pack1) { // first list shorter
		return 1
	} else { // second list shorter
		return -1
	}
}

// EFFECTS: return index of closing bracket for list (closing ] for first [), -1 if there is no closing bracket
func listClosing(list string) int {
	var open int
	for i, c := range list {
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

// EFFECTS: return the index of rune "r" in string "s", -1 if not found
func indexOf(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}

// REQUIRES: each packet of "packets" is a valid packet
// EFFECTS: returns packets sorted
func sortPackets(packets []string) []string {
	sort.Slice(packets, func(i, j int) bool {
		return comparePackets(packets[i], packets[j]) == 1 // packets[i] < packets[j]
	})
	return packets
}

// EFFECTS: returns the index of packet "packet" in packets "packets"
func findPacketIndex(packets []string, packet string) int {
	for i, v := range packets {
		if v == packet {
			return i + 1
		}
	}
	return -1
}
