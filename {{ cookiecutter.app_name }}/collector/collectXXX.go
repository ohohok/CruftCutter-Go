package collector

import (
	"context"
	"database/sql"
	"std_exporter/common"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// mssql sample
func DatabaseStatus(db *sql.DB, ch chan<- prometheus.Metric) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	rows, err := db.QueryContext(ctx, `select @@SERVICENAME instance_name,name db_name,state,state_desc from sys.databases`)
	if err != nil {
		return err
	}
	for rows.Next() {
		var (
			instance string
			database string
			status   float64
			desc     string
		)
		if err := rows.Scan(&instance, &database, &status, &desc); err != nil {
			return err
		}
		ch <- prometheus.MustNewConstMetric(prometheus.NewDesc(prometheus.BuildFQName(common.NameSpace, "", "database_status"),
			"Gauge metric with 数据库状态", []string{"instance", "db_name"}, nil),
			prometheus.GaugeValue, status, instance, database)
	}
	return nil
}
