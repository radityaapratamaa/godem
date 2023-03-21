package component

import "net/http"

type WrittenResponseWriter struct {
	http.ResponseWriter
	written bool
}

func (w *WrittenResponseWriter) WriteHeader(status int) {
	w.written = true
	w.ResponseWriter.WriteHeader(status)
}

func (w *WrittenResponseWriter) Write(b []byte) (int, error) {
	w.written = true
	return w.ResponseWriter.Write(b)
}

func (w *WrittenResponseWriter) Written() bool {
	return w.written
}

func InitHandler(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writtenResponseWriter := &WrittenResponseWriter{
			ResponseWriter: w,
			written:        false,
		}
		w = writtenResponseWriter

		router.ServeHTTP(w, r)
	})
}
