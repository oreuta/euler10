package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	port := os.Getenv("PORT") // for heroku
	if port == "" {
		port = 8080
	}
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
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
