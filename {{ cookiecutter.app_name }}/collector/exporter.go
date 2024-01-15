package collector

import (
	"database/sql"
	"os"
	"os/signal"
	"std_exporter/common"
	"std_exporter/config"
	"std_exporter/logger"
	"syscall"

	"github.com/prometheus/client_golang/prometheus"
)

var logging = logger.GetStdLogger()

// Exporter interface
type ExporterFtx interface {
	// An interface including prometheus's collector
	prometheus.Collector
	// Provide a method to close the interface
	Stop()
}

// Exporter collects DB metrics. It implements prometheus.Collector.
type Exporter struct {
	Type    string
	Version string
	Config  *config.Config
	up      prometheus.Gauge
}

// NewExporter returns a new DB exporter by the provided DSN.
func New(conf *config.Config) *Exporter {
	e := &Exporter{
		Config: conf,
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: common.NameSpace,
			Name:      "up",
			Help:      "Whether the XXX is up.",
		}),
	}
	return e
}

// Describe describes all the metrics exported by exporter.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	metricCh := make(chan prometheus.Metric)
	doneCh := make(chan struct{})

	go func() {
		for m := range metricCh {
			ch <- m.Desc()
		}
		close(doneCh)
	}()

	e.Collect(metricCh)
	close(metricCh)
	<-doneCh

}

// Collect implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	logging.Infoln("collecting data")

	e.scrape(ch)
	ch <- e.up
}

// A specific method for collecting indicators in a database
func (e *Exporter) scrape(ch chan<- prometheus.Metric) {
	db, err := sql.Open("DB", e.Config.Dsn)
	if err != nil {
		logging.Errorln("Error opening connection to database:", err)
		e.up.Set(0)
		return
	}
	defer db.Close()
	e.up.Set(1)

	/*
		Additional requirements code
	*/

	for name, fn := range ScrapeFns {
		if err := fn(db, ch); err != nil {
			logging.Errorf("Error scraping for %s : %s", name, err)
		}
	}
}

// Listening signal
func ListenSignal(fn func()) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	logging.Infoln("Exporter shut down")
	fn()
	_ = os.Remove(logger.LogPath)
	os.Exit(0)
}

// Output stop log
func (e *Exporter) Stop() {
	logging.Infoln("Exporter already stopped!")
}
