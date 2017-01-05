package main

import (
	"euler10/primes"
)

// PrimeFinder finds primes in a specified range and calculates their sum.
type PrimeService interface {
	PrimeSum(uint64, bool, uint8) (uint64, []uint64, error)
}

type primeService struct{}

func (primeService) PrimeSum(n uint64, lst bool, nr uint8) (uint64, []uint64, error) {
	sum, primes, err := primes.PrimeSum(n, lst, nr)
	return sum, primes, err
}
