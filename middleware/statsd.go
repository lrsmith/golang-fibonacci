package middleware

import (
	"log"
	"net/http"
	"strings"

	statsd "github.com/smira/go-statsd"
)

// StatsdClient Global variable so main can create connection to statsd
var StatsdClient *statsd.Client

// StatsdMiddleware ...
func StatsdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Replace '/' in RequestURI with '.' for metric name. Graphite format
		// Then replace parameter with '.' so get value
		// ex. metric would be .v1.fibseq.8
		uriString := strings.Replace(r.RequestURI, "/", ".", -1)
		metric := strings.Replace(uriString, "?index=", ".", -1)

		log.Printf("Metric : %s", metric)
		StatsdClient.Incr(metric, 1)

		next.ServeHTTP(w, r)
	})
}
