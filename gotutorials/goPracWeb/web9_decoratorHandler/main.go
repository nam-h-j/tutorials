package main

import (
	"log"
	"net/http"
	"time"

	"./myapp"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOOGER1] GO!")
	h.ServeHTTP(w, r)
	log.Print("[LOOGER1] Completed time:", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	return mux
}

func main() {
	mux := NewHandler()
	http.ListenAndServe(":1234", mux)
}
