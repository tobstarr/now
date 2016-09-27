package main

import (
	"bytes"
	"encoding/json"
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
		rsp := struct {
			CurrentTime time.Time `json:"current_time"`
		}{
			CurrentTime: now().UTC(),
		}
		buf := &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(rsp)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		io.Copy(w, buf)
		return
	}
}
