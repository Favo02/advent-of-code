// https://adventofcode.com/2022/day/15

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX int = 4000000

type Point struct {
	x, y int
}

func main() {
	sensors := parseInput()
	borderPoints := scanSensorBorder(sensors)
	point := checkBorderPoints(sensors, borderPoints)

	fmt.Println("score of the point not in reach of any sensor (part2):\n\t", point.x*MAX+point.y)
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

// returns the points just outside of each sensor reach
func scanSensorBorder(sensors map[Point]Point) []Point {
	var borderPoints []Point
	for s, v := range sensors {
		dist := manhattanDistance(s, v)

		// add points just out of each direction of the sensor reach ("draw" a diamond)
		for i := dist; i >= 0; i-- {
			borderPoints = append(borderPoints, Point{s.x + (dist - i), s.y + (dist - (dist - i)) + 1})
			borderPoints = append(borderPoints, Point{s.x - (dist - i), s.y - (dist - (dist - i)) - 1})
			borderPoints = append(borderPoints, Point{s.x + (dist - (dist - i)) + 1, s.y + (dist - i)})
			borderPoints = append(borderPoints, Point{s.x - (dist - (dist - i)) - 1, s.y - (dist - i)})
		}
	}
	return borderPoints
}

// returns the manhattand distance between "a" and "b"
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

// returns the (first) point outside of each sensor reach in border points
func checkBorderPoints(sensors map[Point]Point, borderPoints []Point) Point {
	// scan each border point
	for _, borderPoint := range borderPoints {

		// if current border point outside of valid range skip
		if borderPoint.x > MAX || borderPoint.x < 0 {
			continue
		}
		if borderPoint.y > MAX || borderPoint.y < 0 {
			continue
		}

		inReach := false // in reach of a sensor

		// scan each sensor
		for sen, bea := range sensors {

			// distance from border point to sensor
			thisdist := manhattanDistance(sen, borderPoint)

			// distance from sensor to his beacon
			sendist := manhattanDistance(sen, bea)

			// if sensor to beacon distance is greather than border point to beacon then the current border point is in sensor range
			if thisdist <= sendist {
				inReach = true
				break
			}
		}

		// point not in reach of any sensor
		if !inReach {
			return borderPoint
		}
	}

	// no points outside of every sensor range
	return Point{}
}
