package main_test

import (
	"testing"

	popcount "capitulo02/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(1000000000000000000)
	}
}
