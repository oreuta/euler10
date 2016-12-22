package main

//Sieve of Eratosthenes - The Simplest Algorithm
func erat(n int) []int {

	type Number struct {
		val  int  // a number value
		comp bool // if it's composite (false = prime)
	}

	// Init []Sieve
	s := make([]Number, n)
	for i := 0; i < n; i++ {
		s[i].val = i
	}
	s[0].comp, s[1].comp = true, true // 0 and 1 are not primes

	i := 2      // first prime
	pnum := 0   // number of primes
	for i < n { // sieve main loop
		for j := i * i; j < n; j += i {
			s[j].comp = true
		}
		i++
		for i < n && s[i].comp { // find next prime
			i++
		}
		pnum++
	}

	// Pick up primes from []Sieve to []int
	p := make([]int, pnum)
	for i, j := 0, 0; i < n; i++ {
		if !s[i].comp {
			p[j] = s[i].val
			j++
		}
	}
	return p
}
