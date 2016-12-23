package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello euler 10!")
	n := 2000000
	sum, primes := erat(n, false)
	//
	//sum is 142913828922 (for n = 2000000)
	//You are the 228112th person to have solved this problem.
	//
	fmt.Printf("act: %v\n", primes)
	fmt.Printf("exp: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]")
	fmt.Printf("\nSum: %v\n", sum)
}

func print_primes(n []int) {
	fmt.Printf("%v\n", n)
}
