package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	str "strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func makePrimeSumEndpoint(pf PrimeFinder) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Printf("endpoint start")
		req := request.(sumRequest)
		log.Printf("req = %v\n", req)
		start := time.Now()
		sum, primes := pf.PrimeSum(req.N, req.List)
		etime := time.Since(start)
		log.Printf("sum=%v\nt=%v\np=%v\n", sum, etime.String(), primes)
		return sumResponse{
			sum,
			etime.String(),
			str.TrimRight(str.TrimLeft(fmt.Sprintf("%v", primes), "["), "]")}, nil
	}
}

type sumRequest struct {
	N    uint64 `json:"n"`
	List bool   `json:"list"`
}

type sumResponse struct {
	Sum    uint64 `json:"sum"`
	ETime  string `json:"etime"`
	Primes string `json:"primes,omitempty"`
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
