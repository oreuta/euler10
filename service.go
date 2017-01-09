package main

import (
	"euler10/primes"
)

// PrimeFinder finds primes in a specified range and calculates their sum.
type PrimeService interface {
	PrimeSum(int64, bool, int64) (int64, []int64, error)
}

type primeService struct{}

func (primeService) PrimeSum(n int64, lst bool, nr int64) (int64, []int64, error) {
	sum, primes, err := primes.PrimeSum(n, lst, nr)
	return sum, primes, err
}
