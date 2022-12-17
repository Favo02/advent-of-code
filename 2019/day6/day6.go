package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	vals, parentsIndex := parseInput()

	//fmt.Println(vals)
	//fmt.Println(parentsIndex)

	var total int
	for nodeIndex := range vals {
		deepness, _ := visitNodes(vals, parentsIndex, nodeIndex)
		total += deepness
	}

	fmt.Println("Total number of orbits:", total)

	// paths from you and san to root
	var youToRoot, sanToRoot []string

	youIndex := findIndex(vals, "YOU")
	if youIndex != -1 { // calculate path only if you exists
		_, youToRoot = visitNodes(vals, parentsIndex, youIndex)
	}

	sanIndex := findIndex(vals, "SAN")
	if sanIndex != -1 { // calculate path only if san exists
		_, sanToRoot = visitNodes(vals, parentsIndex, sanIndex)
	}

	// calculate distance between you and san (only if both exists)
	if youIndex != -1 && sanIndex != -1 {
		iy, is := comparePaths(youToRoot, sanToRoot) // start of common path
		
		//fmt.Println("Start common path indexes:", iy, is)
		//fmt.Println("Common path to root:", youToRoot[iy:])
		//fmt.Println("Commont path to root", sanToRoot[is:])

		//fmt.Println("PATH FROM YOU TO FIRST COMMON POINT:\n", youToRoot[:iy])
		//fmt.Println("PATH FROM SAN TO FIRST COMMON POINT:\n", sanToRoot[:is])

		var commonPathLength int = len(youToRoot[:iy]) + len(sanToRoot[:is]) - 2 // remove the two last orbits

		fmt.Println("Path length from YOU to SAN:", commonPathLength)
	}
}

// MODIFIES: standard input
// EFFECTS: returns a tree using two parallel arrays, values array and parents index array
func parseInput() (vals []string, parentsIndex []int) {
	parents := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ")")

		// if parent already defined save his index
		parentIndex := findIndex(vals, tokens[0])
		if (parentIndex != -1) {
			parentsIndex = append(parentsIndex, parentIndex)
		} else { // parent not yet defined
			parentsIndex = append(parentsIndex, -1)
		}
		parents = append(parents, tokens[0])
		vals = append(vals, tokens[1])
	}

	// fix indexes of parents defined after their childs
	for i, parent := range parentsIndex {
		if parent == -1 {
			parentsIndex[i] = findIndex(vals, parents[i])
		}
	}
	return vals, parentsIndex
}

// REQUIRES: vals and parentsIndex are a tree
// EFFECTS: returns index of val in values
func findIndex(vals []string, val string) int {
	for i, v := range vals {
		if v == val {
			return i
		}
	}
	return -1
}

// REQUIRES: vals and parentsIndex are a tree
// EFFECTS: returns slice of visited node and number of visited nodes from nodeIndex to root
func visitNodes(vals []string, parentsIndex []int, nodeIndex int) (int, []string) {
	nodesVisited := make([]string, 0)
	
	var res int = 1

	pIndex := parentsIndex[nodeIndex]
	for pIndex != -1 {
		nodesVisited = append(nodesVisited, vals[nodeIndex])
		res++
		nodeIndex = parentsIndex[nodeIndex]
		pIndex = parentsIndex[nodeIndex]
	}

	return res, nodesVisited
}

// EFFECTS: returns indexex of you and san where a common value is stored. returns -1, -1 if there are no common values
func comparePaths(you, san []string) (int, int) {
	for iy, y := range you {
		for is, s := range san {
			if y == s {
				return iy, is
			}
		}
	}
	return -1, -1 // no common path
}
