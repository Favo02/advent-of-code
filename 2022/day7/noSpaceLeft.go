// https://adventofcode.com/2022/day/7
// filesystem implemented using parent parallel array

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// directory saved as a tree, using a parallel array for parent of each node
	nodes, parents := parseInput()

	// map with each directory size
	dirSizes := calcDirsSizes(nodes, parents)

	// calculate part1 (sum of dirs size smaller than 10000)
	part1 := calculatePart1(dirSizes)
	fmt.Println("Sum of directories smaller than 100000 (part1):\n\t", part1)

	// calculate part2 (size of smallest dir to free enough space)
	part2 := calculatePart2(dirSizes)
	fmt.Println("Smallest directory to delete to free 30000000 (part2):\n\t", part2)
}

// node object
type node struct {
	name  string
	size  int  // size for dir = 0
	type_ rune // d = dir, d = file
}

// parse file system to a tree, using a parallel array for parent of each node
func parseInput() ([]node, []string) {
	var nodes []node
	var parent []string

	// current directory
	var curDir string

	// append root
	nodes = append(nodes, node{"/", 0, 'd'})
	parent = append(parent, "")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line[:1] == "$" { // if command
			command := line[2:4]

			// change current directory
			if command == "cd" {
				dest := line[5:]  // destination
				if dest == ".." { // go back
					curDir = curDir[:findLastSlash(curDir)]
				} else if dest == "/" { // go to root
					curDir = "/"
				} else { // go deeper
					curDir += "/" + dest
					curDir = strings.ReplaceAll(curDir, "//", "/")
				}
			}
		} else { // if not command then list of content
			var type_ rune
			var size int
			var name string

			tokens := strings.Split(line, " ")

			if tokens[0] == "dir" { // type directory
				type_ = 'd'
				name = tokens[1]
			} else { // type file
				type_ = 'f'
				size, _ = strconv.Atoi(tokens[0])
				name = tokens[1]
			}

			// append node to tree (and set parent)
			node := node{name, size, type_}
			nodes = append(nodes, node)
			parent = append(parent, curDir)
		}
	}
	return nodes, parent
}

// returns index of last / in dir, used for cd ..
func findLastSlash(dir string) int {
	var index int
	for i, v := range dir {
		if v == '/' {
			index = i
		}
	}
	return index
}

// calculate the size of each directory
func calcDirsSizes(nodes []node, parents []string) map[string]int {
	sizes := make(map[string]int)

	// scan each node
	for i, node := range nodes {

		// if node is a directory
		if node.type_ == 'd' {

			fulldir := strings.ReplaceAll(parents[i]+"/"+node.name, "//", "/")

			sizes = getDirSize(nodes, parents, fulldir, sizes)

		}
	}
	return sizes
}

// returns updated map of sizes with requested directory ("dir")
func getDirSize(nodes []node, parents []string, dir string, sizes map[string]int) map[string]int {
	var sum int

	// scan each node (files + dirs)
	for i, par := range parents {

		// if file is in current dir (has parent = current dir)
		if par == dir {
			if nodes[i].type_ == 'f' { // file: just sum size
				sum += nodes[i].size
			} else { // subdirectory: calculate size of subdirectory

				// directory already calculated
				subSize, found := sizes[dir]
				if found {
					sum += subSize
				}

				// directory not calculated yet
				if !found {

					// calculate dir of subdirectory
					subDir := strings.ReplaceAll(par+"/"+nodes[i].name, "//", "/")

					// calculate size of subdirectory and add to sizes
					sizes = getDirSize(nodes, parents, subDir, sizes)

					// sum subdirectory size to current directory total size
					sum += sizes[subDir]
				}
			}
		}

	}

	// if directory size not calculated yet, add to sizes
	if _, found := sizes[dir]; !found {
		sizes[dir] = sum
	}
	return sizes
}

func calculatePart1(dirSizes map[string]int) int {
	// sum directories smaller than 100000 (part1)
	var sum int
	for _, v := range dirSizes {
		if v > 100000 {
			continue
		}
		sum += v
	}
	return sum
}

func calculatePart2(dirSizes map[string]int) int {
	const TOTAL_SPACE int = 70000000  // total filespace size
	const NEEDED_SPACE int = 30000000 // space needed to update

	// space on filesystem
	totalFree := TOTAL_SPACE - dirSizes["/"]
	missing := NEEDED_SPACE - totalFree

	// smaller dir to delete to free enough space
	var min int = dirSizes["/"]
	for _, v := range dirSizes {
		if v > missing { // if dir bug enough to reach NEEDED_SPACE
			if v < min { // if the smallest yet
				min = v
			}
		}
	}
	return min
}
