package main

import (
	"fmt"
	"strings"
)

// basenmae remove componentes de diretórios e um .sufixo.
// Sem uso de libs
func basename(s string) string {
	// DeSCARTA A ULTIMA '/' E TUDO QUE ESTIVER ANTES
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func newBasename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 se "/" não for encontrada
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	name := basename("/receitas.txt")
	nameTwo := newBasename("/receitas.txt")
	fmt.Println(name)
	fmt.Println(nameTwo)
}
