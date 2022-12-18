// https://adventofcode.com/2022/day/18
// https://github.com/Favo02/advent-of-code

// ------------------
// TODO: decent refactor, now im too tired
// ------------------

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	x, y, z int
	adj     int
}

func main() {
	lavaCubes := parseInput()  // get cubes
	lavaCubes = adj(lavaCubes) // calculate for every cube number of "faces" adjacent to any other cube
	area := calArea(lavaCubes) // calculate external area (no adjacent faces)

	// calculate bounds of positions
	minX, maxX := minmaxX(lavaCubes)
	minY, maxY := minmaxY(lavaCubes)
	minZ, maxZ := minmaxZ(lavaCubes)

	// generate every cube in bounds
	allCubes := generateCubes(Cube{minX, minY, minZ, 0}, Cube{maxX, maxY, maxZ, 0})
	// remove cube occupied by lava
	allCubesEmpty := getEmpty(allCubes, lavaCubes)
	// get cubes (out of every empty cube) trapped inside lava
	trapped := getTrapped(allCubesEmpty, lavaCubes, minX, maxX, minY, maxY, minZ, maxZ)

	// calculate for every trapped cube number of "faces" adjacent to any other trapped cube
	trappedAdj := adj(trapped)
	areaT := calArea(trappedAdj) // calculate external area (no adjacent faces) for trapped spaces

	fmt.Println("area:", area)
	fmt.Println("trapped:", len(trapped))
	fmt.Println("area not trapped:", area-areaT)
}

// returns maximum and minimin x coordinate
func minmaxX(cubes []Cube) (int, int) {
	min := math.MaxInt
	max := math.MinInt
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
	min := math.MaxInt
	max := math.MinInt
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
	min := math.MaxInt
	max := math.MinInt
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

// returns cubes parsed from stdin
// modifies stdin
func parseInput() (cubes []Cube) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		z, _ := strconv.Atoi(tokens[2])
		cube := Cube{x, y, z, 0}
		cubes = append(cubes, cube)
	}
	return cubes
}

// returns the cubes updating them with how many faces of the cube are adjacent to any other cube
func adj(cubes []Cube) []Cube {
	for i, c := range cubes {

		var adjPoints []Cube
		adjPoints = append(adjPoints, Cube{c.x - 1, c.y, c.z, 0})
		adjPoints = append(adjPoints, Cube{c.x + 1, c.y, c.z, 0})
		adjPoints = append(adjPoints, Cube{c.x, c.y + 1, c.z, 0})
		adjPoints = append(adjPoints, Cube{c.x, c.y - 1, c.z, 0})
		adjPoints = append(adjPoints, Cube{c.x, c.y, c.z + 1, 0})
		adjPoints = append(adjPoints, Cube{c.x, c.y, c.z - 1, 0})

		for _, ap := range adjPoints {

			for pi, p := range cubes {
				if pi == i {
					continue
				}
				if ap.x == p.x && ap.y == p.y && ap.z == p.z {
					cubes[i].adj++
				}
			}
		}

	}
	return cubes
}

// returns absolute value of "n"
func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

// returns the total external area of cubes (external = excluding adjacent faces)
func calArea(cubes []Cube) (res int) {
	for _, c := range cubes {
		res += (6 - c.adj)
	}
	return res
}

// returns every possible cube in range min, max
func generateCubes(min, max Cube) (gen []Cube) {
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
func getEmpty(cubes, filled []Cube) (res []Cube) {
	for _, c := range cubes {
		found := false
		for _, f := range filled {
			if c.x == f.x && c.y == f.y && c.z == f.z {
				found = true
			}
		}
		if !found {
			res = append(res, c)
		}
	}
	return res
}

// returns the cubes of "cubes" trapped inside some cubes of "traps"
func getTrapped(cubes, traps []Cube, minX, maxX, minY, maxY, minZ, maxZ int) (res []Cube) {
	for _, c := range cubes {
		if checkTrapped(c, traps, make(map[Cube]bool), minX, maxX, minY, maxY, minZ, maxZ) {
			res = append(res, c)
		}
	}
	return res
}

// returns true if cube "c" is trapped inside "traps" cubes, false otherwise
func checkTrapped(c Cube, traps []Cube, crossed map[Cube]bool, minX, maxX, minY, maxY, minZ, maxZ int) bool {
	if c.x > maxX || c.x < minX {
		return false
	}
	if c.y > maxY || c.y < minY {
		return false
	}
	if c.z > maxZ || c.z < minZ {
		return false
	}

	crossed[c] = true

	var next Cube

	next = Cube{c.x + 1, c.y, c.z, 0}
	if !contains(traps, next) && !crossed[next] {
		if !checkTrapped(next, traps, crossed, minX, maxX, minY, maxY, minZ, maxZ) {
			return false
		}
	}
	next = Cube{c.x - 1, c.y, c.z, 0}
	if !contains(traps, next) && !crossed[next] {
		if !checkTrapped(next, traps, crossed, minX, maxX, minY, maxY, minZ, maxZ) {
			return false
		}
	}

	next = Cube{c.x, c.y + 1, c.z, 0}
	if !contains(traps, next) && !crossed[next] {
		if !checkTrapped(next, traps, crossed, minX, maxX, minY, maxY, minZ, maxZ) {
			return false
		}
	}
	next = Cube{c.x, c.y - 1, c.z, 0}
	if !contains(traps, next) && !crossed[next] {
		if !checkTrapped(next, traps, crossed, minX, maxX, minY, maxY, minZ, maxZ) {
			return false
		}
	}

	next = Cube{c.x, c.y, c.z + 1, 0}
	if !contains(traps, next) && !crossed[next] {
		if !checkTrapped(next, traps, crossed, minX, maxX, minY, maxY, minZ, maxZ) {
			return false
		}
	}
	next = Cube{c.x, c.y, c.z - 1, 0}
	if !contains(traps, next) && !crossed[next] {
		if !checkTrapped(next, traps, crossed, minX, maxX, minY, maxY, minZ, maxZ) {
			return false
		}
	}

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
