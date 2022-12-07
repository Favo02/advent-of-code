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
	x int
	y int
}

func main() {
	// values1: only horizontal, vertical lines (part1)
	// values2: horizontal, vertical, diagonal lines (part2)
	values1, values2 := parseInput()
	score1 := countMoreThan1(values1)
	score2 := countMoreThan1(values2)
	fmt.Println("only horizontal, vertical lines score (part1):", score1)
	fmt.Println("horizontal, vertical, diagonal lines score (part2):", score2)
}

// returns maps with points crossed
func parseInput() (map[Point]int, map[Point]int) {
	// map containing every point crossed by a line, with value how many times the point has been crossed
	values1 := make(map[Point]int)
	values2 := make(map[Point]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// split line into starting and destination point
		tokens := strings.Split(line, " -> ")

		// split each point into x and y
		tokens1 := strings.Split(tokens[0], ",")
		tokens2 := strings.Split(tokens[1], ",")

		// convert x and y (strings) to int
		x1, _ := strconv.Atoi(tokens1[0])
		y1, _ := strconv.Atoi(tokens1[1])
		x2, _ := strconv.Atoi(tokens2[0])
		y2, _ := strconv.Atoi(tokens2[1])

		// calculate points crossed for each line
		values1, values2 = calculateLine(values1, values2, x1, y1, x2, y2)
	}

	// returns maps of points crossed
	return values1, values2
}

// calculates which points a line crosses
func calculateLine(vals1, vals2 map[Point]int, x1, y1, x2, y2 int) (map[Point]int, map[Point]int) {

	// diagonal line
	if !(x1 == x2 || y1 == y2) {
		return vals1, calculateDiagonalLine(vals2, x1, y1, x2, y2)
	}

	// horizonal line
	diffX := int(math.Abs(float64(x1 - x2)))
	if diffX != 0 {
		// number of points crossed = difference
		for i := 0; i <= diffX; i++ {
			if x1 < x2 {
				vals1[Point{x1 + i, y1}]++
				vals2[Point{x1 + i, y1}]++
			} else {
				vals1[Point{x1 - i, y1}]++
				vals2[Point{x1 - i, y1}]++
			}
		}
	}

	// vertical line
	diffY := int(math.Abs(float64(y1 - y2)))
	if diffY != 0 {
		// number of points crossed = difference
		for i := 0; i <= diffY; i++ {
			if y1 < y2 {
				vals1[Point{x1, y1 + i}]++
				vals2[Point{x1, y1 + i}]++
			} else {
				vals1[Point{x1, y1 - i}]++
				vals2[Point{x1, y1 - i}]++
			}
		}
	}

	return vals1, vals2
}

// calculate the points crossed by a diagonal line
func calculateDiagonalLine(vals map[Point]int, x1, y1, x2, y2 int) map[Point]int {
	diff := int(math.Abs(float64(x1 - x2)))
	// number of points crossed = difference
	for i := 0; i <= diff; i++ {
		if x1 > x2 {
			if y1 > y2 {
				vals[Point{x1 - i, y1 - i}]++
			} else {
				vals[Point{x1 - i, y1 + i}]++
			}
		} else {
			if y1 > y2 {
				vals[Point{x1 + i, y1 - i}]++
			} else {
				vals[Point{x1 + i, y1 + i}]++
			}
		}
	}
	return vals
}

// returns the count of points crossed more than 1 time
func countMoreThan1(vals map[Point]int) int {
	var res int
	for _, v := range vals {
		if v >= 2 {
			res++
		}
	}
	return res
}
