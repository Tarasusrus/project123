package middleware

import (
	"context"
	"net/http"
	"sync/atomic"

	"github.com/google/uuid"
)

var req uint64

type key int

const RequestIDKey key = 0

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			// Генерируем уникальный идентификатор
			requestID = uuid.New().String()
		} else {
			// Увеличиваем счетчик запросов, чтобы получить уникальный идентификатор
			requestID = uuid.New().String() + "-" + string(atomic.AddUint64(&req, 1))
		}

		// Вставляем requestID в контекст запроса
		ctx = context.WithValue(ctx, RequestIDKey, requestID)

		// Добавляем requestID в заголовок ответа
		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}
