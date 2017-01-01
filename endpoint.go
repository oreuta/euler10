package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func makePrimeSumEndpoint(ps PrimeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(sumRequest)
		start := time.Now()
		sum, primes, err := ps.PrimeSum(req.N, req.List)
		if err != nil {
			return sumError{"ERROR: " + err.Error()}, nil
		}
		etime := time.Since(start)
		return sumResponse{
			sum,
			etime.String(),
			primes}, nil
	}
}

type sumRequest struct {
	N    uint64 `json:"n"`
	List bool   `json:"list"`
}

type sumResponse struct {
	Sum    uint64   `json:"sum"`
	ETime  string   `json:"etime"`
	Primes []uint64 `json:"primes,omitempty"`
}

type sumError struct {
	Err string `json:"error"`
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
