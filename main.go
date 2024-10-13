package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	JsonMime    = "application/json"
	HtmlMime    = "text/html"
	ContentType = "Content-Type"
)

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JsonMime)
	json.NewEncoder(w).Encode("Hello world!")
}

func AllBooksHandler(w http.ResponseWriter, r *http.Request) {
}

func BookByIdHandler(w http.ResponseWriter, r *http.Request) {
}

func BookByTitleHandler(w http.ResponseWriter, r *http.Request) {
}

func BookByAuthorHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	port := ":80"

	http.HandleFunc("GET /api/", ApiIndexHandler)

	log.Printf("Starting server on port%s", port)
	http.ListenAndServe(port, nil)
}
