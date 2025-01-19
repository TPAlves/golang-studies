package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var input string
var sha int
var shaValues = [3]int{256, 384, 512}

const messageDefault = "Tipo do SHA: %d\nValor em hexadecimal: %x\n"

func main() {
	flag.StringVar(&input, "t", "", "Texto a ser convertido para algum SHA")
	flag.IntVar(&sha, "s", shaValues[0], "Valor do SHA a ser utilizado na conversão")
	flag.Parse()
	if input == "" {
		fmt.Printf("O texto é obrigatório, declare utilizando a flag t\n")
		os.Exit(1)
	}
	switch sha {
	case shaValues[0]:
		fmt.Printf(messageDefault, shaValues[0], sha256.Sum256([]byte(input)))
	case shaValues[1]:
		fmt.Printf(messageDefault, shaValues[1], sha512.Sum384([]byte(input)))
	case shaValues[2]:
		fmt.Printf(messageDefault, shaValues[2], sha512.Sum512([]byte(input)))
	default:
		fmt.Printf("SHA inválido declarado. São aceitos os seguintes valores de SHA: %+v\n", shaValues)
		os.Exit(1)
	}
}
