// zera o conteúdo de um array [32] byte
package main

import "fmt"

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}

func main() {
	array := [32]byte{1, 2, 3, 4}
	fmt.Printf("%v\n", array)
	zero(&array) // endereço da variável array
	fmt.Printf("%v\n", array)
}
