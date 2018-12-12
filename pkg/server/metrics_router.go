package server

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	getBooks = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "get_books_total",
			Help: "Number of get books",
		},
		[]string{"code", "method", "type"},
	)
)

func MetricsRouter(router *mux.Router) *mux.Router {
	router.Path("/metrics").Handler(promhttp.Handler())
	prometheus.MustRegister(getBooks)
	return router
}
