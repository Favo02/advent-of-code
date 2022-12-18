// https://adventofcode.com/2022/day/18
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cube struct: coordinates position and number of faces adjacent to another cube
type Cube struct {
	x, y, z int
	adj     int
}

func main() {
	lavaCubes := parseInput()                // get cubes
	lavaCubes = calculateAdjacent(lavaCubes) // calculate for every cube number of faces adjacent to any other cube
	part1Area := calculateArea(lavaCubes)    // calculate external area (no adjacent faces)

	// calculate bounds of positions
	minX, maxX := minmaxX(lavaCubes)
	minY, maxY := minmaxY(lavaCubes)
	minZ, maxZ := minmaxZ(lavaCubes)
	minCube := Cube{minX, minY, minZ, 0}
	maxCube := Cube{maxX, maxY, maxZ, 0}

	allCubes := generateCubes(minCube, maxCube)                       // generate every possible cube inside are delimited by "minCube" and "maxCube"
	allEmptyCubes := removeFilled(allCubes, lavaCubes)                // remove cube occupied by lava
	trapped := getTrapped(allEmptyCubes, lavaCubes, minCube, maxCube) // get only empty cubes trapped inside lava cubes

	trapped = calculateAdjacent(trapped)  // calculate for every trapped cube number of faces adjacent to any other trapped cube
	trappedArea := calculateArea(trapped) // calculate external area (no adjacent faces) for trapped spaces
	part2Area := part1Area - trappedArea  // remove from complete area the area of trapped cubes

	fmt.Println("lava cubes area, including internal air bubbles (part1):\n\t", part1Area)
	fmt.Println("lava cubes area, excluding internal air bubbles (part2):\n\t", part2Area)
}

// returns cubes parsed from stdin
// modifies stdin
func parseInput() (cubes []Cube) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",") // split coordinates

		// parse coordinates from string to int
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		z, _ := strconv.Atoi(tokens[2])

		// append cube
		cubes = append(cubes, Cube{x, y, z, 0})
	}
	return cubes
}

// returns the cubes, updating each one with how many faces of the cube are adjacent to any other cube
func calculateAdjacent(cubes []Cube) []Cube {
	for i, c := range cubes { // scan each cube

		// cube adjacent on every side
		var adjCubes []Cube
		adjCubes = append(adjCubes, Cube{c.x - 1, c.y, c.z, 0})
		adjCubes = append(adjCubes, Cube{c.x + 1, c.y, c.z, 0})
		adjCubes = append(adjCubes, Cube{c.x, c.y + 1, c.z, 0})
		adjCubes = append(adjCubes, Cube{c.x, c.y - 1, c.z, 0})
		adjCubes = append(adjCubes, Cube{c.x, c.y, c.z + 1, 0})
		adjCubes = append(adjCubes, Cube{c.x, c.y, c.z - 1, 0})

		// check if the adjacent cubes exists
		for _, adj := range adjCubes { // scan each adjacent cube
			for cubeIndex, cube := range cubes { // scan each cube
				if cubeIndex == i { // skip self
					continue
				}
				if adj.x == cube.x && adj.y == cube.y && adj.z == cube.z { // same cube: adjacent exists
					cubes[i].adj++ // update adjacent number for current cube
				}
			}
		}
	}
	return cubes
}

// returns the total external area of cubes (external: excluding adjacent faces)
func calculateArea(cubes []Cube) (res int) {
	for _, c := range cubes { // scan each cube
		res += (6 - c.adj) // add to area number of external faces
	}
	return res
}

// returns maximum and minimin x coordinate
func minmaxX(cubes []Cube) (int, int) {
	min := cubes[0].x
	max := cubes[0].x
	for _, v := range cubes {
		if v.x < min {
			min = v.x
		}
		if v.x > max {
			max = v.x
		}
	}
	return min, max
}

// returns maximum and minimin y coordinate
func minmaxY(cubes []Cube) (int, int) {
	min := cubes[0].y
	max := cubes[0].y
	for _, v := range cubes {
		if v.y < min {
			min = v.y
		}
		if v.y > max {
			max = v.y
		}
	}
	return min, max
}

// returns maximum and minimin z coordinate
func minmaxZ(cubes []Cube) (int, int) {
	min := cubes[0].z
	max := cubes[0].z
	for _, v := range cubes {
		if v.z < min {
			min = v.z
		}
		if v.z > max {
			max = v.z
		}
	}
	return min, max
}

// returns every possible cube in range min, max
func generateCubes(min, max Cube) (gen []Cube) {
	// create cube for every possible coordinate
	for y := min.y; y < max.y; y++ {
		for x := min.x; x < max.x; x++ {
			for z := min.z; z < max.z; z++ {
				gen = append(gen, Cube{x, y, z, 0})
			}
		}
	}
	return gen
}

// returns only the cubes of "cubes" not included in "filled"
func removeFilled(cubes, filled []Cube) (empty []Cube) {
	for _, c := range cubes { // scan each cube
		found := false             // found a filled cube on the same coordinates of current cube
		for _, f := range filled { // scan each filled cube
			if c.x == f.x && c.y == f.y && c.z == f.z { // i
				found = true
			}
		}
		if !found { // no filled cubes on current coords
			empty = append(empty, c) // this is an empty cube
		}
	}
	return empty
}

// returns the cubes of "cubes" trapped inside some cubes of "lavaCubes"
func getTrapped(cubes, lavaCubes []Cube, minCube, maxCube Cube) (trapped []Cube) {
	for _, c := range cubes { // scan each cube
		if checkTrapped(c, lavaCubes, make(map[Cube]bool), minCube, maxCube) { // cube is trapped
			trapped = append(trapped, c)
		}
	}
	return trapped
}

// returns true if cube "c" is trapped inside "lavaCubes" cubes, false otherwise
func checkTrapped(current Cube, lavaCubes []Cube, crossed map[Cube]bool, minCube, maxCube Cube) bool {

	// cube is out of bounds, no lavaCubes out of bounds: not trapped
	if current.x > maxCube.x || current.x < minCube.x {
		return false
	}
	if current.y > maxCube.y || current.y < minCube.y {
		return false
	}
	if current.z > maxCube.z || current.z < minCube.z {
		return false
	}

	// mark current cube as crossed
	crossed[current] = true

	// next cubes to expand and check (adjacent to current)
	next := []Cube{{current.x + 1, current.y, current.z, 0}, {current.x - 1, current.y, current.z, 0}, {current.x, current.y + 1, current.z, 0}, {current.x, current.y - 1, current.z, 0}, {current.x, current.y, current.z + 1, 0}, {current.x, current.y, current.z - 1, 0}}

	for i := 0; i < 6; i++ { // scan all next cubes
		if !crossed[next[i]] && !contains(lavaCubes, next[i]) { // cube not already checked and cube is not a lava cube
			if !checkTrapped(next[i], lavaCubes, crossed, minCube, maxCube) { // if recursive search reaches bounds: not trapped
				return false
			}
		}
	}

	// scanned recursively all next cubes and no cube leads to border: trapped
	return true
}

// returns true if "cube" is contained in "cubes"
func contains(cubes []Cube, cube Cube) bool {
	for _, c := range cubes {
		if c.x == cube.x && c.y == cube.y && c.z == cube.z {
			return true
		}
	}
	return false
}
