package main

import (
	"fmt"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// read starting fishes
	fishesStr := ""
	Scan(&fishesStr)
	fishes := strings.Split(fishesStr, ",")

	// parse strings slice to ints slice
	var fishesInt = []int{}
	for _, i := range fishes {
		j, _ := strconv.Atoi(i)
		fishesInt = append(fishesInt, j)
	}

	// current day
	var curDay int = 1

	// number of days
	var days int
	if len(os.Args) < 2 {
		fmt.Println("insert number of days to calculate as command line argument")
		return
	}
	daysStr := os.Args[1]
	days, _ = strconv.Atoi(daysStr)

	// calculate starting statuses of starting fishes
	fishesNumArray := printUniqueValue(fishesInt)
	fishesNum := fishesNumArray[:] //array to slice

	// for each day
	for curDay <= days {
		// fishes at status 0 are the only one that do something
		zeroFishes := fishesNum[0]

		// shift left fish slice (fish in status 8 --> status 7, status 7 --> 6, ...)
		// move fishes in status 0 to status 8 (new fishes generated)
		fishesNum = append(fishesNum[1:], zeroFishes)
		// move parent fishes (that just generated a fish) to status 6
		fishesNum[6] += zeroFishes

		curDay++
	}
	Println(sumSlice(fishesNum))

}

// returns number of fishes in each status
func printUniqueValue(arr []int) [9]int {
	var fishesNum [9]int
	for _, v := range arr {
		fishesNum[v]++
	}
	return fishesNum
}

// returns the sum of slice
func sumSlice(sl []int) int {
	res := 0
	for _, v := range sl {
		res += v
	}
	return res
}
