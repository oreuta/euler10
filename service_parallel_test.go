package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/go-kit/kit/log"
)

// Skipped: client.Do(req) - problem
func TestPrimeService_Parallel(t *testing.T) {
	var logBuf bytes.Buffer
	logger := log.NewLogfmtLogger(&logBuf)
	sh := buildServiceHandler(logger, false)
	ts := httptest.NewServer(http.Handler(sh))
	defer ts.Close()
	url := fmt.Sprintf("%s/sum", ts.URL)
	bReqBody := []byte(
		`{"n": 2, "lst": false, "nr": 0}`)
	want := `"sum":2`

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			reqBody := bytes.NewBuffer(bReqBody)
			req, _ := http.NewRequest("POST", url, reqBody)
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			bResBody, err := ioutil.ReadAll(res.Body)
			defer res.Body.Close()
			if err != nil {
				t.Error(err)
				return
			}
			resBody := string(bResBody)
			if !strings.Contains(resBody, want) {
				t.Errorf("SERVICE: got  %vwant %v",
					res.Body, want)
			}
		}()
	}
	wg.Wait()
}
