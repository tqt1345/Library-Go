package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/tqt1345/Library-Go/model"
)

var (
	repo *model.Repository
	wd   string
	hd   Header
	nv   NavItems
)

const (
	Json        = "application/json"
	Html        = "text/html"
	ContentType = "Content-Type"
)

type Server struct {
	Repo *model.Repository
	Port string
}

func (s *Server) Start() {
	var err error
	wd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	nv = []NavItem{}
	nv.Add("Home", "/")
	nv.Add("Catalogue", "/books/catalogue")

	// File server
	fs := http.FileServer(http.Dir("./static"))
	http.HandleFunc("GET /static/", http.StripPrefix("/static/", fs).ServeHTTP)

	// Index
	http.HandleFunc("GET /", s.IndexHandler)
	log.Print("Starting...")

	// Book handlers
	http.HandleFunc("GET /api/books/all", s.ApiAllBooksHandler)
	http.HandleFunc("GET /api/books/{id}", s.ApiBookByIdHandler)
	http.HandleFunc("GET /api/books/title", s.ApiBookByTitleHandler)
	http.HandleFunc("GET /books/catalogue", s.AllBooks)
	http.HandleFunc("GET /books/details/{id}", s.BookDetailsTemplate)

	// Author handlers
	http.HandleFunc("GET /api/authors/all", s.ApiAllAuthorsHandler)
	http.HandleFunc("GET /api/authors/firstName", s.ApiAuthorByFirstName)

	// Html templates
	http.HandleFunc("GET /template/books/catalogue", s.AllBooksTemplate)
	http.HandleFunc("GET /template/headers/books", s.BookCatalogueHeaderTemplate)
	http.HandleFunc("GET /template/headers/index", s.IndexHeaderTemplate)

	// http.HandleFunc("GET /template/books/details/{id}", BookDetailsTemplate)
	http.ListenAndServe(s.Port, nil)
}
