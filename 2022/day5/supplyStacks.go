// https://adventofcode.com/2022/day/5

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stacks []stack

func main() {
	// false --> part1
	// true --> part2
	parseInput(false) // modifies stacks
	// printStacks()
	fmt.Println(readMessage())
}

// REQUIRES: stdin is a valid challenge input
// MODIFIES: stacks, stdin
// EFFECTS: parses input stacks into stacks global var,
// then parses instructions executing them on stacks.
// if multipleMoveCrane = false: moving operationg will move only one crane at a time
// if multipleMoveCrane = true: moving operations will move more than one crane at a time
func parseInput(multipleMoveCrane bool) {
	var startingLines stack

	scanner := bufio.NewScanner(os.Stdin)

	// --- save starting stacks (as strings) ---

	// save starting lines in a stack, I need to push first the lower content (Z, M, P)
	// so I need to "reverse" this part of the input
	//	   [D]
	// [N] [C]
	// [Z] [M] [P]
	//  1   2   3
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // end of starting stacks part
			break
		}
		startingLines.push(line)
	}

	// --- convert starting stacks (as strings) to stacks ---

	// parse number of stacks (line of numbers, last line pushed to stack, first to pop)
	stacksNumberTokens := strings.Split(startingLines.pop(), " ")
	stacksNum, _ := strconv.Atoi(stacksNumberTokens[len(stacksNumberTokens)-2]) // -2 because of final space

	// push in stacks
	stacks = make([]stack, stacksNum)
	for !startingLines.isEmpty() { // keep popping from initial lines
		line := startingLines.pop()

		payloadSize := 0
		stackIndex := 0
		payload := ""
		for i := 0; i < len(line); i++ {
			payload += string(line[i])
			payloadSize++
			if payloadSize == 3 { // every value is 3 characters [X]
				if payload != "   " { // skip empty values
					stacks[stackIndex].push(payload[1:2]) // push without []
				}
				payload = ""
				payloadSize = 0
				stackIndex++
				i++ //skip space separating values
			}
		}
	}

	// --- parse move instructions ---

	for scanner.Scan() {
		line := scanner.Text()

		// extract amount to move (token 1), from stack (token 3) and to stack (token 5):
		// move 2 from 7 to 2
		//   0  1   2  3  4 5
		tokens := strings.Split(line, " ")
		amount, _ := strconv.Atoi(tokens[1])
		from, _ := strconv.Atoi(tokens[3])
		to, _ := strconv.Atoi(tokens[5])

		moveStacks(amount, from-1, to-1, multipleMoveCrane) // perform the move
	}
}

// REQUIRES: amount >= 0, from and to valid stacks indexes
// MODIFIES: stacks
// EFFECTS: perform the move: move "amount" values from "from" to "to"
// if multiple = false: values will be moved only one at a time
// if multiple = true: values will be moved more than ones at a time
func moveStacks(amount, from, to int, multiple bool) {
	if !multiple { // pop from "from" stack and push to "to" stack
		for i := 0; i < amount; i++ { // move "amout" times
			stacks[to].push(stacks[from].pop())
		}
	} else { // pop from "from" stack, save the values in a buffer stack, when the amount is over push to "to" stack popping from the buffer stack
		var buffer stack // buffer stack
		for i := 0; i < amount; i++ {
			buffer.push(stacks[from].pop())
		}
		for i := 0; i < amount; i++ {
			stacks[to].push(buffer.pop())
		}
	}
}

// REQUIRES: every stack is not empty
// MODIFIES: stacks
// EFFECTS: returns the first values of each stack (removing it (pop)) concatenated in a string
func readMessage() string {
	var msg string
	for i := 0; i < len(stacks); i++ {
		msg += stacks[i].pop()
	}
	return msg
}

// REQUIRES: s != nil
// MODIFIES: stdout
// EFFECTS: print to stdout each stack
func printStacks() {
	fmt.Println("---")
	for _, s := range stacks {
		s.print()
	}
	fmt.Println("---")
}

// STACK IMPLEMENTATION //

type listNode struct {
	next *listNode
	item string
}

type stack struct {
	head *listNode
}

func newNode(item string) *listNode {
	return &listNode{nil, item}
}

func (list *stack) push(item string) {
	newNode := newNode(item)
	newNode.next = list.head
	list.head = newNode
}

func (list *stack) pop() string {
	node := list.head
	list.head = node.next
	return node.item
}

func (list stack) isEmpty() bool {
	if list.head == nil {
		return true
	}
	return false
}

func (list stack) print() {
	fmt.Print("[ ")
	var node *listNode = list.head
	for node != nil {
		fmt.Print(node.item, " ")
		node = node.next
	}
	fmt.Println("]")
}
