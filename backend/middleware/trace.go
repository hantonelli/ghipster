package middleware

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

type contextKey string

func (c contextKey) String() string {
	return "github.com/hantonelli/ghipster context key " + string(c)
}

const TracingKey contextKey = "tracing"

var tracingHeaders = []string{
	"x-request-id",
	"x-b3-traceid",
	"x-b3-spanid",
	"x-b3-sampled",
	"x-b3-parentspanid",
	"x-b3-flags",
	"x-ot-span-context",
}

func TracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tracing []string
		for _, key := range tracingHeaders {
			if val := r.Header.Get(key); val != "" {
				tracing = append(tracing, key, val)
			}
		}
		if 0 == len(tracing) {
			next.ServeHTTP(w, r)
		} else {
			ctx := context.WithValue(r.Context(), TracingKey, tracing)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func AddTraceToLog(ctx context.Context, log *zap.SugaredLogger) *zap.SugaredLogger {
	trace := ctx.Value(TracingKey)
	tr, ok := trace.([]string)
	if ok && 0 < len(tr) {
		for i := 0; i+1 < len(tr); i += 2 {
			log = log.With(tr[i], tr[i+1])
		}
	}
	return log
}
