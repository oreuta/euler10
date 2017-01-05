package primes

import (
	"math"
)

//Sieve of Eratosthenes - The Simplest Algorithm

// https://habrahabr.ru/post/133037/
// 142913828922 - 19.0011ms
func Erat2(n uint64, lst bool, nr uint64) (uint64, []uint64, error) {
	if n < 2 {
		return 0, []uint64{}, ErrEmptyRange
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

	if lst {
		return sum, makePrimeLst(s, pnum), nil
	}

	return sum, nil, nil

}

//Parallel version 1 (fixed g-rout number)
func Erat1(n uint64, lst bool, nr uint64) (uint64, []uint64, error) {
	if n < 2 {
		return 0, []uint64{}, ErrEmptyRange
	}

	var i uint64 = 0        // first prime
	var sum uint64 = 0      // sum of primes
	var pnum uint64 = 0     // number of primes
	s := make([]bool, n)    // Sieve: false-prime true-composite
	s[0], s[1] = true, true // 0 and 1 are composite

	ch := make(chan int, 4)

	go rmcomp(s, 2, ch)
	go rmcomp(s, 3, ch)
	go rmcomp(s, 5, ch)
	go rmcomp(s, 7, ch)
	go rmcomp(s, 11, ch)
	go rmcomp(s, 13, ch)

	sum = 2 + 3 + 5 + 7 + 11 + 13
	pnum = 6

	<-ch
	<-ch
	<-ch
	<-ch
	<-ch
	<-ch

	i = 13

	for i < n { // sieve main loop

		// go routine starts here ----
		for j := i * i; j < n; j += i {
			s[j] = true
		}
		// ----------------------------

		//find next prime, count it, update sum
		i++
		for i < n && s[i] {
			i++
		}
		pnum++
		sum += i
	}

	// prepare list of primes if needed
	if lst {
		return sum, makePrimeLst(s, pnum), nil
	}

	return sum, nil, nil

}

func rmcomp(s []bool, i int, c chan int) {
	for j := i * i; j < len(s); j += i {
		s[j] = true
	}
	c <- i
}

//Sieve of Eratosthenes - The Simplest Algorithm
func Erat0(n uint64, lst bool, nr uint64) (uint64, []uint64, error) {
	if n < 2 {
		return 0, []uint64{}, ErrEmptyRange
	}

	var i uint64 = 2        // first prime
	var sum uint64 = i      // sum of primes
	var pnum uint64 = 1     // number of primes
	s := make([]bool, n)    // Sieve: false-prime true-composite
	s[0], s[1] = true, true // 0 and 1 are composite

	for i < n { // sieve main loop

		// go routine starts here ----
		for j := i * i; j < n; j += i {
			s[j] = true
		}
		// ----------------------------

		//find next prime, count it, update sum
		i++
		for i < n && s[i] {
			i++
		}
		pnum++
		sum += i
	}

	// prepare list of primes if needed
	if lst {
		return sum, makePrimeLst(s, pnum), nil
	}

	return sum, nil, nil

}

func makePrimeLst(s []bool, pnum uint64) []uint64 {
	// Pick up primes from []bool to []uint64
	p := make([]uint64, pnum)
	for i, j := uint64(0), uint64(0); i < uint64(len(s)); i++ {
		if !s[i] {
			p[j] = i
			j++
		}
	}
	return p
}
