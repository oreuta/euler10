package main

import (
	"net/http"
	"os"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	port := os.Getenv("PORT") // for heroku
	if port == "" {
		port = "8080"
	}

	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "euler10",
		Subsystem: "prime_sum_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "euler10",
		Subsystem: "prime_sum_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	var svc PrimeService
	svc = primeService{}
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, svc}

	primeSumHandler := httptransport.NewServer(
		ctx,
		makePrimeSumEndpoint(svc),
		decodePrimeSumRequest,
		encodeResponse,
	)

	fs := http.FileServer(http.Dir("ui"))
	http.Handle("/", fs)
	http.Handle("/sum", primeSumHandler)
	http.Handle("/metrics", stdprometheus.Handler())

	logger.Log("msg", "HTTP", "addr", ":"+port)
	logger.Log("err", http.ListenAndServe(":"+port, nil))
}
