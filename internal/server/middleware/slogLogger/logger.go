package slogLogger

import (
	"BaseApi/internal/server/middleware"
	"log/slog"
	"net/http"
	"time"
)

func New(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log = log.With(slog.String(
			"component", "middleware/slogLogger"),
		)
		log.Info("slogLogger middleware enabled")
		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_adr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				slog.String("request_id", middleware.GetRequestID(r.Context())),
			)

			ww := middleware.WrapResponseWriter(w)

			t1 := time.Now()
			defer func() {
				entry.Info("request completed",
					slog.Int("status", ww.Status()),
					slog.Int("bytes", ww.BytesWritten()),
					slog.String("duratione", time.Since(t1).String()),
				)
			}()

			next.ServeHTTP(ww, r)

		}

		return http.HandlerFunc(fn)
	}
}
