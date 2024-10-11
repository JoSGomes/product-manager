package middleware

import (
	"golang.org/x/exp/slog"
	"net/http"
	"os"
)

var l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.Info("Received request", "method", r.Method, "url", r.URL.String(), "headers", r.Header)
		next.ServeHTTP(w, r)
	})
}
