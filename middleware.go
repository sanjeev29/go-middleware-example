package main

import (
	"log"
	"net/http"
	"time"
)

// LoggerHandler middleware
func LoggerHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		handler.ServeHTTP(w, r)
		log.Printf(
			"method=%s path=%s duration=%v",
			r.Method, r.URL.Path, time.Since(startTime),
		)
	})
}
