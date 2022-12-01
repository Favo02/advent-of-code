// pesci lanterna - advent of code 2021 day 6
// parti 1, 2

package main

import (
	. "fmt"
	"strconv"
	"strings"
)

func main() {

	// prendo in input i pesci iniziali come string
	pesciStr := ""
	Scan(&pesciStr)
	pesci := strings.Split(pesciStr, ",")

	// converto lo slice di stringhe a slice di interi
	var pesciInt = []int{}
	for _, i := range pesci {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		pesciInt = append(pesciInt, j)
	}

	// prendo in input il numero di giorni
	giorni := 0
	giorno := 1
	Scan(&giorni)

	// calcolo il numero di stati per i pesci iniziali
	// i pesci iniziali possono essere solo da 0 a 6 ma predispongo fino ad 8
	// per quando faranno spawnare nuovi pesciolini
	num_pesci_array := printUniqueValue(pesciInt)
	num_pesci := num_pesci_array[:] //conversione array to slice

	// inizio a scorrere i giorni
	for giorno <= giorni {
		// metto da parte i pesci nello stato 0 (gli unici che effettivamente fanno qualcosa)
		num_pesci_a_zero := num_pesci[0]

		// faccio scorrere a sinistra la slice (quelli nello stato 8 vanno a 7, i 6 a 5, ...)
		// facendo finire i pesci nello stato 0 nello stato 8 (nuovi pesciolini)
		num_pesci = append(num_pesci[1:], num_pesci_a_zero)
		// faccio tornare i pesci "genitori" che hanno appena fatto un altro pesce nello stato 6
		// sommandoli a quelli che sono evoluti nello stato 6 (dal 7)
		num_pesci[6] += num_pesci_a_zero

		Println("Giorno:", giorno, "- pesci", sommaSlice(num_pesci))

		giorno++
	}
	Println(sommaSlice(num_pesci))

}

// calcola il numero di pesci in ogni stato
func printUniqueValue(arr []int) [9]int {
	var num_pesci [9]int
	for _, v := range arr {
		num_pesci[v]++
	}
	Println(num_pesci)
	return num_pesci
}

// somma i numeri contenuti nella slice
func sommaSlice(sl []int) int {
	res := 0
	for _, v := range sl {
		res += v
	}
	return res
}
