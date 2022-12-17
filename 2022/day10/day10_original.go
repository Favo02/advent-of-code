// https://adventofcode.com/2022/day/10

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	instr := parseInput()
	sum := signalPower(instr)
	fmt.Println(sum)
	crt := checkCTR(instr)
	fmt.Println(crt)
}

func parseInput() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func signalPower(instr []string) int {
	var x, clock, sum int
	x = 1
	for _, ins := range instr {
		if ins[:4] == "noop" {
			clock++
			if clock == 20 || clock == 60 || clock == 100 || clock == 140 || clock == 180 || clock == 220 {
				fmt.Println(clock)
				sum += clock * x
				fmt.Println(clock * x)
				fmt.Println(clock, x)
			}
		} else if ins[:4] == "addx" {
			clock++
			if clock == 20 || clock == 60 || clock == 100 || clock == 140 || clock == 180 || clock == 220 {
				fmt.Println(clock)
				sum += clock * x
				fmt.Println(clock * x)
				fmt.Println(clock, x)
			}
			clock++
			if clock == 20 || clock == 60 || clock == 100 || clock == 140 || clock == 180 || clock == 220 {
				fmt.Println(clock)
				sum += clock * x
				fmt.Println(clock * x)
				fmt.Println(clock, x)
			}
			n, _ := strconv.Atoi(ins[5:])
			x += n
		}
	}
	return sum
}

func checkCTR(instr []string) string {
	var x, clock int
	var res string
	x = 1
	for _, ins := range instr {
		if clock%40 == x-1 || clock%40 == x || clock%40 == x+1 {
			fmt.Println(clock%40, x)
			res += "#"
			fmt.Println("#")
		} else {
			fmt.Println(clock%40, x)
			res += "."
			fmt.Println(".")
		}
		if ins[:4] == "noop" {
			clock++
			if clock == 40 || clock == 80 || clock == 120 || clock == 160 || clock == 200 || clock == 240 {
				res += "\n"
			}
		} else if ins[:4] == "addx" {
			clock++
			if clock == 40 || clock == 80 || clock == 120 || clock == 160 || clock == 200 || clock == 240 {
				res += "\n"
			}

			if clock%40 == x-1 || clock%40 == x || clock%40 == x+1 {
				fmt.Println(clock%40, x)
				res += "#"
				fmt.Println("#")
			} else {
				fmt.Println(clock%40, x)
				res += "."
				fmt.Println(".")
			}

			clock++
			if clock == 40 || clock == 80 || clock == 120 || clock == 160 || clock == 200 || clock == 240 {
				res += "\n"
			}

			n, _ := strconv.Atoi(ins[5:])
			x += n
		}
	}
	return res
}
