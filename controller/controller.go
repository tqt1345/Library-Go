package controller

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/tqt1345/Library-Go/model"
)

type Header struct {
	Heading  string
	NavItems []NavItem
}

func BookCatalogueHeaderTemplate(w http.ResponseWriter, r *http.Request) {
	h := Header{"Book Catalogue", nv}

	tmpl, err := template.ParseFiles(wd + "/view/fragments/header.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, Html)
	tmpl.Execute(w, h)
}

type NavItem struct {
	Title string
	Url   string
}

type NavItems []NavItem

func (n *NavItems) Add(title, url string) {
	// *n = append(*n, NavItem{Title: title, Url: title})
	*n = append(*n, NavItem{Title: title, Url: url})
}

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

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, Json)

	json.NewEncoder(w).Encode("Hello world!")
}

func ApiAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := repo.FindAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	if books == nil {
		books = []model.Book{}
	}

	w.Header().Set(ContentType, Json)
	json.NewEncoder(w).Encode(books)
}

func ApiBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := repo.FindBookById(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set(ContentType, Json)
	json.NewEncoder(w).Encode(book)
}

func ApiBookByTitleHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var title string

	if params["title"] == nil {
		log.Print("Bad Request")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	title = params.Get("title")

	books, err := repo.FindBooksByTitle(title)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if books == nil {
		books = []model.Book{}
	}

	w.Header().Set(ContentType, Json)
	json.NewEncoder(w).Encode(books)
}

func ApiBookByAuthorHandler(w http.ResponseWriter, r *http.Request) {
}

func ApiAllAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	authors, err := repo.FindAllAuthors()
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if authors == nil {
		authors = []model.Author{}
	}

	w.Header().Set(ContentType, Json)
	json.NewEncoder(w).Encode(authors)
}

func ApiAuthorByFirstName(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if params["firstName"] == nil {
		err := errors.New("Bad request")
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := params.Get("firstName")
	authors, err := repo.FindAuthorsByFirstName(name)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if authors == nil {
		authors = []model.Author{}
	}

	w.Header().Set(ContentType, Json)
	json.NewEncoder(w).Encode(authors)
}

func AllBooksTemplate(w http.ResponseWriter, r *http.Request) {
	books, err := repo.FindAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	if books == nil {
		books = []model.Book{}
	}

	tmpl, err := template.ParseFiles(wd + "/view/fragments/catalogue-table.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, Html)
	tmpl.Execute(w, books)
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set(ContentType, Html)
	// http.ServeFile(w, r, wd+"/view/books.html")
	servePage(w, r, "/view/books.html")
}

func IndexHeaderTemplate(w http.ResponseWriter, r *http.Request) {
	h := Header{"Home", nv}

	tmpl, err := template.ParseFiles(wd + "/view/fragments/header.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, Html)
	tmpl.Execute(w, h)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	servePage(w, r, "/view/index.html")
}

func servePage(w http.ResponseWriter, r *http.Request, filePath string) {
	w.Header().Set(ContentType, Html)
	http.ServeFile(w, r, wd+filePath)
}

func Init(r *model.Repository) {
	repo = r

	var err error
	wd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	nv = []NavItem{}
	nv.Add("Home", "/")
	nv.Add("Catalogue", "/books/catalogue")

	// Index
	http.HandleFunc("GET /", IndexHandler)

	// Book handlers
	http.HandleFunc("GET /api/", ApiIndexHandler)
	http.HandleFunc("GET /api/books/all", ApiAllBooksHandler)
	http.HandleFunc("GET /api/books/{id}", ApiBookByIdHandler)
	http.HandleFunc("GET /api/books/title", ApiBookByTitleHandler)
	http.HandleFunc("GET /books/catalogue", AllBooks)

	// Author handlers
	http.HandleFunc("GET /api/authors/all", ApiAllAuthorsHandler)
	http.HandleFunc("GET /api/authors/firstName", ApiAuthorByFirstName)

	// Html templates
	http.HandleFunc("GET /template/books/catalogue", AllBooksTemplate)
	http.HandleFunc("GET /template/headers/books", BookCatalogueHeaderTemplate)
	http.HandleFunc("GET /template/headers/index", IndexHeaderTemplate)
}
