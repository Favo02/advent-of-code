// https://adventofcode.com/2022/day/25
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	snafuNums := parseInput()             // read input
	decNums := snafuToDecArray(snafuNums) // convert snafu to dec
	decSum := sumArray(decNums)           // sum dec array
	snafu := decToSnafu(decSum)           // convert sum to snafu
	fmt.Println("sum of all snafu numbers in snafu base (part1):\n\t", snafu)
}

// returns numbers parse from stdin
// modifies stdin
func parseInput() (nums []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		nums = append(nums, line)
	}
	return nums
}

// returns snafu numbers converted to decimal numbers
func snafuToDecArray(snafu []string) (dec []int) {
	dec = make([]int, len(snafu))
	for i, s := range snafu {
		dec[i] = snafuToDec(s)
	}
	return dec
}

// returns the snafu number "num" converted to decimal
func snafuToDec(snafu string) (dec int) {
	for i := 0; i < len(snafu); i++ {
		var digit int
		if snafu[i] == '=' {
			digit = -2
		} else if snafu[i] == '-' {
			digit = -1
		} else {
			digit, _ = strconv.Atoi(string(rune(snafu[i])))
		}
		dec += digit * pow(5, len(snafu)-i-1)
	}
	return dec
}

// returns the sum of all numbers in "nums"
func sumArray(nums []int) (sum int) {
	for _, n := range nums {
		sum += n
	}
	return sum
}

// returns decimal number "dec" converted to snafu
func decToSnafu(dec int) (snafu string) {
	var snafuDigits []int
	var overflow int
	for dec != 0 {
		remDigit := (dec % 5) + overflow

		if remDigit > 2 {
			overflow = 1
		} else {
			overflow = 0
		}

		snafuDigits = append(snafuDigits, remDigit)
		dec /= 5
	}

	// still overflow after number is over
	if overflow > 0 {
		snafuDigits = append(snafuDigits, overflow)
	}

	// parse array of digits to string, in reverse order
	for i := len(snafuDigits) - 1; i >= 0; i-- {
		digit := snafuDigits[i]
		if snafuDigits[i] <= 2 {
			snafu = fmt.Sprint(snafu, digit)
		} else if snafuDigits[i] == 3 {
			snafu = fmt.Sprint(snafu, "=")
		} else if snafuDigits[i] == 4 {
			snafu = fmt.Sprint(snafu, "-")
		} else {
			snafu = fmt.Sprint(snafu, "0")
		}
	}
	return snafu
}

// returns a^b (math.Pow(a, b)) as int
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
