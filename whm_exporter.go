package main
import (
    "log"

    "github.com/prometheus/client_golang/prometheus"

    "whm_exporter/cache"
    "whm_exporter/collector"
    "whm_exporter/config"
    "whm_exporter/server"
    "whm_exporter/whm"
)

func main() {
    cfg := config.Load()

    client := whm.NewClient(cfg.WHM)

    c := cache.New(cfg.WHM.CacheTTL)

    enabledCollectors := collector.LoadEnabled(
        cfg.Collectors.Enabled,
        client,
        c,
    )

    wc := collector.NewWHMCollector(enabledCollectors)

    reg := prometheus.NewRegistry()
    reg.MustRegister(wc)

    if err := server.Start(cfg.Server, reg); err != nil {
        log.Fatal(err)
    }
}
