package server
import (
    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start(cfg struct{ Port string }, reg *prometheus.Registry) error {
    http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
    return http.ListenAndServe(":"+cfg.Port, nil)
}

