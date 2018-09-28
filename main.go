package main

import (
    "net/http"

    "github.com/gorilla/mux"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/weaveworks/common/middleware"
)

var (
	//RequestDuration a prometheus metric
    RequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
        Name:    "request_duration_seconds",
        Help:    "Time (in seconds) spent serving HTTP requests.",
        Buckets: prometheus.DefBuckets,
    }, []string{"method", "route", "status_code", "ws"})
)

func init() {
    prometheus.MustRegister(RequestDuration)
}

func main() {
    router := mux.NewRouter()
    router.Path("/metrics").Handler(prometheus.Handler())
    router.Path("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello world"))
    }))
    http.ListenAndServe(":8000", middleware.Instrument{
        Duration: RequestDuration,
    }.Wrap(router))
}