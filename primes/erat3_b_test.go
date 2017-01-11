package primes

import (
	"testing"
)

func benchmarkErat3(b *testing.TB, nr int64) {
	b.ReportAllocs()
	var n int64 = 2000000
	for i := 0; i < b.N; i++ {
		Erat3(n, false, nr)
	}
}

func BenchmarkErat3w0(b *testing.TB) {
	benchmarkErat3(b, 0)
}

func BenchmarkErat3w1(b *testing.TB) {
	benchmarkErat3(b, 1)
}

func BenchmarkErat3w10(b *testing.TB) {
	benchmarkErat3(b, 10)
}

func BenchmarkErat3w100(b *testing.TB) {
	benchmarkErat3(b, 100)
}

func BenchmarkErat3w1000(b *testing.TB) {
	benchmarkErat3(b, 1000)
}
