package main

//Sieve of Eratosthenes - The Simplest Algorithm
func erat(n int) ([]int, int) {

	// All numbers are set to be prime (false)
	s := make([]bool, n)
	s[0], s[1] = true, true // 0 and 1 are composite

	i := 2      // first prime
	pnum := 0   // number of primes
	for i < n { // sieve main loop
		for j := i * i; j < n; j += i {
			s[j] = true
		}
		i++
		for i < n && s[i] { // find next prime
			i++
		}
		pnum++
	}

	// Pick up primes from []bool to []int
	// and calculate a sum
	p := make([]int, pnum)
	sum := 0
	for i, j := 0, 0; i < n; i++ {
		if !s[i] {
			p[j] = i
			sum += i
			j++
		}
	}
	return p, sum
}
