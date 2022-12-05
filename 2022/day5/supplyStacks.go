// ----------------------------
// NOT REFACTORED YET
// ----------------------------

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
	parseInput(true)
	printStacks()
	msg := readMessage()
	fmt.Println(msg)
}

func parseInput(multipleMoveCrane bool) {
	var startingLines stack

	scanner := bufio.NewScanner(os.Stdin)

	// parse starting stack: save starting lines
	for scanner.Scan() {

		line := scanner.Text()
		if line == "" {
			break
		}
		startingLines.push(line)

	}

	// convert starting lines to stacks

	// get number of stacks (line of numbers)
	stacksNumberStr := startingLines.pop()
	stacksNumberTokens := strings.Split(stacksNumberStr, " ")
	stacksNum, _ := strconv.Atoi(stacksNumberTokens[len(stacksNumberTokens)-2])

	// insert in stacks
	stacks = make([]stack, stacksNum)
	for !startingLines.isEmpty() {
		line := startingLines.pop()

		payloadSize := 0
		stackIndex := 0
		payload := ""
		for i := 0; i < len(line); i++ {
			payload += string(line[i])
			payloadSize++
			if payloadSize == 3 {
				if payload != "   " {
					stacks[stackIndex].push(payload)
				}
				payload = ""
				payloadSize = 0
				stackIndex++
				i++ //skip space
			}
		}
	}

	// parse instructions
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		tokens := strings.Split(line, " ")
		amount, _ := strconv.Atoi(tokens[1])
		from, _ := strconv.Atoi(tokens[3])
		to, _ := strconv.Atoi(tokens[5])

		moveStacks(amount, from-1, to-1, multipleMoveCrane)
	}
}

// multiple true =
func moveStacks(amount, from, to int, multiple bool) {
	if !multiple {
		for i := 0; i < amount; i++ {
			fromVal := stacks[from].pop()
			stacks[to].push(fromVal)
		}
	} else {
		var buffer stack
		for i := 0; i < amount; i++ {
			buffer.push(stacks[from].pop())
		}
		for i := 0; i < amount; i++ {
			stacks[to].push(buffer.pop())
		}
	}
}

func readMessage() string {
	var msg string
	for i := 0; i < len(stacks); i++ {
		msg += stacks[i].pop()
	}
	msg = strings.ReplaceAll(msg, "[", "")
	msg = strings.ReplaceAll(msg, "]", "")
	return msg
}

func printStacks() {
	fmt.Println("---")
	for _, s := range stacks {
		s.print()
	}
	fmt.Println("---")
}

// STACK //

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
