package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func makePrimeSumEndpoint(pf PrimeFinder) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Printf("endpoint start")
		req := request.(sumRequest)
		log.Printf("req = %v\n", req)
		sum, primes := pf.PrimeSum(req.N, req.List)
		log.Printf("sum = %v\np = %v\n", sum, primes)
		return sumResponse{sum, primes}, nil
	}
}

type sumRequest struct {
	N    int64 `json:"n"`
	List bool  `json:"list"`
}

type sumResponse struct {
	Sum    int64   `json:"sum"`
	Primes []int64 `json:"primes,omitempty"`
}

func decodePrimeSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request sumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}