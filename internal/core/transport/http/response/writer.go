package core_http_response

import "net/http"

var (
	StatusCodeUninitialized = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statucCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statucCode:     StatusCodeUninitialized,
	}
}
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.statucCode = statusCode
}

func (rw *ResponseWriter) GetStatusCodeOrPanic() int {
	if rw.statucCode == StatusCodeUninitialized {
		panic("no status code set")
	}
	return rw.statucCode
}
