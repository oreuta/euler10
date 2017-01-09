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
		sum, primes, err := ps.PrimeSum(req.N, req.Lst, req.Nr)
		if err != nil {
			return sumError{err.Error()}, nil
		}
		etime := time.Since(start)
		return sumResponse{
			sum,
			etime.String(),
			primes}, nil
	}
}

type sumRequest struct {
	N   int64 `json:"n"`
	Lst bool  `json:"lst"`
	Nr  int64 `json:"nr"`
}

type sumResponse struct {
	Sum    int64   `json:"sum"`
	ETime  string  `json:"etime"`
	Primes []int64 `json:"primes,omitempty"`
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
