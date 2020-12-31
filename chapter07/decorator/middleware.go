package decorator

import (
	"log"
	"net/http"
	"time"
)

//Logger é um dos nossos decoradores de 'middleware'
func Logger(l *log.Logger) Decorator {
	return func(c http.RoundTripper) http.RoundTripper {
		return TransportFunc(func(r *http.Request) (*http.Response, error) {
			start := time.Now()
			l.Printf("requisição iniciada para %s em %s", r.URL, start.Format("2020-01-02 14:30:00"))

			resp, err := c.RoundTrip(r)
			l.Printf("requisição concluida para %s em %s", r.URL, time.Since(start))

			return resp, err
		})
	}
}

//BasicAuth é outro de nossos decoradores de 'middleware'
func BasicAuth(username, password string) Decorator {
	return func(c http.RoundTripper) http.RoundTripper {
		return TransportFunc(func(r *http.Request) (*http.Response, error) {
			r.SetBasicAuth(username, password)
			resp, err := c.RoundTrip(r)
			return resp, err
		})
	}
}
