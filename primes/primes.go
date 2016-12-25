//
//sum is 142913828922 (for n = 2000000)
//You are the 228112th person to have solved this problem.
//
package primes

import (
	"math"
)

//https://habrahabr.ru/post/133037/
func Erat2(n int, list bool) (int, []int) {
	s := make([]bool, n)    // Sieve: false-prime true-composite
	s[0], s[1] = true, true // 0 and 1 are composite

	i := 2 // first prime

	kmax := int(math.Sqrt(float64(n)))

	for i <= kmax { // sieve main loop

		for j := i * i; j < n; j += i {
			s[j] = true
		}

		i++
		for i <= kmax && s[i] {
			i++
		}
	}

	sum := 0
	pnum := 0

	for k := 2; k < n; k++ {
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
func Erat1(n int, list bool) (int, []int) {

	// All numbers are set to be prime (false)
	s := make([]bool, n)
	s[0], s[1] = true, true // 0 and 1 are composite

	i := 0    // first prime
	sum := 0  // sum of primes
	pnum := 0 // number of primes

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
func Erat0(n int, list bool) (int, []int) {

	// All numbers are set to be prime (false)
	s := make([]bool, n)
	s[0], s[1] = true, true // 0 and 1 are composite

	i := 2    // first prime
	sum := i  // sum of primes
	pnum := 1 // number of primes

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

func make_prime_list(s []bool, pnum int) []int {
	// Pick up primes from []bool to []int
	p := make([]int, pnum)
	for i, j := 0, 0; i < len(s); i++ {
		if !s[i] {
			p[j] = i
			j++
		}
	}
	return p
}
