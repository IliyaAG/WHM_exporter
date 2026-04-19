package collector
import (
    "github.com/prometheus/client_golang/prometheus"
    "whm_exporter/cache"
    "whm_exporter/whm"
)

type accountCollector struct {
    client *whm.Client
    cache  *cache.Cache
}

func init() {
    Register("account", NewAccountCollector)
}

func NewAccountCollector(client *whm.Client, cache *cache.Cache) (Collector, error) {
    return &accountCollector{client: client, cache: cache}, nil
}

func (c *accountCollector) Name() string {
    return "account"
}

func (c *accountCollector) Update(ch chan<- prometheus.Metric) error {
    data, err := c.client.ListAccounts()
    if err != nil {
        return err
    }

    for _, acc := range data.Accounts {
        ch <- prometheus.MustNewConstMetric(
            prometheus.NewDesc(
                "whm_account_suspended",
                "Account suspended",
                []string{"user"},
                nil,
            ),
            prometheus.GaugeValue,
            boolToFloat(acc.Suspended),
            acc.User,
        )
    }

    return nil
}
