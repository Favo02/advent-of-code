// https://adventofcode.com/2022/day/20
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Number struct {
	value int
	index int
}

const DECRYPTION_KEY int = 811589153

func main() {
	nums_p1 := parseInput()       // starting array for part1
	nums_p2 := copyArray(nums_p1) // starting array for part2

	// part1

	// order numbers 1 time
	orderedPart1 := orderNumbers(nums_p1, 1)
	// get 1000th, 2000th, 3000th after 0
	n1000_p1, n2000_p1, n3000_p1 := get10002000300(orderedPart1)
	// multiply 1000th, 2000th, 3000th values after 0 value
	part1 := orderedPart1[n1000_p1].value + orderedPart1[n2000_p1].value + orderedPart1[n3000_p1].value

	// part2

	// multiply decryption key
	nums_p2 = applyDecryptionKey(nums_p2)
	// order numbers 10 times
	orderedPart2 := orderNumbers(nums_p2, 10)
	// get 1000th, 2000th, 3000th after 0
	n1000_p2, n2000_p2, n3000_p2 := get10002000300(orderedPart2)
	// multiply 1000th, 2000th, 3000th values after 0 value
	part2 := orderedPart2[n1000_p2].value + orderedPart2[n2000_p2].value + orderedPart2[n3000_p2].value

	fmt.Println("No decryption key, 1 mixing (part1):\n\t", part1)
	fmt.Println("Decryption key, 10 mixing  (part2):\n\t", part2)
}

// returns the numbers parsed in structs containing value and index
// modifies stdin
func parseInput() (nums []Number) {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		nums = append(nums, Number{n, i})
	}
	return nums
}

// returns nums ordered ("times" times) by moving each number by its value
func orderNumbers(nums []Number, times int) []Number {
	for i := 0; i < times; i++ { // times sorted
		indexToMove := 0                                // index of the Number to sort
		for nMoved := 0; nMoved < len(nums); nMoved++ { // numbers moved
			for i := 0; i < len(nums); i++ { // find number to sort
				if nums[i].index == indexToMove { // current number to sort found
					nums = moveNumber(nums, i, nums[i].value) // move number by its value
					indexToMove++                             // sort next number
					break
				}
			}
		}
	}
	return nums
}

// returns nums with the number at position "index" moved by "value"
func moveNumber(nums []Number, index, value int) []Number {
	if value == 0 { // no need to move
		return nums
	}

	swap := reflect.Swapper(nums)
	nSwaps := value % (len(nums) - 1) // number of swap to reach final position

	if value > 0 { // move forward
		for i := 0; i < nSwaps; i++ {
			// move number with next number, for "nSwaps" times
			swap((index+i)%len(nums), (index+i+1)%len(nums))
		}
	} else { // move backward
		for i := 0; i < -nSwaps; i++ {
			// move number with number before, for "nSwaps" times
			swap(((index-i)%len(nums)+len(nums))%len(nums), ((index-i-1)%len(nums)+len(nums))%len(nums))
		}
	}
	return nums
}

// returns the representation of nums, without indexes
func printNums(nums []Number) (res string) {
	for _, n := range nums {
		res = fmt.Sprint(res, n.value, " ")
	}
	return res
}

// returns indexes of positions 1000, 2000 and 3000 after value 0
func get10002000300(nums []Number) (int, int, int) {
	var index int
	for i, n := range nums { // search 0
		if n.value == 0 {
			index = i
		}
	}
	// calculate positions in circular array
	return (index + 1000) % len(nums), (index + 2000) % len(nums), (index + 3000) % len(nums)
}

// returns a copy of "nums"
func copyArray(nums []Number) []Number {
	copy := make([]Number, len(nums))
	for i, v := range nums {
		copy[i] = v
	}
	return copy
}

// returns nums modified multiplying each Number value with DECRYPTION_KEY
func applyDecryptionKey(nums []Number) []Number {
	for i := 0; i < len(nums); i++ {
		nums[i].value *= DECRYPTION_KEY
	}
	return nums
}
