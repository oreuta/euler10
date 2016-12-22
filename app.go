package main

import (
	conf "euler10/config"
	"fmt"
	"log"
)

func main() {
	n, err := conf.Getconf(100)
	if err != nil {
		log.Fatalf("Bad config: %v", err)
	}
	fmt.Println("Hello euler 10!")
	primes := erat(n)
	fmt.Printf("act: %v\n", primes)
	fmt.Printf("exp: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]")
	fmt.Println()
}

func print_primes(n []int) {
	fmt.Printf("%v\n", n)
}
