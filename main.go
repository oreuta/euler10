package main

import (
	"fmt"
	"time"

	"euler10/primes"
)

func main() {
	fmt.Println("Hello euler 10!")
	n := 2000000
	list := false

	fmt.Printf("exp: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]")
	fmt.Println("\n-----------------")

	start := time.Now()
	sum, p := primes.Erat0(int64(n), list)
	elapsed := time.Since(start)
	pres(sum, elapsed, p, list)
	fmt.Println("==============")

	start = time.Now()
	sum, p = primes.Erat1(int64(n), list)
	elapsed = time.Since(start)
	pres(sum, elapsed, p, list)
	fmt.Println("==============")

	start = time.Now()
	sum, p = primes.Erat2(int64(n), list)
	elapsed = time.Since(start)
	pres(sum, elapsed, p, list)
	fmt.Println("==============")

	start = time.Now()
	sum, p = primes.Primesum(int64(n), list)
	elapsed = time.Since(start)
	pres(sum, elapsed, p, list)

}

func pres(sum int64, elapsed time.Duration, p []int64, list bool) {
	if list {
		fmt.Printf("act: %v\n", p)
	}
	fmt.Printf("Sum: %v ", sum)
	fmt.Printf("Dur: %dns\n", elapsed.Nanoseconds())
}
