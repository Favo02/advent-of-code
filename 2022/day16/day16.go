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

type Person struct {
	valve Valve
	time  int
}

func main() {
	graph, nodes := parseInput()              // parse valves and connections between valves as a graph and all nodes
	distances := calculateAllDistances(graph) // distance from each valve to every valve

	startingValve := Valve{"AA", 0, false}                          // starting valve
	time := 30                                                      // time to open valves
	part1 := maxPressureSolo(nodes, startingValve, distances, time) // calculate max pressure releasable in "time" time

	fmt.Println("Max pressure releasable, only me (part1):\n\t", part1)

	time = 26                                                    // time to open valves
	me := Person{startingValve, time}                            // person starting point and time remaining
	elefant := Person{startingValve, time}                       // person starting point and time remaining
	part2 := maxPressureDuo(nodes, distances, me, elefant, true) // calculate max pressure releasable in "time" time, with me and elefant working

	fmt.Println("Max pressure releasable, me and elefant (part2):\n\t", part2)
}

// returns a graph representing each valve with the possible valves reachable in 1 step
func parseInput() (map[Valve][]string, []Valve) {
	scanner := bufio.NewScanner(os.Stdin)

	graph := make(map[Valve][]string) // graph: each valve with possible valve to reach in 1 step
	var nodes []Valve                 // all nodes
	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " ") // split line

		name := tokens[1]                  // valve name
		flowStr := tokens[4][5:]           // valve flow (ending with ;)
		flowStr = flowStr[:len(flowStr)-1] // remove final ; from flow
		flow, _ := strconv.Atoi(flowStr)   // parse flow (string) to in

		var dest []string                  // destinations reachable from this valve
		for i := 9; i < len(tokens); i++ { // scan destinations
			if len(tokens[i]) == 3 { // has final ;
				dest = append(dest, tokens[i][:2]) // add destination removing final ;
			} else { // no final ;
				dest = append(dest, tokens[i]) // add destination
			}
		}

		valve := Valve{name, flow, false} // create valve object
		graph[valve] = dest               // add valve and destinations to graph
		nodes = append(nodes, valve)      // add valve to nodes list
	}
	return graph, nodes
}

// returns the distance from each valve to every valve
func calculateAllDistances(graph map[Valve][]string) map[Valve]map[Valve]int {
	distances := make(map[Valve]map[Valve]int) // each valve is associated with every valve associated with the distance
	for k := range graph {                     // scan each valve
		dist := depthFirstSearch(graph, k) // get distance to each valve for current valve using a dfs
		distances[k] = dist                // set distances
	}
	return distances
}

// returns the distance from each valve for "cur" valve
func depthFirstSearch(graph map[Valve][]string, cur Valve) map[Valve]int {
	queue := queue{nil}              // nodes in queue
	distances := make(map[Valve]int) // distances calculated
	distances[cur] = 0               // distance to starting node = 0
	reached := make(map[Valve]bool)  // reached nodes
	reached[cur] = true              // starting node reached

	queue.enqueue(cur) // add to analyze queue starting node

	for !queue.isEmpty() {
		u := queue.dequeue() // element to analyze

		for _, vString := range graph[u] { // each nodes reachable from u (as string)
			v := findByName(graph, vString) // get node analyzed (as Valve)

			if !reached[v] { // not yet analyzed
				distances[v] = distances[u] + 1 // calculate distance
				reached[v] = true               // set reached
				queue.enqueue(v)                // add to nodes to analyze
			}
		}
	}
	return distances
}

// returns the Valve with name "name"
func findByName(graph map[Valve][]string, name string) Valve {
	for k := range graph { // scan each Valve
		if k.name == name {
			return k
		}
	}
	return Valve{} // not found, return empty Valve
}

// returns max pressure releasable working solo, in time "time"
func maxPressureSolo(notVisited []Valve, cur Valve, distances map[Valve]map[Valve]int, time int) int {
	if len(notVisited) == 0 || time <= 0 { // no valves to open or time is over
		return 0
	}

	var totPressure int // pressure released
	var maxPressure int // maximum pressure

	for _, k := range notVisited { // scan each node
		if k.flow == 0 { // no flow valve, skip
			continue
		}

		timeRemeaning := time - (distances[cur][k] + 1)
		pressure := timeRemeaning * k.flow // calculate pressure releasable from this valve only over time

		// calculate pressure released from next valves recursively
		nextValvesPress := maxPressureSolo(removeFromSlice(notVisited, k), k, distances, time-(distances[cur][k]+1))
		totPressure = pressure + nextValvesPress // total pressure: this valve + next nodes

		// save only max pressure
		if totPressure > maxPressure {
			maxPressure = totPressure
		}
	}
	return maxPressure
}

// returns max pressure releasable working as a sue, in time "time"
func maxPressureDuo(notVisited []Valve, distances map[Valve]map[Valve]int, p1 Person, p2 Person, alt bool) int {

	// no valves to open remeaning
	if len(notVisited) == 0 {
		return 0
	}

	var cur Valve // current valve
	var time int  // current time remeaning

	// alternate person values
	if alt { // set valve and time from person 1
		cur = p1.valve
		time = p1.time
	} else { // set valve and time from person 2
		cur = p2.valve
		time = p2.time
	}

	var totPressure int // pressure released
	var maxPressure int // maximum pressure

	for _, k := range notVisited { // scan each node
		if k.flow == 0 { // no flow valve, skip
			continue
		}

		if (time - (distances[cur][k] + 1)) < 0 { // not enough time to reach node, skip
			continue
		}

		pressure := (time - (distances[cur][k] + 1)) * k.flow // calculate pressure releasable from this valve only over time

		// alternate this node results
		if alt { // results belongs to person 1
			p1.valve = k
			p1.time = time - (distances[cur][k] + 1)
		} else { // results belongs to person 2
			p2.valve = k
			p2.time = time - (distances[cur][k] + 1)
		}

		// calculate pressure released from next valves recursively
		nextValvesPress := maxPressureDuo(removeFromSlice(notVisited, k), distances, p1, p2, !alt)
		totPressure = pressure + nextValvesPress // total pressure: this valve + next nodes

		// save only max pressure
		if totPressure > maxPressure {
			maxPressure = totPressure
		}
	}
	return maxPressure
}

// returns a slice without "remove" item
func removeFromSlice(slice []Valve, remove Valve) (res []Valve) {
	for _, v := range slice {
		if v != remove {
			res = append(res, v)
		}
	}
	return res
}

// QUEUE used in bfs

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
