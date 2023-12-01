package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// parse input and caculate ones and zeros for part 1
	binValues, ones, zeros := parseInput()

	// part1
	gamma, epsilon := getGammaEpsilon(ones, zeros)

	gammaDec := binToDec(gamma)
	epsilonDec := binToDec(epsilon)

	fmt.Println("power consumption (gamma * epsilon):")
	fmt.Println("\tpart1 =", gammaDec*epsilonDec)

	// part2
	oxygen := getValid(binValues, gamma, 0, '1')
	co2 := getValid(binValues, epsilon, 0, '0')

	oxygenDec := binToDec(oxygen)
	co2Dec := binToDec(co2)

	fmt.Println("life support (o2 * co2):")
	fmt.Println("\tpart2 =", oxygenDec*co2Dec)
}

func parseInput() ([]string, []int, []int) {
	scanner := bufio.NewScanner(os.Stdin)

	var binary []string
	var ones []int
	var zeros []int

	for scanner.Scan() {
		line := scanner.Text()
		binary = append(binary, line)

		for i, bit := range line {
			if i >= len(zeros) {
				zeros = append(zeros, 0)
			}
			if i >= len(ones) {
				ones = append(ones, 0)
			}
			if bit == '0' {
				zeros[i]++
			} else if bit == '1' {
				ones[i]++
			}
		}
	}

	return binary, ones, zeros
}

func getOnesZeros(values []string) ([]int, []int) {
	var ones []int
	var zeros []int

	for _, val := range values {
		for i, bit := range val {
			if i >= len(zeros) {
				zeros = append(zeros, 0)
			}
			if i >= len(ones) {
				ones = append(ones, 0)
			}
			if bit == '0' {
				zeros[i]++
			} else if bit == '1' {
				ones[i]++
			}
		}
	}

	return ones, zeros
}

func getGammaEpsilon(ones, zeros []int) (string, string) {
	var gamma, epsilon string

	for i := 0; i < len(ones); i++ {
		if ones[i] > zeros[i] {
			gamma += "1"
			epsilon += "0"
		} else if zeros[i] > ones[i] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "-"
			epsilon += "-"
		}
	}
	return gamma, epsilon
}

func binToDec(binary string) int {
	dec, _ := strconv.ParseInt(binary, 2, 64)
	return int(dec)
}

// filter values: value[indexToFilter] == criteria[indexToFilter]
// if criteria == - then value[indexToFilter] == takeOnEqual
func getValid(values []string, criteria string, indexToFilter int, takeOnEqual rune) string {
	if len(values) <= 1 {
		return values[0]
	}

	var validValues []string

	for _, val := range values {
		if criteria[indexToFilter] == '-' {
			if val[indexToFilter] == byte(takeOnEqual) {
				validValues = append(validValues, val)
			}
		} else if val[indexToFilter] == criteria[indexToFilter] {
			validValues = append(validValues, val)
		}
	}

	ones, zeros := getOnesZeros(validValues)
	if takeOnEqual == '1' {
		criteria, _ = getGammaEpsilon(ones, zeros)
	} else {
		_, criteria = getGammaEpsilon(ones, zeros)
	}

	return getValid(validValues, criteria, indexToFilter+1, takeOnEqual)
}
