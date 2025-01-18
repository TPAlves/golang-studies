package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("TESTE"))
	c2 := sha256.Sum256([]byte("teste"))
	count := CountDiffBits(c1, c2)
	fmt.Print(count)
}

func CountDiffBits(c1, c2 [32]byte) (count int) {
	for i := range c1 {
		// j = j &  (j - 1) == j &= (j - 1)
		for j := c1[i] ^ c2[i]; j != 0; j &= (j - 1) {
			count++
		}
	}
	return count
}
