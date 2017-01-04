// Package primes provides functions to find sum
// of all primes lesser then predefined number n
// and optionaly get all of them as a slice.
// For n = 2000000 sum is 142913828922
// This is a 228112th solusion this problem on
// https://projecteuler.net/
package primes

import (
	"errors"
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
var PrimeSum func(uint64, bool) (uint64, []uint64, error) = Erat3

// Parallel sum calculation
func Erat3(n uint64, list bool) (uint64, []uint64, error) {
	if n < 2 {
		return 0, []uint64{}, ErrBadRange
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

	var rnum uint64 = 10 // Max # of goroutins
	var sums chan uint64 = make(chan uint64, rnum)
	var pnums chan uint64 = make(chan uint64, rnum)
	var ran = n / uint64(rnum)
	var imin, imax uint64
	for r := uint64(1); r <= rnum; r++ {
		imin, imax = imax, r*ran
		if r == rnum {
			imax = n
		}
		go func(s []bool) {
			var psum, ppnum, k uint64
			for k = 0; k < uint64(len(s)); k++ {
				if !s[k] {
					psum += k
					ppnum++
				}
			}
			sums <- psum
			pnums <- ppnum
		}(s[imin:imax])
	}

	for r := uint64(1); r <= rnum; r++ {
		sum += <-sums
		pnum += <-pnums
	}

	if list {
		return sum, make_prime_list(s, pnum), nil
	}

	return sum, nil, nil

}

var ErrBadRange error = errors.New("No primes in the range")
