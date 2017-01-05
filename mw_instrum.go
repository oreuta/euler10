package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           PrimeService
}

func (mw instrumentingMiddleware) PrimeSum(n uint64, lst bool, nr uint64) (sum uint64, primes []uint64, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "PrimeSum", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	sum, primes, err = mw.next.PrimeSum(n, lst, nr)
	return
}
