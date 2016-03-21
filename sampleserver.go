package sampleserver

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello, world!\n")
	return
}

func bye(w http.ResponseWriter, r *http.Request) {
	log.Println("")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Good bye, this world..\n")
	return
}

// CreateMux create mux
func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/hello", loggingMiddleware(http.HandlerFunc(hello)))
	mux.Handle("/bye", loggingMiddleware(http.HandlerFunc(bye)))
	return mux
}

// CreateServer create server
func CreateServer() *http.Server {
	mux := CreateMux()
	server := &http.Server{
		Handler: mux,
		Addr:    "localhost:8810",
	}
	return server
}
