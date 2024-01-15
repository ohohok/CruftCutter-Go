package main

import (
	"flag"
	"net/http"
	"std_exporter/collector"
	"std_exporter/common"
	"std_exporter/config"
	"std_exporter/logger"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Defining the basic content
var (
	Version       = common.Version
	NameSpace     = common.NameSpace
	listenAddress = flag.String("web.listen-address", ":9601", "Address to listen on for web interface and telemetry.")
	metricPath    = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	landingPage   = []byte("<html><head><title>" + NameSpace + " " + Version + "</title></head><body><h1>" + NameSpace +
		" " + Version + "</h1><p><a href='" + *metricPath + "'>Metrics</a></p></body></html>")
)

// Program entry, collect command line content, register Exporter and listening
func main() {
	var host, port, user, password, logLevel string
	flag.StringVar(&host, "host", "", "host")
	flag.StringVar(&port, "port", "", "port")
	flag.StringVar(&user, "user", "", "user")
	flag.StringVar(&password, "password", "", "password")
	flag.StringVar(&logLevel, "logLevel", "error", "logLevel")
	var passEncrypt bool
	flag.BoolVar(&passEncrypt, "passEncrypt", false, "password is or not encrypt")
	flag.Parse()
	logging := logger.GetStdLogger()
	logging.SetLevel(logLevel)
	if host == "" || port == "" {
		logging.Errorln("host or port is needed")
	}
	if user == "" {
		logging.Errorln("user is needed")
	}
	if password == "" {
		logging.Errorln("password is null,please set password")
	}
	// 密码加密
	if passEncrypt {
		logging.Debugln("Encrypted password: ", password)
		passwd, err := common.RSADecrypt(password, []byte(common.PrivateKey))
		if err == nil {
			password = string(passwd)
			logging.Debugln("Decrypted password: ", password)
		} else {
			logging.Errorln("RSADecrypt Error: ", err, "  [Try to use plain password.....]")
		}
	}

	logging.Infoln("Starting mssql_exporter " + Version)
	conf := &config.Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		LogLevel: logLevel,
		Dsn:      "", //Fill in the DSN according to the actual situation
	}
	exporter := collector.New(conf)
	// The method is typically used exporter.Stop
	go collector.ListenSignal(exporter.Stop)
	registry := prometheus.NewRegistry()
	registry.MustRegister(exporter)
	http.Handle(*metricPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(landingPage)
		if err != nil {
			return
		}
	})
	logging.Infoln("Listening on", *listenAddress)
	logging.Fatal(http.ListenAndServe(*listenAddress, nil))
}
