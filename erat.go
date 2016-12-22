package main

//Sieve of Eratosthenes - The Simplest Algorithm
func erat(n int) []int {

	type Sieve struct {
		i int  // a number
		f bool // false = prime
	}

	// Init []Sieve
	s := make([]Sieve, n)
	for i := 0; i < n; i++ {
		s[i].i = i
	}
	s[0].f, s[1].f = true, true // 0 and 1 are not primes

	i := 2      // first prime
	pnum := 0   // number of primes
	for i < n { // sieve main loop
		for j := i * i; j < n; j += i {
			s[j].f = true
		}
		i++
		for i < n && s[i].f { // find next prime
			i++
		}
		pnum++
	}

	// Pick up primes from []Sieve to []int
	p := make([]int, pnum)
	for i, j := 0, 0; i < n; i++ {
		if !s[i].f {
			p[j] = s[i].i
			j++
		}
	}
	return p
}
