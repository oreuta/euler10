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

	//
	//sum is 142913828922 (for n = 2000000)
	//You are the 228112th person to have solved this problem.
	//
	fmt.Printf("exp: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]")
	fmt.Println("\n-----------------")

	ts := time.Now()
	sum, p := primes.Erat0(n, list)
	te := time.Now()
	pres(sum, ts, te, p, list)
	fmt.Println("==============")
	ts = time.Now()
	sum, p = primes.Erat1(n, list)
	te = time.Now()
	pres(sum, ts, te, p, list)
	fmt.Println("==============")
	ts = time.Now()
	sum, p = primes.Erat2(n, list)
	te = time.Now()
	pres(sum, ts, te, p, list)

}

func pres(sum int, ts time.Time, te time.Time, p []int, list bool) {
	if list {
		fmt.Printf("act: %v\n", p)
	}
	fmt.Printf("Sum: %v ", sum)
	fmt.Printf("Dur: %v\n", te.Nanosecond()-ts.Nanosecond())
}
