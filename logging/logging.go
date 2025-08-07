package logging

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request: ", r.Header.Get("method"), " ", r.Header.Get("request-target"), " ", r.Header.Get("protocol"))
		f(w, r)
	}
}
