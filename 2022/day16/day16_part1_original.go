// https://adventofcode.com/2022/day/16
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Valve struct {
	name   string
	flow   int
	isOpen bool
}

func main() {
	graph := parseInput()
	// cur := Valve{"AA", 0, false}

	distances := make(map[Valve]map[Valve]int)
	for k := range graph {
		dist, _ := depthFirstSearch(graph, k)
		distances[k] = dist
	}

	for k, v := range distances {
		fmt.Println(k, ":")
		fmt.Println(v)
		fmt.Println("---")
	}

	notVisided := make(map[Valve]bool)
	for k := range distances {
		if k.name == "AA" {
			continue
		}
		notVisided[k] = true
	}

	fmt.Println(maxPressure(notVisided, Valve{"AA", 0, false}, distances, 30))
}

func parseInput() map[Valve][]string {
	scanner := bufio.NewScanner(os.Stdin)

	graph := make(map[Valve][]string)
	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		name := tokens[1]
		flowStr := tokens[4][5:]
		flowStr = flowStr[:len(flowStr)-1]
		flow, _ := strconv.Atoi(flowStr)
		var dest []string
		for i := 9; i < len(tokens); i++ {
			if len(tokens[i]) == 3 {
				dest = append(dest, tokens[i][:2])
			} else {
				dest = append(dest, tokens[i])
			}
		}

		valve := Valve{name, flow, false}
		graph[valve] = dest
	}
	return graph
}

func findByName(graph map[Valve][]string, name string) Valve {
	for k := range graph {
		if k.name == name {
			return k
		}
	}
	return Valve{}
}

func depthFirstSearch(graph map[Valve][]string, cur Valve) (map[Valve]int, map[Valve][]Valve) {
	queue := queue{nil}
	distances := make(map[Valve]int)
	distances[cur] = 0
	reached := make(map[Valve]bool)
	reached[cur] = true
	visitedToReach := make(map[Valve][]Valve)

	queue.enqueue(cur)

	for !queue.isEmpty() {
		u := queue.dequeue()

		for _, vString := range graph[u] {
			v := findByName(graph, vString)
			if !reached[v] {
				distances[v] = distances[u] + 1
				visitedToReach[v] = visitedToReach[u]
				visitedToReach[v] = append(visitedToReach[v], u)
				reached[v] = true
				queue.enqueue(v)
			}
		}

	}
	return distances, visitedToReach
}

func maxPressure(notVisited map[Valve]bool, cur Valve, distances map[Valve]map[Valve]int, time int) int {
	if len(notVisited) == 0 || time <= 0 {
		return 0
	}

	var totPressure int
	max := 0
	for k := range notVisited {
		if k.flow == 0 {
			continue
		}
		pressure := (time - (distances[cur][k] + 1)) * k.flow
		totPressure = pressure + maxPressure(reducedMap(notVisited, k), k, distances, time-(distances[cur][k]+1))
		if totPressure > max {
			max = totPressure
		}
	}

	return max
}

func reducedMap(mapp map[Valve]bool, remove Valve) map[Valve]bool {
	res := make(map[Valve]bool)
	for k, v := range mapp {
		if k != remove {
			res[k] = v
		}
	}
	return res
}

// QUEUE

type queue struct {
	head *queueNode
}

type queueNode struct {
	next    *queueNode
	payload Valve
}

func (q *queue) enqueue(p Valve) {
	if q.head == nil {
		q.head = &queueNode{nil, p}
		return
	}
	node := q.head
	for node.next != nil {
		node = node.next
	}
	newNode := queueNode{nil, p}
	node.next = &newNode
}

func (q *queue) dequeue() Valve {
	head := q.head
	q.head = q.head.next
	return head.payload
}

func (q *queue) isEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}
