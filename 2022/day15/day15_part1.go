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

const VOID int = -1
const SENSOR int = 1
const BEACON int = 2

const LINE int = 2000000

type Point struct {
	x, y int
}

func main() {
	sensors := parseInput()
	cave := placeSensorsAndBeacons(sensors)

	totVoid := placeVoidOnLine(sensors, cave)

	fmt.Print("Total number of void on line ", LINE, " (part1):\n\t", totVoid, "\n")
}

// returns each sensor associated to his beacon
func parseInput() map[Point]Point {
	sensors := make(map[Point]Point)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		// get x,y as strings
		senXstr, senYstr := tokens[2], tokens[3]
		beaXstr, beaYstr := tokens[8], tokens[9]

		// parse sensor x,y to int
		senX, _ := strconv.Atoi(senXstr[2 : len(senXstr)-1])
		senY, _ := strconv.Atoi(senYstr[2 : len(senYstr)-1])

		// parse beacon x,y to int
		beaX, _ := strconv.Atoi(beaXstr[2 : len(beaXstr)-1])
		beaY, _ := strconv.Atoi(beaYstr[2:])

		// add sensor
		sensors[Point{senX, senY}] = Point{beaX, beaY}
	}
	return sensors
}

// returns a cave with sensors and beacons placed on it
func placeSensorsAndBeacons(sensors map[Point]Point) map[Point]int {
	cave := make(map[Point]int)
	for s, b := range sensors {
		cave[s] = SENSOR
		cave[b] = BEACON
	}
	return cave
}

// returns the number of void points on line "LINE"
func placeVoidOnLine(sensors map[Point]Point, cave map[Point]int) (count int) {
	// sensors with the minimum and maximum x
	minS := mostLeftX(sensors)
	maxS := mostRightX(sensors)

	// minimum and maximum point on line: min x - distance from his beacon (max x + distance from his beacon)
	min := minS.x - manhattanDistance(minS, sensors[minS])
	max := maxS.x + manhattanDistance(maxS, sensors[maxS])

	// scan each point of line
	for x := min; x < max; x++ {
		linePoint := Point{x, LINE}

		// skip sensors and beacons
		if cave[linePoint] == SENSOR || cave[linePoint] == BEACON {
			continue
		}

		// scan each sensor to check if covers the current line point
		for sen, bea := range sensors {

			// distance from line point to sensor
			thisdist := manhattanDistance(sen, linePoint)

			// distance from sensor to his beacon
			sendist := manhattanDistance(sen, bea)

			// if sensor to beacon distance is greather than line point to beacon then the current line point is in sensor range
			if thisdist <= sendist {
				count++
				break
			}
		}
	}
	return count
}

// return manhattan distance between "a" and "b"
func manhattanDistance(a, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

// returns the absolute value of "a"
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -(a)
}

// returns the point with the minimum x
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

// returns the point with the maximum x
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
