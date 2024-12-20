package main

import "fmt"

// pc[i] é a população de i

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount devolve a população (número de bits definidos) de x
// @TODO: EXERCICIO 2.4
// @TODO: EXERCICIO 2.5
func PopCount(x uint64) (population int) {
	for i := range 8 {
		population += int(pc[byte(x>>(i*8))])

	}
	return
}

func main() {
	x := PopCount(1)
	fmt.Println(x)
}
