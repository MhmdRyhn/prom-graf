package app

import (
	"fmt"
	"net/http"

	"github.com/MhmdRyhn/poke"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(poke.NewMetrics().WithCounterVector(
		"monitoring", "promgraf", "request_count",
	).WithHistogramVector(
		"monitoring", "promgraf", "latency_histogram", nil,
	).Middleware)
	r.Get("/hello-world", HelloWorldHandler)
	r.Mount("/metrics", promhttp.Handler())
	return r
}

func Start(port int) error {
	r := Router()
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	return err
}
