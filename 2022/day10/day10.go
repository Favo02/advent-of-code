// https://adventofcode.com/2022/day/10

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	instructions := parseInput()
	signalPower := signalPower(instructions)
	crt := generateCRT(instructions)

	fmt.Println("sum of signals powers (part1):", signalPower)
	fmt.Print("crt display (part2):\n", crt, "\n")
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

func signalPower(instructions []string) int {
	var x, clock int                                 // CPU values
	x = 1                                            // x register initialized to 1
	var sumSignalPower int                           // sum of signal powers at breakpoints
	breakpoints := []int{20, 60, 100, 140, 180, 220} // points where to sum signalPower

	for _, ins := range instructions {
		switch ins[:4] {

		case "noop":
			clock++
			// after every clock change check if clock is in breakpoint
			if equalBreakpoint(clock, breakpoints) {
				sumSignalPower += clock * x
			}

		case "addx":
			clock++
			// after every clock change check if clock is in breakpoint
			if equalBreakpoint(clock, breakpoints) {
				sumSignalPower += clock * x
			}

			clock++
			// after every clock change check if clock is in breakpoint
			if equalBreakpoint(clock, breakpoints) {
				sumSignalPower += clock * x
			}

			// clock MUST be checked before changing register x state
			n, _ := strconv.Atoi(ins[5:])
			x += n

		}
	}

	return sumSignalPower
}

func equalBreakpoint(x int, breakpoints []int) bool {
	for _, bp := range breakpoints {
		if bp == x {
			return true
		}
	}
	return false
}

func generateCRT(instr []string) string {
	var x, clock int     // CPU values
	x = 1                // x register initialized to 1
	var CRTscreen string // "ASCII art" displayed by CRT

	for _, ins := range instr {

		// print on CRT if printing position (clock) contains x (clock = x+-1)
		if math.Abs(float64(clock%40-x)) < 2 {
			CRTscreen += "#"
		} else {
			CRTscreen += "."
		}

		switch ins[:4] {

		case "noop":
			clock++

		case "addx":
			clock++

			// special checks after clock increase

			// if CRT dislay line length reached, go to new line
			if clock%40 == 0 {
				CRTscreen += "\n"
			}

			// print on CRT if printing position (clock) contains x (clock = x+-1)
			if math.Abs(float64(clock%40-x)) < 2 {
				CRTscreen += "#"
			} else {
				CRTscreen += "."
			}

			clock++

			// change register x state
			n, _ := strconv.Atoi(ins[5:])
			x += n
		}

		// if CRT dislay line length reached, go to new line
		if clock%40 == 0 {
			CRTscreen += "\n"
		}
	}

	return CRTscreen
}
