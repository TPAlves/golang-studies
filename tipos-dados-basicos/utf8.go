package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word := "Opa &$"
	for i, r := range word {
		// \t é voltado a tabulação
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println("\nValor total: ", utf8.RuneCountInString(word))
	fmt.Println(string(1234567)) // CARACTERE DE SUBSTITUIÇÃO
	fmt.Println(string(65)) // LETRA `A`
}
