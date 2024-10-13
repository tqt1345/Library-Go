package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const (
	JsonMime    = "application/json"
	HtmlMime    = "text/html"
	ContentType = "Content-Type"
)

var (
	queries *Queries
	ctx     context.Context
	ddl     string
)

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JsonMime)

	json.NewEncoder(w).Encode("Hello world!")
}

func ApiAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JsonMime)
	books, err := queries.ListBooks(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}

	if books == nil {
		books = []Book{}
	}

	json.NewEncoder(w).Encode(books)
}

func BookByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JsonMime)

	book, err := queries.BookId1(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(book)
}

func BookByTitleHandler(w http.ResponseWriter, r *http.Request) {
}

func BookByAuthorHandler(w http.ResponseWriter, r *http.Request) {
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

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}
	log.Print("Database connected...")

	queries = New(db)
	log.Print("Queries loaded...")

	port := ":80"

	http.HandleFunc("GET /api/", ApiIndexHandler)
	http.HandleFunc("GET /api/books", ApiAllBooksHandler)
	http.HandleFunc("GET /api/books/one", BookByIdHandler)
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
