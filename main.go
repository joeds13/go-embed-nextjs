package main

import (
	"encoding/json"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

//go:generate go run embed_gen.go
func main() {
	devMode := flag.Bool("dev", false, "Dev mode proxies to dev mode client on http://localhost:3000/")
	flag.Parse()

	// if in dev mode, proxy through to localhost:3000, don't embed/fileserver
	// this allows next dev with hot reloading etc and a sane dev workflow
	if *devMode == true {
		log.Println("Proxying client to: http://localhost:3000/")
		origin, _ := url.Parse("http://localhost:3000/")

		director := func(req *http.Request) {
			req.Header.Add("X-Forwarded-Host", req.Host)
			req.Header.Add("X-Origin-Host", origin.Host)
			req.URL.Scheme = "http"
			req.URL.Host = origin.Host
		}

		proxy := &httputil.ReverseProxy{Director: director}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			proxy.ServeHTTP(w, r)
		})
	} else {
		// Return the built client as the filesystem root
		// content is provided by go generating embed_gen.go
		client, err := fs.Sub(content, "client/out")
		if err != nil {
			panic(err)
		}

		// Serve static files of built client
		http.Handle("/", http.FileServer(http.FS(client)))
	}

	// Define ping endpoint that responds with pong
	http.HandleFunc("/api/ping", pingHandler())

	// Start the server
	log.Println("Server starting on: http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type pong struct {
	Pong int64 `json:"pong"`
}

// pingHandler returns pong
func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data = pong{time.Now().UnixNano()}
		jData, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	}
}
