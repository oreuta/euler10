package primes

import (
	"testing"
)

func benchmarkErat3(b *testing.B, nr int64) {
	var n int64 = 20000
	for i := 0; i < b.N; i++ {
		Erat0(n, false, nr)
	}
}

func BenchmarkErat3w0(b *testing.B) {
	benchmarkErat3(b, 0)
}

func BenchmarkErat3w1(b *testing.B) {
	benchmarkErat3(b, 1)
}

func BenchmarkErat3w10(b *testing.B) {
	benchmarkErat3(b, 10)
}

func BenchmarkErat3w100(b *testing.B) {
	benchmarkErat3(b, 100)
}

func BenchmarkErat3w1000(b *testing.B) {
	benchmarkErat3(b, 1000)
}
