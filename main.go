package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tqt1345/Library-Go/database"
	"github.com/tqt1345/Library-Go/model"
)

const (
	JsonMime    = "application/json"
	HtmlMime    = "text/html"
	ContentType = "Content-Type"
)

var (
	repo    *model.Repository
	queries *database.Queries
	ctx     context.Context
	ddl     string
)

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JsonMime)

	encode(w, "Hello World")
}

func ApiAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := repo.FindAllBooks()
	if err != nil {
		InternalServerError(w, err)
	}

	if books == nil {
		books = []model.Book{}
	}

	encode(w, books)
}

func ApiBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		BadRequest(w, err)
		return
	}

	book, err := repo.FindBookById(id)
	if err != nil {
		NotFound(w, err)
		return
	}

	encode(w, book)
}

func ApiBookByTitleHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var title string

	if params["title"] == nil {
		err := errors.New("Bad request param")
		BadRequest(w, err)
		return
	}
	title = params.Get("title")

	books, err := repo.FindBooksByTitle(title)
	if err != nil {
		NotFound(w, err)
		return
	}

	if books == nil {
		books = []model.Book{}
	}

	encode(w, books)
}

func NotFound(w http.ResponseWriter, err error) {
	log.Print(err.Error())
	http.Error(w, err.Error(), http.StatusNotFound)
}

func ApiBookByAuthorHandler(w http.ResponseWriter, r *http.Request) {
}

func ApiAllAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	authors, err := repo.FindAllAuthors()
	if err != nil {
		InternalServerError(w, err)
	}

	if authors == nil {
		authors = []model.Author{}
	}

	encode(w, authors)
}

func ApiAuthorByFirstName(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if params["firstName"] == nil {
		err := errors.New("Bad request param")
		BadRequest(w, err)
		return
	}

	name := params.Get("firstName")
	authors, err := repo.FindAuthorsByFirstName(name)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	if authors == nil {
		authors = []model.Author{}
	}

	encode(w, authors)
}

func BadRequest(w http.ResponseWriter, err error) {
	log.Print(err.Error())
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func InternalServerError(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func encode(w http.ResponseWriter, i any) {
	w.Header().Set(ContentType, JsonMime)
	json.NewEncoder(w).Encode(i)
}

func stringToInt64(s string) (int64, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return int64(i), nil
}

func run() error {
	ctx = context.Background()

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	// if _, err := db.ExecContext(ctx, ddl); err != nil {
	// 	return err
	// }

	log.Print("Database connected...")

	queries = database.New(db)
	log.Print("Queries loaded...")

	repo = model.NewRepo(db)
	log.Print("Repository loaded...")

	port := ":8080"

	// Book handlers
	http.HandleFunc("GET /api/", ApiIndexHandler)
	http.HandleFunc("GET /api/books/all", ApiAllBooksHandler)
	http.HandleFunc("GET /api/books/{id}", ApiBookByIdHandler)
	http.HandleFunc("GET /api/books/title", ApiBookByTitleHandler)

	// Author handlers
	http.HandleFunc("GET /api/authors/all", ApiAllAuthorsHandler)
	http.HandleFunc("GET /api/authors/firstName", ApiAuthorByFirstName)

	log.Print("Routes loaded...")

	log.Printf("Server started on port%s", port)
	http.ListenAndServe(port, nil)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
