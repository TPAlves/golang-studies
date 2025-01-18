package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("Salveeeeee1"))
	c2 := sha256.Sum256([]byte("Salveeeeee"))

	// %x -> exibe o elemento como hexadecimal
	// %t -> exibe um booleano
	// %T -> exibe o tipo de dado
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
