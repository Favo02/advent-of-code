// https://adventofcode.com/2022/day/15

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SENSOR int = 1
const BEACON int = 2
const VOID int = 0
const FORCEDVOID int = -1

const MAX int = 4000000

func main() {
	sensors := parseInput()
	cave := placeSensorsBeacon(sensors)

	borders := scanSensorBorder(sensors)

	for borderPoint, v := range borders {

		if !v {
			continue
		}

		if cave[borderPoint] == SENSOR {
			continue
		}
		if cave[borderPoint] == BEACON {
			continue
		}

		close := false
		// scan each sensor
		for sen, bea := range sensors {

			// distance from line point to sensor
			thisdist := manhattanDistance(sen, borderPoint)

			// distance from sensor to his beacon
			sendist := manhattanDistance(sen, bea)

			if thisdist <= sendist {
				close = true
				break
			}
		}

		if !close {
			if borderPoint.x > MAX || borderPoint.x < 0 {
				continue
			}
			if borderPoint.y > MAX || borderPoint.y < 0 {
				continue
			}
			fmt.Println(borderPoint)
		}

	}
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

func forcedVoidOnLine(sensors map[Point]Point, cave map[Point]int) (int, int) {
	min := 0
	max := MAX

	// scan each point of line
	for y := min; y < max; y++ {
		for x := 0; x < max; x++ {

			thispoint := Point{x, y}

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
					break
				}
			}
		}

		// check if line has void
		index := lineContaints(cave, y)
		if index >= 0 {
			return index, y
		}
	}
	return -1, -1
}

func scanSensorBorder(sensors map[Point]Point) map[Point]bool {
	border := make(map[Point]bool)
	for s, v := range sensors {
		dist := manhattanDistance(s, v)

		for i := dist; i >= 0; i-- {
			border[Point{s.x + (dist - i), s.y + (dist - (dist - i)) + 1}] = true
			border[Point{s.x - (dist - i), s.y - (dist - (dist - i)) - 1}] = true

			border[Point{s.x + (dist - (dist - i)) + 1, s.y + (dist - i)}] = true
			border[Point{s.x - (dist - (dist - i)) - 1, s.y - (dist - i)}] = true
		}
	}
	return border
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -(a)
}

func lineContaints(cave map[Point]int, line int) int {
	for i := 0; i < MAX; i++ {
		if cave[Point{i, line}] == VOID {
			return i
		}
	}
	return -1
}
