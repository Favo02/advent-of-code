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
	fmt.Print("CRT display (part2):\n", crt, "\n")
}

// REQUIRES: stdin is a valid challenge input
// MODIFIES: stdin
// EFFECTS: returns the instructions
func parseInput() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

// REQUIRES: instructions contains valid challenge instructions
// EFFECTS: returns the sum of signals powers at clock cycle 20, 60, 100, 140, 180, 220
func signalPower(instructions []string) int {
	var x, clock int                                 // CPU values
	x = 1                                            // x register initialized to 1
	var sumSignalPower int                           // sum of signal powers at breakpoints
	breakpoints := []int{20, 60, 100, 140, 180, 220} // points where to sum signalPower

	for _, instruction := range instructions {
		switch instruction[:4] {

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
			n, _ := strconv.Atoi(instruction[5:])
			x += n

		}
	}

	return sumSignalPower
}

// EFFECTS: returns true if x is in breakpoints, false otherwise
func equalBreakpoint(x int, breakpoints []int) bool {
	for _, bp := range breakpoints {
		if bp == x {
			return true
		}
	}
	return false
}

// REQUIRES: instructions contains valid challenge instructions
// EFFECTS: returns the visualization of CRT display after the execution of instructions
func generateCRT(instructions []string) string {
	var x, clock int     // CPU values
	x = 1                // x register initialized to 1
	var CRTscreen string // "ASCII art" displayed by CRT

	for _, instruction := range instructions {

		// print character on screen ("█" or " ")
		CRTscreen = printCRTcharacter(clock, x, CRTscreen)

		switch instruction[:4] {

		case "noop":
			clock++

		case "addx":
			clock++

			// special checks after clock increase

			// print newline on screen if CRT dislay line length reached
			CRTscreen = checkCRTnewline(clock, CRTscreen)

			// print character on screen ("█" or " ")
			CRTscreen = printCRTcharacter(clock, x, CRTscreen)

			clock++

			// change register x state
			n, _ := strconv.Atoi(instruction[5:])
			x += n
		}

		// print newline on screen if CRT dislay line length reached
		CRTscreen = checkCRTnewline(clock, CRTscreen)

	}

	return CRTscreen
}

// EFFECTS: returns CRTscreen updated with "█" if printing position (clock) contains x (clock = x+-1), updated with " " otherwise
func printCRTcharacter(clock, x int, CRTscreen string) string {
	if math.Abs(float64(clock%40-x)) < 2 {
		CRTscreen += "█"
	} else {
		CRTscreen += " "
	}
	return CRTscreen
}

// EFFECTS: returns CRTscreen updated with a new line if CRT display line length reached (each 40 characters)
func checkCRTnewline(clock int, CRTscreen string) string {
	if clock%40 == 0 {
		CRTscreen += "\n"
	}
	return CRTscreen
}
