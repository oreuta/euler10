package primes

import (
	"testing"
)

func BenchmarkErat2(b *testing.B) {
	var n int64 = 20000
	for i := 0; i < b.N; i++ {
		Erat2(n, false, 0)
	}
}
