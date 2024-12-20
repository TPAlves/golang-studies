// Echo1 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	// 1º Versão
	// var s, sep string
	// for i:= 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// 2º Versão
	s, sep := "", ""
	for indice, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(indice)
		fmt.Println(arg)
	}
	fmt.Println(os.Args[0])
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], " "))

	teste  := []string{"batata","teste3", "shazam", "aaaaaaaaa", "shazam3"}
	fmt.Println(teste[2:])
	fmt.Println(strings.Join(teste[1:], ""))
}