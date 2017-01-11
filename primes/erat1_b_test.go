package primes

import (
	"testing"
)

func BenchmarkErat1(b *testing.B) {
	var n int64 = 2000000
	for i := 0; i < b.N; i++ {
		Erat1(n, false, 0)
	}
}
