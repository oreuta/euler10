package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestRootHandler(t *testing.T) {
	rw := httptest.NewRecorder()
	rh := buildRootHandler()
	req, _ := http.NewRequest("GET", "/", nil)
	rh.ServeHTTP(rw, req)
	if !strings.Contains(rw.Body.String(),
		"Summation of primes") {
		t.Errorf("Unexpected output: %s", rw.Body)
	}
}

func TestPrimeServiceHandler(t *testing.T) {
	data := []struct {
		req, resp, log string
	}{
		{`{"n": 2, "lst": false, "nr": 0}`,
			`"sum":2`, `input=2[f]#0 sum=2 err=null`},
		{`{"n": 10, "lst": false, "nr": 0}`,
			`"sum":17`, `input=10[f]#0 sum=17 err=null`},
		{`{"n": 100, "lst": false, "nr": 0}`,
			`"sum":1060`, `input=100[f]#0 sum=1060 err=null`},
		{`{"n": -1, "lst": false, "nr": 0}`,
			`"error":"Bad range"`, `input=-1[f]#0 sum=0 err="Bad range"`},
		{`{"n": 1, "lst": false, "nr": 0}`,
			`"error":"No primes in the range"`, `input=1[f]#0 sum=0 err="No primes in the range"`},
		{`{"n": 2, "lst": true, "nr": 0}`,
			``, `primes="#1 [2]`},
		{`{"n": 7, "lst": true, "nr": 0}`,
			`"primes":[2,3,5,7]`, `primes="#4 [2 3 5 7]"`},
		{`{"n": 20, "lst": true, "nr": 0}`,
			`"primes":[2,3,5,7,11,13,17,19]`, `primes="#8 [2 3 5 7 11...]"`},

		{`bad request`,
			`invalid character`, ``},
	}

	var logBuf bytes.Buffer
	logger := log.NewLogfmtLogger(&logBuf)
	sh := buildServiceHandler(logger, true)
	url := "GET /sum HTTP/1.0\r\n\r\n"

	for _, d := range data {
		rw := httptest.NewRecorder()
		logBuf.Reset()
		body := bytes.NewBuffer([]byte(d.req))
		req, _ := http.NewRequest("POST", url, body)
		req.Header.Set("Content-Type", "application/json")
		sh.ServeHTTP(rw, req)
		if !strings.Contains(rw.Body.String(), d.resp) {
			t.Errorf("SERVICE:\nFor %v\ngot  %vwant %v",
				d.req, rw.Body, d.resp)
		}
		if !strings.Contains(logBuf.String(), d.log) {
			t.Errorf("LOGGING:\nFor %v\ngot  %vwant %v",
				d.req, logBuf.String(), d.log)
		}
	}
}
