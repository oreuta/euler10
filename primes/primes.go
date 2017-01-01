// Package primes provides functions to find sum
// of all primes lesser then predefined number n
// and optionaly get all of them as a slice.
// For n = 2000000 sum is 142913828922
// This is a 228112th solusion this problem on
// https://projecteuler.net/
package primes

import (
	_ "errors"
	"math"
)

type Res struct {
	Sum    uint64
	Primes []uint64
}

type Params struct {
	N    uint64
	List bool
	Nrut uint8
}

// Default implementation
var PrimeSum func(uint64, bool) (uint64, []uint64, error) = Erat2

// https://habrahabr.ru/post/133037/
// 142913828922 - 19.0011ms
func Erat2(n uint64, list bool) (uint64, []uint64, error) {
	if n < 2 {
		return 0, []uint64{}, nil
	}

	var i uint64 = 2        // first prime
	var sum uint64 = 0      // sum of primes
	var pnum uint64 = 0     // number of primes
	s := make([]bool, n)    // Sieve: false-prime true-composite
	s[0], s[1] = true, true // 0 and 1 are composite

	kmax := uint64(math.Sqrt(float64(n)))

	for i <= kmax { // sieve main loop
		for j := i * i; j < n; j += i {
			s[j] = true
		}
		i++
		for i <= kmax && s[i] {
			i++
		}
	}

	for k := uint64(2); k < n; k++ {
		if !s[k] {
			sum += k
			pnum++
		}
	}

	if list {
		return sum, make_prime_list(s, pnum), nil
	}

	return sum, nil, nil

}
