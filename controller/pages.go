package controller

import (
	"fmt"
	"net/http"
)

// BEGIN PAGES
func (s *Server) AllBooks(w http.ResponseWriter, r *http.Request) {
	servePage(w, r, "/view/books.html")
}

func wtf(w http.ResponseWriter) {
	fmt.Fprintf(w, "Hello world")
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	servePage(w, r, "/view/index.html")
}

func BookDetails(w http.ResponseWriter, r *http.Request) {
	servePage(w, r, "/view/bookDetails.html")
}

// END PAGES
