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
	values1, values2 := parseInput()
	score1 := countMoreThan2(values1)
	score2 := countMoreThan2(values2)
	fmt.Println("only horizontal, vertical lines score (part1):", score1)
	fmt.Println("horizontal, vertical, diagonal lines score (part2):", score2)
}

func parseInput() (map[Point]int, map[Point]int) {
	values1 := make(map[Point]int)
	values2 := make(map[Point]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " -> ")
		tokens1 := strings.Split(tokens[0], ",")
		tokens2 := strings.Split(tokens[1], ",")
		x1, _ := strconv.Atoi(tokens1[0])
		y1, _ := strconv.Atoi(tokens1[1])
		x2, _ := strconv.Atoi(tokens2[0])
		y2, _ := strconv.Atoi(tokens2[1])
		values1, values2 = calculateLine(values1, values2, x1, y1, x2, y2)
	}

	return values1, values2
}

func calculateLine(vals1, vals2 map[Point]int, x1, y1, x2, y2 int) (map[Point]int, map[Point]int) {
	if !(x1 == x2 || y1 == y2) {
		return vals1, calculateDiagonalLine(vals2, x1, y1, x2, y2)
	}

	diffX := int(math.Abs(float64(x1 - x2)))
	if diffX != 0 {
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

	diffY := int(math.Abs(float64(y1 - y2)))
	if diffY != 0 {
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

func calculateDiagonalLine(vals map[Point]int, x1, y1, x2, y2 int) map[Point]int {
	diff := int(math.Abs(float64(x1 - x2)))
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

func countMoreThan2(vals map[Point]int) int {
	var res int
	for _, v := range vals {
		if v >= 2 {
			res++
		}
	}
	return res
}
