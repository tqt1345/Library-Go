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

// func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello world")
// }

func (s *Server) Start() {
	var err error
	wd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// File server
	fs := http.FileServer(http.Dir("./static"))
	http.HandleFunc("GET /static/", http.StripPrefix("/static/", fs).ServeHTTP)

	// Index
	http.HandleFunc("GET /", s.IndexHandler)
	http.HandleFunc("GET /api/", s.ApiIndexHandler)
	log.Print("Starting...")

	// Book handlers
	http.HandleFunc("GET /api/books/all", s.ApiAllBooksHandler)

	http.HandleFunc("GET /api/books/{id}", s.ApiBookByIdHandler)
	http.HandleFunc("GET /api/books/title", s.ApiBookByTitleHandler)
	http.HandleFunc("GET /books/catalogue", s.AllBooks)
	http.HandleFunc("GET /books/details/{id}", s.BookDetailsTemplate)
	//
	// // Author handlers
	// http.HandleFunc("GET /api/authors/all", ApiAllAuthorsHandler)
	// http.HandleFunc("GET /api/authors/firstName", ApiAuthorByFirstName)
	//
	// // Html templates
	// http.HandleFunc("GET /template/books/catalogue", AllBooksTemplate)
	// http.HandleFunc("GET /template/headers/books", BookCatalogueHeaderTemplate)
	// http.HandleFunc("GET /template/headers/index", IndexHeaderTemplate)
	// // http.HandleFunc("GET /template/books/details/{id}", BookDetailsTemplate)
	http.ListenAndServe(s.Port, nil)
}

// func Init(r *model.Repository) {
// 	repo = r
//
// 	var err error
// 	wd, err = os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	nv = []NavItem{}
// 	nv.Add("Home", "/")
// 	nv.Add("Catalogue", "/books/catalogue")
//
// 	// File server
// 	fs := http.FileServer(http.Dir("./static"))
// 	http.HandleFunc("GET /static/", http.StripPrefix("/static/", fs).ServeHTTP)
//
// 	// Index
// 	http.HandleFunc("GET /", IndexHandler)
//
// 	// Book handlers
// 	http.HandleFunc("GET /api/", ApiIndexHandler)
// 	http.HandleFunc("GET /api/books/all", ApiAllBooksHandler)
// 	http.HandleFunc("GET /api/books/{id}", ApiBookByIdHandler)
// 	http.HandleFunc("GET /api/books/title", ApiBookByTitleHandler)
// 	http.HandleFunc("GET /books/catalogue", AllBooks)
// 	http.HandleFunc("GET /books/details/{id}", BookDetailsTemplate)
//
// 	// Author handlers
// 	http.HandleFunc("GET /api/authors/all", ApiAllAuthorsHandler)
// 	http.HandleFunc("GET /api/authors/firstName", ApiAuthorByFirstName)
//
// 	// Html templates
// 	http.HandleFunc("GET /template/books/catalogue", AllBooksTemplate)
// 	http.HandleFunc("GET /template/headers/books", BookCatalogueHeaderTemplate)
// 	http.HandleFunc("GET /template/headers/index", IndexHeaderTemplate)
// 	// http.HandleFunc("GET /template/books/details/{id}", BookDetailsTemplate)
// }
