package web

import (
	"log"
	"net/http"
	"shortlink/models"
	"time"
)

type Middeware struct {
}

func (m Middeware) LoggingHandler(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %q %v", r.Method, r.URL.String(), time.Now().Sub(startTime))
	}
	return http.HandlerFunc(f)
}

func (m Middeware) RecoverHandler(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover from panic:%+v", err)
				sendErrorResponse(w, models.ErrorInternalFaults)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}
