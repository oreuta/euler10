package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/context"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	ctx := context.Background()
	svc := primeService{}

	primeSumHandler := httptransport.NewServer(
		ctx,
		makePrimeSumEndpoint(svc),
		decodePrimeSumRequest,
		encodeResponse,
	)

	http.HandleFunc("/", indexHandler)
	http.Handle("/sum", primeSumHandler)
	log.Fatal(http.ListenAndServe(":5050", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := ioutil.ReadFile("./ui/index.html")
	fmt.Fprint(w, string(page))
	return
}
