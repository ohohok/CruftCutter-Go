package collector

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

type ScrapeFn func(db *sql.DB, ch chan<- prometheus.Metric) error

// A map used to call the collect metrics method
var ScrapeFns = map[string]func(db *sql.DB, ch chan<- prometheus.Metric) error{
	"DatabaseStatus": DatabaseStatus,
}
