// https://adventofcode.com/2022/day/15

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const SENSOR int = 1
const BEACON int = 2
const VOID int = 0
const FORCEDVOID int = -1

// const LINE int = 10

const LINE int = 2000000

func main() {
	sensors := parseInput()
	cave := placeSensorsBeacon(sensors)

	cave = forcedVoidOnLine(sensors, cave)
	forcedVoid := countForcedVoid(cave)

	fmt.Println(forcedVoid)
}

type Point struct {
	x, y int
}

func parseInput() map[Point]Point {
	cave := make(map[Point]Point)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		sx, sy := tokens[2], tokens[3]
		bx, by := tokens[8], tokens[9]

		sxx, _ := strconv.Atoi(sx[2 : len(sx)-1])
		syy, _ := strconv.Atoi(sy[2 : len(sy)-1])

		bxx, _ := strconv.Atoi(bx[2 : len(bx)-1])
		byy, _ := strconv.Atoi(by[2:])

		cave[Point{sxx, syy}] = Point{bxx, byy}
	}
	return cave
}

func manhattanDistance(a, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func placeSensorsBeacon(sensors map[Point]Point) map[Point]int {
	cave := make(map[Point]int)
	for s, b := range sensors {
		cave[s] = SENSOR
		cave[b] = BEACON
	}
	return cave
}

func forcedVoidOnLine(sensors map[Point]Point, cave map[Point]int) map[Point]int {
	minS := mostLeftX(sensors)
	min := minS.x - manhattanDistance(minS, sensors[minS])

	maxS := mostRightX(sensors)
	max := maxS.x + manhattanDistance(maxS, sensors[maxS])

	fmt.Println(min, max)

	// scan each point of line
	for i := min; i < max; i++ {
		thispoint := Point{i, LINE}

		if cave[thispoint] == SENSOR {
			continue
		}
		if cave[thispoint] == BEACON {
			continue
		}

		// scan each sensor
		for sen, bea := range sensors {

			// distance from line point to sensor
			thisdist := manhattanDistance(sen, thispoint)

			// distance from sensor to his beacon
			sendist := manhattanDistance(sen, bea)

			if thisdist <= sendist {
				cave[thispoint] = FORCEDVOID
			}
		}
	}
	return cave
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -(a)
}

func countLine(cave map[Point]int) (count int) {
	for p, v := range cave {
		if p.y == LINE {
			if v == FORCEDVOID {
				count++
			}
		}
	}
	return count
}

func mostLeftX(sensors map[Point]Point) Point {
	min := math.MaxInt
	var senMin Point
	for s := range sensors {
		if s.x < min {
			min = s.x
			senMin = s
		}
	}
	return senMin
}

func mostRightX(sensors map[Point]Point) Point {
	max := math.MinInt
	var senMax Point
	for s := range sensors {
		if s.x > max {
			max = s.x
			senMax = s
		}
	}
	return senMax
}

func countForcedVoid(cave map[Point]int) (count int) {
	for _, v := range cave {
		if v == FORCEDVOID {
			count++
		}
	}
	return count
}
