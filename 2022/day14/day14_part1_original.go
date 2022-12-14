// https://adventofcode.com/2022/day/14

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

var sandStart Point = Point{500, 0}

func main() {
	rocks := parseInput()
	cave := fillGrid(rocks)
	// fmt.Println(cave)
	// fmt.Println(len(cave))
	cave = simulateSand(cave)

	sand := countSand(cave)
	fmt.Println(sand)

	// fmt.Println(cave)
}

func parseInput() (rocks [][]Point) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " -> ")
		var rock []Point
		for _, t := range tokens {
			xy := strings.Split(t, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			rock = append(rock, Point{x, y})
		}
		rocks = append(rocks, rock)
	}
	return rocks
}

func fillGrid(rocks [][]Point) map[Point]int {
	cave := make(map[Point]int)
	for _, rock := range rocks {
		for i := 0; i < len(rock)-1; i++ {
			drawLine(cave, rock[i], rock[i+1])
		}
	}
	return cave
}

func drawLine(cave map[Point]int, s, d Point) map[Point]int {
	diffX := s.x - d.x
	diffY := s.y - d.y
	if diffX > 0 {
		for i := 0; i <= diffX; i++ {
			cave[Point{s.x - i, s.y}] = 1
		}
	} else {
		for i := 0; i <= -diffX; i++ {
			cave[Point{s.x + i, s.y}] = 1
		}
	}

	if diffY > 0 {
		for i := 0; i <= diffY; i++ {
			cave[Point{s.x, s.y - i}] = 1
		}
	} else {
		for i := 0; i <= -diffY; i++ {
			cave[Point{s.x, s.y + i}] = 1
		}
	}
	return cave
}

func sandFall(cave map[Point]int, sand Point) map[Point]int {
	var rest bool
	min := lowerPoint(cave)
	for !rest {
		if sand.y > min {
			cave[sand] = 0
			break
		}
		// fmt.Println("cursand:", sand)
		// fmt.Println(countSand(cave))
		if cave[Point{sand.x, sand.y + 1}] == 0 {
			cave[Point{sand.x, sand.y + 1}] = -1
			cave[sand] = 0
			sand = Point{sand.x, sand.y + 1}
			continue
		}
		if cave[Point{sand.x - 1, sand.y + 1}] == 0 {
			cave[Point{sand.x - 1, sand.y + 1}] = -1
			cave[sand] = 0
			sand = Point{sand.x - 1, sand.y + 1}
			continue
		}
		if cave[Point{sand.x + 1, sand.y + 1}] == 0 {
			cave[Point{sand.x + 1, sand.y + 1}] = -1
			cave[sand] = 0
			sand = Point{sand.x + 1, sand.y + 1}
			continue
		}
		rest = true
	}
	// fmt.Println("endsand:", sand)
	return cave
}

func simulateSand(cave map[Point]int) map[Point]int {
	// var rest bool
	for true {
		oldMap := copyMap(cave)

		cave[sandStart] = -1
		cave = sandFall(cave, sandStart)

		if sameMap(oldMap, cave) {
			break
		}
	}
	return cave
}

func countSand(cave map[Point]int) int {
	var sand int
	for _, v := range cave {
		if v == -1 {
			sand++
		}
	}
	return sand
}

func countRocks(cave map[Point]int) int {
	var rocks int
	for _, v := range cave {
		if v == -1 {
			rocks++
		}
	}
	return rocks
}

func lowerPoint(cave map[Point]int) int {
	low := math.MinInt
	for k := range cave {
		if k.y > low {
			low = k.y
		}
	}
	return low
}

func copyMap(cave map[Point]int) map[Point]int {
	newMap := make(map[Point]int)
	for k, v := range cave {
		newMap[k] = v
	}
	return newMap
}

func sameMap(m1, m2 map[Point]int) bool {
	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	for k := range m2 {
		if m1[k] != m2[k] {
			return false
		}
	}
	return true
}
