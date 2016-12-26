// Package primes provides functions to find sum
// of all primes lesser then predefined number n
// and optionaly get all of them as a slice.
// For n = 2000000 sum is 142913828922
// This is a 228112th solusion this problem on
// https://projecteuler.net/
package primes

import (
	"math"
)

// Default implementation of
var PrimeSum func(int64, bool) (int64, []int64) = Erat2

// https://habrahabr.ru/post/133037/
func Erat2(n int64, list bool) (int64, []int64) {
	i := int64(2)           // first prime
	sum := int64(0)         // sum of primes
	pnum := int64(0)        // number of primes
	s := make([]bool, n)    // Sieve: false-prime true-composite
	s[0], s[1] = true, true // 0 and 1 are composite

	kmax := int64(math.Sqrt(float64(n)))

	for i <= kmax { // sieve main loop

		for j := i * i; j < n; j += i {
			s[j] = true
		}

		i++
		for i <= kmax && s[i] {
			i++
		}
	}

	for k := int64(2); k < n; k++ {
		if !s[k] {
			sum += k
			pnum++
		}
	}

	// prepare list of primes if needed
	if list {
		return sum, make_prime_list(s, pnum)
	}

	return sum, nil

}

//Sieve of Eratosthenes - The Simplest Algorithm
//Parallel version 1 (fixed g-rout number)
func Erat1(n int64, list bool) (int64, []int64) {
	var i int64 = 0         // first prime
	var sum int64 = 0       // sum of primes
	var pnum int64 = 0      // number of primes
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
	if list {
		return sum, make_prime_list(s, pnum)
	}

	return sum, nil

}

func rmcomp(s []bool, i int, c chan int) {
	for j := i * i; j < len(s); j += i {
		s[j] = true
	}
	c <- i
}

//Sieve of Eratosthenes - The Simplest Algorithm
func Erat0(n int64, list bool) (int64, []int64) {
	var i int64 = 2         // first prime
	var sum int64 = i       // sum of primes
	var pnum int64 = 1      // number of primes
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
	if list {
		return sum, make_prime_list(s, pnum)
	}

	return sum, nil

}

func make_prime_list(s []bool, pnum int64) []int64 {
	// Pick up primes from []bool to []int
	p := make([]int64, pnum)
	for i, j := int64(0), int64(0); i < int64(len(s)); i++ {
		if !s[i] {
			p[j] = i
			j++
		}
	}
	return p
}
