package main

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	ctx := context.Background()
	svc := PrimeService{}

	primeSumHandler := httptransport.NewServer(
		ctx,
		makePrimeSumEndpoint(svc),
		decodePrimeSumRequest,
		encodeResponse,
	)

	fs := http.FileServer(http.Dir("ui"))
	http.Handle("/", fs)
	http.Handle("/sum", primeSumHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
