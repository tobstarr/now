package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	l := log.New(os.Stderr, "", 0)
	if err := run(l); err != nil {
		log.Fatal(err)
	}
}

func run(l *log.Logger) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if !strings.Contains(port, ":") {
		port = "0.0.0.0:" + port
	}
	l.Printf("running port %q", port)
	return http.ListenAndServe(port, http.HandlerFunc(index(time.Now)))
}

func index(now func() time.Time) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, now().UTC().Format(time.RubyDate))
		return
	}
}
