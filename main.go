package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"time"
)

// Embed Next.js exported site
//go:embed client/out/*
//go:embed client/out/*/*
//go:embed client/out/*/*/*
//go:embed client/out/*/*/*/*
//go:embed client/out/*/*/*/*/*
//go:embed client/out/*/*/*/*/*/*
//go:embed client/out/*/*/*/*/*/*/*
var content embed.FS

func main() {
	// Return the built ui as the filesystem root
	client, err := fs.Sub(content, "client/out")
	if err != nil {
		panic(err)
	}

	// Serve static files of built ui
	http.Handle("/", http.FileServer(http.FS(client)))

	// Define ping endpoint that responds with pong
	http.HandleFunc("/api/ping", pingHandler())

	// Start the server
	log.Println("Server starting on: http://localhost:3000/")
	if err := http.ListenAndServe(":3000", nil); err != nil {
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
