package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(parseInput())
}

func parseInput() (int, int) {
	var wrongScore int
	var correctScore int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		w, c := lineToScore(line)
		wrongScore += w
		correctScore += c
	}

	return wrongScore, correctScore
}

// A = rock
// B = paper
// C = scissors

// X = rock			= 1
// Y = paper		= 2
// Z = scissors	= 3

// X = lose
// Y = draw
// Z = win

func lineToScore(line string) (int, int) {
	tokens := strings.Split(line, " ")
	opponent := tokens[0]
	me := tokens[1]

	var wrongScore, correctScore int

	wrongScore += meToScore(me)
	wrongScore += calculateWin(me, opponent)

	me = calculateResponse(opponent, me)
	correctScore += meToScore(me)
	correctScore += calculateWin(me, opponent)

	return wrongScore, correctScore
}

func meToScore(me string) int {
	switch me {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}
	return -1
}

// 6 = win
// 3 = draw
// 0 = lose

func calculateWin(me, opponent string) int {
	switch me {

	case "X": // rock
		switch opponent {
		case "A": // rock
			return 3
		case "B": // paper
			return 0
		case "C": // scissors
			return 6
		}

	case "Y": // paper
		switch opponent {
		case "A": // rock
			return 6
		case "B": // paper
			return 3
		case "C": // scissors
			return 0
		}

	case "Z": // scissors
		switch opponent {
		case "A": // rock
			return 0
		case "B": // paper
			return 6
		case "C": // scissors
			return 3
		}

	}
	return -1
}

func calculateResponse(opponent, result string) string {
	switch result {

	case "X": // lose
		switch opponent {
		case "A": // rock
			return "Z" // scissors
		case "B": // paper
			return "X" // rock
		case "C": // scissors
			return "Y" // paper
		}

	case "Y": // draw
		switch opponent {
		case "A": // rock
			return "X" // rock
		case "B": // paper
			return "Y" // paper
		case "C": // scissors
			return "Z" // scissors
		}

	case "Z": // win
		switch opponent {
		case "A": // rock
			return "Y" // paper
		case "B": // paper
			return "Z" // scissors
		case "C": // scissors
			return "X" //rock
		}

	}
	return "A"
}
