package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var invalidPoints int
	var missingPoints []int

	for scanner.Scan() {
		line := scanner.Text()
		wellFormed, firstInvalid, missingChunks := checkLine(line)
		if !wellFormed {
			if len(missingChunks) > 0 { // incomplete line
				missingPoints = append(missingPoints, getMissingPoints(missingChunks))
			} else { // invalid line
				invalidPoints += getInvalidPoints(firstInvalid)
			}
		}
	}

	sort.Ints(missingPoints)
	var middleScore int = missingPoints[len(missingPoints)/2]

	fmt.Println("invalidPoints (part1):\t", invalidPoints)
	fmt.Println("middleScore (part2):\t", middleScore)
}

func checkLine(line string) (bool, rune, []rune) {
	var stack stack = stack{nil}
	const opening string = "([{<"
	const closing string = ")]}>"

	for _, c := range line {
		if strings.Contains(opening, string(c)) {
			stack.push(c)
		} else if strings.Contains(closing, string(c)) {
			if getCorresponding(c) != stack.pop() {
				// invalid line
				return false, c, nil
			}
		} else {
			fmt.Println("invalid")
		}
	}

	// valid line
	if stack.isEmpty() {
		return true, '0', nil
	}

	// incomplete line
	var missingChunks = make([]rune, 0)
	for !stack.isEmpty() {
		missingChunks = append(missingChunks, getCorresponding(stack.pop()))
	}
	return false, '0', missingChunks
}

func getCorresponding(chunk rune) rune {
	switch chunk {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	case '>':
		return '<'
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	default:
		fmt.Println("Invalid rune")
		return '0'
	}
}

func getInvalidPoints(chunk rune) int {
	switch chunk {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		fmt.Println("invalid get invalid points")
		return 0
	}
}

func getMissingPoints(chunk []rune) int {
	var res int = 0
	for _, v := range chunk {
		res *= 5
		switch v {
		case ')':
			res += 1
		case ']':
			res += 2
		case '}':
			res += 3
		case '>':
			res += 4
		}
	}
	return res
}

// IMPLEMENTAZIONE PILA TRAMITE LISTA //

type listNode struct {
	next *listNode
	item rune
}

type stack struct {
	head *listNode
}

func newNode(item rune) *listNode {
	return &listNode{nil, item}
}

func (list *stack) push(item rune) {
	newNode := newNode(item)
	newNode.next = list.head
	list.head = newNode
}

func (list *stack) pop() rune {
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
