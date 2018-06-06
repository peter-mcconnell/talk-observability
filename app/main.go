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
	addr      = flag.String("addr", ":8080", "The address to listen on")
	histogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "foo_milliseconds",
		Help:    "Time taken to render foo",
		Buckets: prometheus.LinearBuckets(200, 100, 10),
	}, []string{"code"})
)

func init() {
	prometheus.MustRegister(histogram)
}

func fooHandler(histogram *prometheus.HistogramVec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer r.Body.Close()

		defer func() {
			duration := time.Since(start)
			ms := duration.Seconds() * 1e3
			ms *= 200 // fake http://fakethirdpartysite/ feeling "real"
			log.Println(ms)
			histogram.WithLabelValues(fmt.Sprintf("%d", 200)).Observe(ms)
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

	http.Handle("/api/foo", fooHandler(histogram))
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(*addr, nil))
}
