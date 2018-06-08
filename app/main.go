package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr    = flag.String("addr", ":8080", "The address to listen on")
	summary = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "foo_milliseconds",
		Help:       "Time taken to render foo",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"code"})
)

func init() {
	prometheus.MustRegister(summary)
}

func fooHandler(summary *prometheus.SummaryVec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer r.Body.Close()

		defer func() {
			duration := time.Since(start)
			ms := duration.Seconds() * 1e3
			log.Println(ms)
			summary.WithLabelValues(fmt.Sprintf("%d", 200)).Observe(ms)
		}()

		resp, err := http.Get("http://fakethirdpartysite/")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, string(body))
	}
}

func main() {

	flag.Parse()

	http.Handle("/api/foo", fooHandler(summary))
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(*addr, nil))
}
