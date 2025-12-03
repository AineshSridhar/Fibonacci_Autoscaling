package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fib_requests_total",
			Help: "Total number of requests",
		},
		[]string{"path", "method"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "fib_request_duration_seconds",
			Help:    "Request duration histogram",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)

	PanicCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "fib_panics_total",
			Help: "Number of recovered panics",
		},
	)
)

func init() {
	prometheus.MustRegister(RequestCounter, RequestDuration, PanicCounter)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func() {
			if rec := recover(); rec != nil {
				PanicCounter.Inc()
			}
		}()
		next.ServeHTTP(w, r)

		RequestCounter.WithLabelValues(r.URL.Path, r.Method).Inc()
		RequestDuration.WithLabelValues(r.URL.Path, r.Method).
			Observe(time.Since(start).Seconds())
	})

}

func Handler() http.Handler {
	return promhttp.Handler()
}
