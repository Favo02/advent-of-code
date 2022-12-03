package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	single, triple := parseInput()
	fmt.Println("sigle value (part1):", single)
	fmt.Println("triple value (part2):", triple)
}

func parseInput() (int, int) {
	var prec, curr int
	var count int

	var curr3sum int
	var last3 [3]int
	var count3 int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// part1
		curr, _ = strconv.Atoi(line)
		if curr > prec {
			count++
		}
		prec = curr

		// part2
		last3sum := sumArray(last3)
		curr3sum = last3sum - last3[0] + curr
		last3 = shiftArray(last3, curr)
		if curr3sum > last3sum {
			count3++
		}
	}

	// count-1: first element has no prec
	// count-3: first 3 elements have no prec
	return count - 1, count3 - 3
}

// shifts left the array, inserting in at last position
func shiftArray(arr [3]int, in int) [3]int {
	arr[0] = arr[1]
	arr[1] = arr[2]
	arr[2] = in
	return arr
}

// sums the array values
func sumArray(arr [3]int) int {
	return arr[0] + arr[1] + arr[2]
}
