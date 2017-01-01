package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   PrimeService
}

func (mw loggingMiddleware) PrimeSum(n uint64, list bool) (sum uint64, primes []uint64, err error) {
	var plen int
	var p string

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"input", fmt.Sprintf("%v(%t)", n, list),
			"sum", sum,
			"err", err,
			"took", time.Since(begin).String(),
			"primes", fmt.Sprintf("#%v [%v]", len(primes), p),
		)
	}(time.Now())

	sum, primes, err = mw.next.PrimeSum(n, list)

	// nice output for primes
	plen = len(primes)
	var pend string
	if plen > 5 {
		plen = 5
		pend = "..."
	}
	p = fmt.Sprintf("%v", primes[0:plen])
	if len(p) > 2 {
		p = p[1:len(p)-1] + pend
	} else {
		p = pend
	}
	return
}
