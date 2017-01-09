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
	"runtime"
)

// Default implementation for Summation of primes.
// It gets n - an upper limit for primes in the sum,
// lst - if list of primes has to be generated,
// nr - number of goroutines used.
var PrimeSum func(n int64, lst bool, nr int64) (sum int64, primes []int64, err error) = Erat3

// Parallel sum calculation
func Erat3(n int64, lst bool, nr int64) (int64, []int64, error) {
	if n < 0 {
		return 0, []int64{}, ErrBadRange
	}
	if n < 2 {
		return 0, []int64{}, ErrEmptyRange
	}

	var i int64 = 2         // first prime
	var sum int64 = 0       // sum of primes
	var pnum int64 = 0      // number of primes
	s := make([]bool, n+1)  // Sieve: false-prime true-composite
	s[0], s[1] = true, true // 0 and 1 are composite

	kmax := int64(math.Sqrt(float64(n)))

	for i <= kmax { // sieve main loop
		for j := i * i; j <= n; j += i {
			s[j] = true
		}
		i++
		for i <= kmax && s[i] {
			i++
		}
	}

	var rnum int64 // Max # of goroutines
	if nr == 0 {
		rnum = int64(runtime.NumCPU()) // default!
	} else {
		rnum = nr // user defined
	}

	var sums chan int64 = make(chan int64, rnum)
	var pnums chan int64 = make(chan int64, rnum)
	var ran = (n + 1) / rnum
	var imin, imax int64

	for r := int64(1); r <= rnum; r++ {
		imin, imax = imax, r*ran
		if r == rnum {
			imax = n + 1
		}

		go func(s []bool, imin int64, imax int64) {
			var psum, ppnum, k int64
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

	for r := int64(1); r <= rnum; r++ {
		sum += <-sums
		if sum <= 0 {
			return 0, []int64{}, ErrOverflow
		}
		pnum += <-pnums
	}

	if lst {
		return sum, makePrimeLst(s, pnum), nil
	}

	return sum, nil, nil

}

var (
	ErrEmptyRange error = errors.New("No primes in the range")
	ErrBadRange   error = errors.New("Bad range")
	ErrOverflow   error = errors.New("Overflow occurred")
)
