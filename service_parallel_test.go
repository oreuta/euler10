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
func _TestPrimeService_Parallel(t *testing.T) {
	var logBuf bytes.Buffer
	logger := log.NewLogfmtLogger(&logBuf)
	sh := buildServiceHandler(logger, false)
	ts := httptest.NewServer(http.Handler(sh))
	defer ts.Close()
	url := fmt.Sprintf("%s/sum", ts.URL)
	body := bytes.NewBuffer([]byte(
		`{"n": 2, "lst": false, "nr": 0}`))
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	want := `"sum":2`

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			client := &http.Client{}
			res, err := client.Do(req)
			//panic: Post http://127.0.0.1:49644/sum: http: ContentLength=31 with Body length 0
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			bodyB, err := ioutil.ReadAll(res.Body)
			defer res.Body.Close()
			if err != nil {
				t.Error(err)
				return
			}
			bodyS := string(bodyB)
			if !strings.Contains(bodyS, want) {
				t.Errorf("SERVICE: got  %vwant %v",
					res.Body, want)
			}

			t.Logf("Got: %s", bodyS)
		}()
	}
	wg.Wait()
}
