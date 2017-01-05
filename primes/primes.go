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

// Default implementation for Summation of primes.
// It gets n - an upper limit for primes in the sum,
// lst - if list of primes has to be generated,
// nr - number of goroutines used.
var PrimeSum func(n uint64, lst bool, nr uint8) (sum uint64, primes []uint64, err error) = Erat3

// Parallel sum calculation
func Erat3(n uint64, lst bool, nr uint8) (uint64, []uint64, error) {
	if n < 2 {
		return 0, []uint64{}, ErrEmptyRange
	}

	var i uint64 = 2        // first prime
	var sum uint64 = 0      // sum of primes
	var pnum uint64 = 0     // number of primes
	s := make([]bool, n+1)  // Sieve: false-prime true-composite
	s[0], s[1] = true, true // 0 and 1 are composite

	kmax := uint64(math.Sqrt(float64(n)))

	for i <= kmax { // sieve main loop
		for j := i * i; j <= n; j += i {
			s[j] = true
		}
		i++
		for i <= kmax && s[i] {
			i++
		}
	}

	var rnum uint64 // Max # of goroutins
	if nr == 0 {
		rnum = 3 // default for now...
	}
	var sums chan uint64 = make(chan uint64, rnum)
	var pnums chan uint64 = make(chan uint64, rnum)
	var ran = (n + 1) / uint64(rnum)
	var imin, imax uint64
	for r := uint64(1); r <= rnum; r++ {
		imin, imax = imax, r*ran
		if r == rnum {
			imax = n + 1
		}

		go func(s []bool, imin uint64, imax uint64) {
			var psum, ppnum, k uint64
			for k = imin; k < imax; k++ {
				if !s[k] {
					psum += k
					ppnum++
				}
			}
			sums <- psum
			pnums <- ppnum
		}(s, imin, imax)
	}

	for r := uint64(1); r <= rnum; r++ {
		sum += <-sums
		pnum += <-pnums
	}

	if lst {
		return sum, makePrimeLst(s, pnum), nil
	}

	return sum, nil, nil

}

var ErrEmptyRange error = errors.New("No primes in the range")
var ErrBadRange error = errors.New("Bad range")
