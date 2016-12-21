package main

import (
	conf "euler10/config"
	"fmt"
	"log"
)

func main() {
	n, err := conf.Getconf()
	if err != nil {
		log.Fatal("Bad config")
	}
	fmt.Printf("Hello euler config: %d\n", n)
}
