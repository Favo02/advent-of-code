// https://adventofcode.com/2022/day/2
// rock paper scissors - implemented using circular arrays

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wrongScore, correctScore := parseInput()
	fmt.Println("wrong score (part 1):", wrongScore)
	fmt.Println("correct score (part 2):", correctScore)
}

// REQUIRES: stdin is a valid puzzle input, terminated by EOF
// MODIFIES: stdin
// EFFECTS: calculates and returns the wrongScore and the correctScore
func parseInput() (int, int) {
	var wrongScore int   // score based on part 1 of the challenge (X, Y, Z = rock, paper, scissors)
	var correctScore int // score based on part 2 of the challenge (X, Y, Z = lose, draw, win)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		wrongScoreLine, correctScoreLine := lineToScore(line)
		wrongScore += wrongScoreLine
		correctScore += correctScoreLine
	}

	return wrongScore, correctScore
}

// REQUIRES: line is in format "X Y", with X = A, B, C and Y = X, Y, Z
// EFFECTS: calculates and returns the wrongScore and the correctScore for line
func lineToScore(line string) (int, int) {
	tokens := strings.Split(line, " ")
	opponentStr := tokens[0]
	meStr := tokens[1]

	opponent := stringCodeToIntCode(opponentStr)
	me := stringCodeToIntCode(meStr)

	var wrongScore, correctScore int

	// me is already the move to be performed
	wrongScore += me
	wrongScore += calculateWin(me, opponent)

	// me is (win, lose or draw), not the move to be performed
	myMove := calculateMove(opponent, meStr)

	// myMove is the move to be performed
	correctScore += myMove
	correctScore += calculateWin(myMove, opponent)

	return wrongScore, correctScore
}

// REQUIRES: me = X, Y or Z
// EFFECTS: returns 1 for rock, 2 for paper, 3 for scissors, -1 otherwise
func stringCodeToIntCode(strCode string) int {
	// X = A = rock			= 1
	// Y = B = paper		= 2
	// Z = C = scissors	= 3

	switch strCode {
	case "X", "A": // rock
		return 1
	case "Y", "B": // paper
		return 2
	case "Z", "C": // scissors
		return 3
	}
	return -1
}

// REQUIRES: me, opponent = 1, 2, 3 (code for rock, paper, scissors)
// EFFECTS: returns 6 if me wins, 3 if me and opponents draw, 0 if opponent wins
func calculateWin(me, opponent int) int {
	// 6 = win
	// 3 = draw
	// 0 = lose
	if me == opponent {
		return 3
	}
	// convert codes to circular array:
	// convert code from 1,2,3 to 0,1,2
	me--
	opponent--
	if (me+1)%3 == opponent {
		return 0
	} else {
		return 6
	}
}

// REQUIRES: opponent = 1, 2, 3 (code for rock, paper, scissors), result = X, Y, Z (lose, draw, lose)
// EFFECTS: returns the move needed to get the result desiderated by result parameter
func calculateMove(opponent int, result string) int {
	switch result {
	case "X": // lose
		move := opponent - 1
		if move == 0 { // outside of circular array
			move = 3
		}
		return move
	case "Y": // draw
		return opponent
	case "Z": // win
		move := opponent + 1
		if move == 4 { // outside of circular array
			move = 1
		}
		return move
	}
	return 0
}
