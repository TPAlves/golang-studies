package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

	con "capitulo02/tempconv"
)

var n = flag.Bool("n", false, "Omite a quebra de linha")
var sep = flag.String("s", " ", "separador")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	var kelvin con.Kelvin = 10.0 
	fmt.Println(math.Round(float64(con.KToF(kelvin)))) 
}
