// https://adventofcode.com/2022/day/18
// https://github.com/Favo02/advent-of-code

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
	cubes := parseInput()
	cubes = adj(cubes)
	area := calArea(cubes)

	// calculate bounds of positions
	minX, maxX := minmaxX(cubes)
	minY, maxY := minmaxY(cubes)
	minZ, maxZ := minmaxZ(cubes)

	allCubes := generateCubes(Cube{minX, minY, minZ, 0}, Cube{maxX, maxY, maxZ, 0})
	allCubesEmpty := getEmpty(allCubes, cubes)
	trapped := getTrapped(allCubesEmpty, cubes, minX, maxX, minY, maxY, minZ, maxZ)

	trappedAdj := adj(trapped)
	areaT := calArea(trappedAdj)

	fmt.Println("area:", area)
	fmt.Println("trapped:", len(trapped))
	fmt.Println("area not trapped:", area-areaT)
}

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

func adj2(cubes, cfr []Cube) []Cube {
	for i, c := range cubes {

		var adjPoints []Cube
		adjPoints = append(adjPoints, Cube{c.x - 1, c.y, c.z, 0})
		adjPoints = append(adjPoints, Cube{c.x + 1, c.y, c.z, 0})
		adjPoints = append(adjPoints, Cube{c.x, c.y + 1, c.z, 0})
		adjPoints = append(adjPoints, Cube{c.x, c.y - 1, c.z, 0})
		adjPoints = append(adjPoints, Cube{c.x, c.y, c.z + 1, 0})
		adjPoints = append(adjPoints, Cube{c.x, c.y, c.z - 1, 0})

		for _, ap := range adjPoints {

			for _, p := range cfr {
				if ap.x == p.x && ap.y == p.y && ap.z == p.z {
					cubes[i].adj++
				}
			}
		}

	}
	return cubes
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func calArea(cubes []Cube) (res int) {
	for _, c := range cubes {
		res += (6 - c.adj)
	}
	return res
}

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

func get6adj(cubes []Cube) (res []Cube) {
	for _, c := range cubes {
		if c.adj == 6 {
			res = append(res, c)
		}
	}
	return res
}

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

func getTrapped(cubes, traps []Cube, minX, maxX, minY, maxY, minZ, maxZ int) (res []Cube) {
	for _, c := range cubes {
		if checkTrapped(c, traps, make(map[Cube]bool), minX, maxX, minY, maxY, minZ, maxZ) {
			res = append(res, c)
		}
	}
	return res
}

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

func contains(cubes []Cube, cube Cube) bool {
	for _, c := range cubes {
		if c.x == cube.x && c.y == cube.y && c.z == cube.z {
			return true
		}
	}
	return false
}
