package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	user         = kingpin.Flag("user", "PostgreSQL username").Default("fusionpbx").String()
	password     = kingpin.Flag("password", "PostgreSQL password").Default("password").String()
	dbname       = kingpin.Flag("dbname", "PostgreSQL database name").Default("fusionpbx").String()
	host         = kingpin.Flag("host", "PostgreSQL host").Default("localhost").String()
	port         = kingpin.Flag("port", "PostgreSQL port").Default("5432").String()
	listenAddr   = kingpin.Flag("web.listen-address", "Address to listen on for web interface and metrics").Default(":9240").String()
)

var reg = prometheus.NewPedanticRegistry()

func main() {
	kingpin.Parse()

	go func() {
		for {
			CollectMetrics()
			time.Sleep(10 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		log.Fatal("Failed to start HTTP server: ", err)
	}
}
