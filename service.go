package main

import "euler10/primes"

// PrimeFinder finds primes in a specified range and calculates their sum.
type PrimeFinder interface {
	PrimeSum(uint64, bool) (uint64, []uint64)
}

type PrimeService struct{}

func (PrimeService) PrimeSum(n uint64, list bool) (uint64, []uint64) {
	sum, primes := primes.PrimeSum(n, list)
	return sum, primes
}
