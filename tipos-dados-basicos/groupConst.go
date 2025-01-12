package main

import (
	"fmt"
)

type Semana int
type Flags uint

const (
	Domingo Semana = iota // O iota icrementa os demais sucessivamente
	Segunda
	Terca
	Quarta
	Quinta
	Sexta
)

const (
	FlagUp           Flags = 1 << iota // Ativa  1 * 2^0
	FlagBroadcast                      // Suporta recurso de acesso broadcast 1 * 2^1
	FlagLoopBack                       // é uma interface loopback 1 * 2^2
	FlagPointToPoint                   // pertence a um link ponto a ponto 1 * 2^3
	FlagMulticast                      // suporta recurso de acesso multicast 1 * 2^4
)

const (
	// untyped constante sem tipo definido 
	_ = 1 << (10 * iota)
	Kilobyte
	Megabyte
	Gigabyte
	Terabyte
	Petabyte
	Exabyte
	Zetabyte
	Yottabyte
)

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
