package collector
import (
    "sync"

    "whm_exporter/cache"
    "whm_exporter/whm"
)

type Collector interface {
    Name() string
    Update(ch chan<- prometheus.Metric) error
}

type Factory func(client *whm.Client, cache *cache.Cache) (Collector, error)

var (
    factories = map[string]Factory{}
    mu        sync.Mutex
)

func Register(name string, factory Factory) {
    mu.Lock()
    defer mu.Unlock()
    factories[name] = factory
}

func LoadEnabled(names []string, client *whm.Client, cache *cache.Cache) []Collector {
    var result []Collector

    for _, name := range names {
        if f, ok := factories[name]; ok {
            c, err := f(client, cache)
            if err == nil {
                result = append(result, c)
            }
        }
    }

    return result
}
