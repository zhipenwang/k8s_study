package go_demo

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cast"
	"math/rand"
	"net/http"
	"time"
)

var HttpCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "test",
		Subsystem: "one",
		Name:      "request_total",
		Help:      "http request total",
	},
	[]string{"code", "method"},
)

var HttpSpeed = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: "test",
		Subsystem: "one",
		Name:      "request_total",
		Help:      "http request total",
		Buckets:   []float64{0.5, 1},
	},
	[]string{"code", "method"},
)

func RegisterPrometheus() {
	prometheus.Register(HttpCount)
}

func PromSrv(port int) {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+cast.ToString(port), nil)
}

func Register() {
	RegisterPrometheus()
}

func Incr() {
	t := time.NewTicker(time.Second)
	code := []string{"200", "500"}
	method := []string{"/demo", "test"}
	for {
		select {
		case <-t.C:
			HttpCount.WithLabelValues(code[rand.Intn(2)], method[rand.Intn(2)]).Add(1)
		}
	}
}
