package main

import (
	. "fmt"
	"strconv"
	"strings"
)

func main() {
	pesciStr := ""
	Scan(&pesciStr)
	pesci := strings.Split(pesciStr, ",")

	var pesciInt = []int{}
	for _, i := range pesci {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		pesciInt = append(pesciInt, j)
	}
	
	giorni := 0
	giorno := 1
	Scan(&giorni)

	for(giorno <= giorni) {

		for i:=0; i<len(pesciInt); i++ {
			if(pesciInt[i] == 0) {
				pesciInt[i] = 6
				pesciInt = append(pesciInt, 9)
				continue
			}
			pesciInt[i]--
		}
		// Println("Giorno:", giorno, "- pesci", len(pesciInt), "(", pesciInt, ")")
		Println("Giorno:", giorno, "- pesci", len(pesciInt))

		giorno++

	}
}
