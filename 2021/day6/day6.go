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
	pesciStr := ""
	Scan(&pesciStr)
	pesci := strings.Split(pesciStr, ",")

	// parse strings slice to ints slice
	var pesciInt = []int{}
	for _, i := range pesci {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		pesciInt = append(pesciInt, j)
	}

	// current day
	var giorno int = 1

	// number of days
	var giorni int
	if len(os.Args) < 2 {
		fmt.Println("insert number of days to calculate as command line argument")
		return
	}
	giorniStr := os.Args[1]
	giorni, _ = strconv.Atoi(giorniStr)

	// calculate starting statuses of starting fishes
	num_pesci_array := printUniqueValue(pesciInt)
	num_pesci := num_pesci_array[:] //array to slice

	// for each day
	for giorno <= giorni {
		// fishes at status 0 are the only one that do something
		num_pesci_a_zero := num_pesci[0]

		// shift left fish slice (fish in status 8 --> status 7, status 7 --> 6, ...)
		// move fishes in status 0 to status 8 (new fishes generated)
		num_pesci = append(num_pesci[1:], num_pesci_a_zero)
		// move parent fishes (that just generated a fish) to status 6
		num_pesci[6] += num_pesci_a_zero

		giorno++
	}
	Println(sommaSlice(num_pesci))

}

// returns number of fishes in each status
func printUniqueValue(arr []int) [9]int {
	var num_pesci [9]int
	for _, v := range arr {
		num_pesci[v]++
	}
	return num_pesci
}

// returns the sum of slice
func sommaSlice(sl []int) int {
	res := 0
	for _, v := range sl {
		res += v
	}
	return res
}
