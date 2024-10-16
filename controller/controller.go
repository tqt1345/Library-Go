package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/tqt1345/Library-Go/model"
)

var repo *model.Repository

const (
	JsonMime    = "application/json"
	HtmlMime    = "text/html"
	ContentType = "Content-Type"
)

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JsonMime)

	json.NewEncoder(w).Encode("Hello world!")
}

func ApiAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := repo.FindAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err)
	}

	if books == nil {
		books = []model.Book{}
	}

	w.Header().Set(ContentType, JsonMime)
	json.NewEncoder(w).Encode(books)
}

func ApiBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Print(err)
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

	w.Header().Set(ContentType, JsonMime)
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

	w.Header().Set(ContentType, JsonMime)
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

	w.Header().Set(ContentType, JsonMime)
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

	w.Header().Set(ContentType, JsonMime)
	json.NewEncoder(w).Encode(authors)
}

func Init(r *model.Repository) {
	repo = r

	// Book handlers
	http.HandleFunc("GET /api/", ApiIndexHandler)
	http.HandleFunc("GET /api/books/all", ApiAllBooksHandler)
	http.HandleFunc("GET /api/books/{id}", ApiBookByIdHandler)
	http.HandleFunc("GET /api/books/title", ApiBookByTitleHandler)

	// Author handlers
	http.HandleFunc("GET /api/authors/all", ApiAllAuthorsHandler)
	http.HandleFunc("GET /api/authors/firstName", ApiAuthorByFirstName)
}
