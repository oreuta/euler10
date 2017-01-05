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

func (mw loggingMiddleware) PrimeSum(n uint64, lst bool, nr uint8) (sum uint64, primes []uint64, err error) {
	var plen int
	var p string

	defer func(begin time.Time) {
		var t rune
		if lst {
			t = 't'
		} else {
			t = 'f'
		}
		_ = mw.logger.Log(
			"input", fmt.Sprintf("%v[%c]#%v", n, t, nr),
			"sum", sum,
			"err", err,
			"took", time.Since(begin).String(),
			"primes", fmt.Sprintf("#%v [%v]", len(primes), p),
		)
	}(time.Now())

	sum, primes, err = mw.next.PrimeSum(n, lst, nr)

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
