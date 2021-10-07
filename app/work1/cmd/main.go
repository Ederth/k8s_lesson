package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

func main() {
	_ = startServer(context.Background())
}

func startServer(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandle)
	mux.HandleFunc("/healthz", healthzHandle)
	return http.ListenAndServe(":80", mux)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	logReq(r, statusCode)

	for k, v := range r.Header {
		for _, s := range v {
			w.Header().Add(k, s)
		}
	}
	w.Header().Add("Version", os.Getenv("VERSION"))

	w.WriteHeader(statusCode)
	w.Write([]byte("hello"))
}

func healthzHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func logReq(r *http.Request, code int) {
	fmt.Printf("client address: %s\nstatus code: %d", r.RemoteAddr, code)
}
