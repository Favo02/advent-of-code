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
	lines := parseInput()

	sum := 0
	for _, num := range lines {
		sum += snafuToDec(num)
	}
	fmt.Println(sum)

	snafu := decToSnafu(sum)
	fmt.Println(snafu)

}

func parseInput() (lines []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func snafuToDec(num string) (res int) {
	for i := 0; i < len(num); i++ {
		var digit int
		if num[i] == '=' {
			digit = -2
		} else if num[i] == '-' {
			digit = -1
		} else {
			digit, _ = strconv.Atoi(string(rune(num[i])))
		}
		res += digit * pow(5, len(num)-i-1)
	}
	return res
}

func decToSnafu(num int) (res string) {
	var rem []int
	var overflow int
	for num != 0 {
		remDigit := (num % 5) + overflow

		if remDigit > 2 {
			overflow = 1
		} else {
			overflow = 0
		}

		rem = append(rem, remDigit)
		num /= 5
	}

	// missing final overflow
	if overflow > 0 {
		rem = append(rem, overflow)
	}

	for i := len(rem) - 1; i >= 0; i-- {
		digit := rem[i]
		if rem[i] <= 2 {
			res = fmt.Sprint(res, digit)
		} else if rem[i] == 3 {
			res = fmt.Sprint(res, "=")
		} else if rem[i] == 4 {
			res = fmt.Sprint(res, "-")
		} else {
			res = fmt.Sprint(res, "0")
		}
	}
	return res
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
