package main

import "euler10/primes"

// PrimeFinder finds primes in a specified range and calculates their sum.
type PrimeFinder interface {
	PrimeSum(int64, bool) (int64, []int64)
}

type primeService struct{}

func (primeService) PrimeSum(n int64, list bool) (int64, []int64) {
	sum, primes := primes.PrimeSum(n, list)
	return sum, primes
}
