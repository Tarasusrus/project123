package middleware

import (
	"bufio"
	"net"
	"net/http"
)

// ResponseWriter is an interface that extends the standard http.ResponseWriter
type ResponseWriter interface {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	http.Pusher
	Status() int
	BytesWritten() int
	Unwrap() http.ResponseWriter
}

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
	bytes       int
}

func WrapResponseWriter(w http.ResponseWriter) ResponseWriter {
	bw := &responseWriter{ResponseWriter: w}
	bw.status = http.StatusOK // Default status code
	return bw
}

func (w *responseWriter) WriteHeader(code int) {
	if !w.wroteHeader {
		w.status = code
		w.ResponseWriter.WriteHeader(code)
		w.wroteHeader = true
	}
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.WriteHeader(http.StatusOK) // Ensure the header is written
	n, err := w.ResponseWriter.Write(b)
	w.bytes += n
	return n, err
}

func (w *responseWriter) Status() int {
	return w.status
}

func (w *responseWriter) BytesWritten() int {
	return w.bytes
}

func (w *responseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

func (w *responseWriter) Flush() {
	if fl, ok := w.ResponseWriter.(http.Flusher); ok {
		fl.Flush()
	}
}

func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hj, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, http.ErrNotSupported
	}
	return hj.Hijack()
}

func (w *responseWriter) Push(target string, opts *http.PushOptions) error {
	if ps, ok := w.ResponseWriter.(http.Pusher); ok {
		return ps.Push(target, opts)
	}
	return http.ErrNotSupported
}
