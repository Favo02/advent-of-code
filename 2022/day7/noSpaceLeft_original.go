// https://adventofcode.com/2022/day/7

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TOTAL_SPACE int = 70000000
const NEEDED_SPACE int = 30000000

func main() {
	nodes, parents := parseInput()
	calcDir := calcDir(nodes, parents)
	var sum int
	for _, v := range calcDir {
		if v > 100000 {
			continue
		}
		sum += v
	}
	fmt.Println(sum)

	totalFree := TOTAL_SPACE - calcDir["/"]
	missing := NEEDED_SPACE - totalFree

	var min int = calcDir["/"]
	for _, v := range calcDir {
		if v > missing {
			if v < min {
				min = v
			}
		}
	}
	fmt.Println(min)
}

type node struct {
	name  string
	size  int
	type_ rune
}

type size_ struct {
	name string
	size int
}

func parseInput() ([]node, []string) {
	var nodes []node
	var parent []string

	var curDir string

	nodes = append(nodes, node{"/", 0, 'd'})
	parent = append(parent, "")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:1] == "$" {
			command := line[2:4]
			if command == "cd" {
				dest := line[5:]
				if dest == ".." {
					curDir = curDir[:findLastSlash(curDir)]
				} else if dest == "/" {
					curDir = "/"
				} else {
					if curDir[len(curDir)-1] != '/' {

						curDir += "/" + dest
					} else {
						curDir += dest

					}
				}

			}
		} else {
			var type_ rune
			var size int
			var name string

			tokens := strings.Split(line, " ")

			if tokens[0] == "dir" {
				type_ = 'd'
				name = tokens[1]
			} else {
				type_ = 'f'
			}

			if type_ == 'f' {
				size, _ = strconv.Atoi(tokens[0])
				name = tokens[1]
			}

			node := node{name, size, type_}
			nodes = append(nodes, node)
			parent = append(parent, curDir)
		}
	}
	return nodes, parent
}

func findLastSlash(dir string) int {
	var index int
	var count int
	for i, v := range dir {
		if v == '/' {
			index = i
			count++
		}
	}
	if count == 1 {
		return index + 1
	}
	return index
}

func calcDir(nodes []node, parents []string) map[string]int {
	var sizes map[string]int = make(map[string]int)
	for i, node := range nodes {
		if node.type_ == 'd' {
			fulldir := parents[i] + "/" + node.name
			fulldir = strings.ReplaceAll(fulldir, "//", "/")
			dir, size := getDirSize(nodes, parents, fulldir, sizes)
			if dir == "/" {
				sizes[dir] = size
			}
		}
	}
	return sizes
}

func getDirSize(nodes []node, parents []string, dir string, sizes map[string]int) (string, int) {
	var sum int
	for i, v := range parents {
		// par := v[findLastSlash(v):]
		// if len(par) == 0 {
		// 	par = "/"
		// } else if par[0] == '/' {
		// 	par = par[1:]
		// }
		par := v

		if par == dir {
			if nodes[i].type_ == 'f' {
				sum += nodes[i].size
			} else {
				var skip bool
				if s, found := sizes[dir]; found {
					sum += s
					skip = true
				}

				if !skip {
					fulldir := v + "/" + nodes[i].name
					fulldir = strings.ReplaceAll(fulldir, "//", "/")
					d, s := getDirSize(nodes, parents, fulldir, sizes)
					sizes[d] = s
					sum += s
				}
			}
		}
	}
	return dir, sum
}
