package main

import (
	conf "euler10/config"
	"fmt"
	"log"
)

func main() {
	n, err := conf.Getconf(10)
	if err != nil {
		log.Fatalf("Bad config: %v", err)
	}
	fmt.Printf("Hello euler %v config: %d\n", erat(n), n)
}
