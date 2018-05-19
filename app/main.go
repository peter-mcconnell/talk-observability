package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr      = flag.String("addr", ":8080", "The address to listen on")
	histogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "foo_seconds",
		Help:    "Time taken to render foo",
		Buckets: prometheus.LinearBuckets(0.3, 0.2, 10),
	}, []string{"hostname", "code"})
	hostname, _ = os.Hostname()
)

func init() {
	prometheus.MustRegister(histogram)
}

func fooHandler(histogram *prometheus.HistogramVec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer r.Body.Close()

		// using os.Hostname() as a label as scape config is load balanced (and
		// as such prints app:8080 for all containers). for demo purposes
		defer func() {
			duration := time.Since(start)
			log.Println(duration.Seconds())
			histogram.WithLabelValues(hostname, fmt.Sprintf("%d", 200)).Observe(duration.Seconds())
		}()

		resp, err := http.Get("https://api.github.com/users/pemcconnell/repos")
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
