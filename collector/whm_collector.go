package collector
import (
    "sync"

    "github.com/prometheus/client_golang/prometheus"
)

type WHMCollector struct {
    collectors []Collector
}

func NewWHMCollector(cs []Collector) *WHMCollector {
    return &WHMCollector{collectors: cs}
}

func (w *WHMCollector) Describe(ch chan<- *prometheus.Desc) {}

func (w *WHMCollector) Collect(ch chan<- prometheus.Metric) {
    var wg sync.WaitGroup

    for _, c := range w.collectors {
        wg.Add(1)

        go func(col Collector) {
            defer wg.Done()
            _ = col.Update(ch)
        }(c)
    }

    wg.Wait()
}
