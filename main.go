package main

import (
	"context"
	"database/sql"
	"encoding/json"
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

	json.NewEncoder(w).Encode("Hello world!")
}

func ApiAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := queries.FindAllBooks(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err)
	}

	if books == nil {
		books = []database.Book{}
	}

	w.Header().Set(ContentType, JsonMime)
	json.NewEncoder(w).Encode(books)
}

func ApiBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	// strId := r.PathValue("id")
	// i, err := strconv.Atoi(strId)
	// id := int64(i)

	// id, err := stringToInt64(r.PathValue("id"))
	// if err != nil {
	// 	log.Print(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	//
	id, err := stringToInt(r.PathValue("id"))
	if err != nil {
		log.Print(err)
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// book, err := queries.FindBookById(ctx, id)
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

	books, err := queries.FindBookByTitle(ctx, sql.NullString{String: title, Valid: true})
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set(ContentType, JsonMime)
	json.NewEncoder(w).Encode(books)
}

func ApiBookByAuthorHandler(w http.ResponseWriter, r *http.Request) {
}

func stringToInt64(s string) (int64, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return int64(i), nil
}

func stringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
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

	http.HandleFunc("GET /api/", ApiIndexHandler)
	http.HandleFunc("GET /api/books/all", ApiAllBooksHandler)
	http.HandleFunc("GET /api/books/{id}", ApiBookByIdHandler)
	http.HandleFunc("GET /api/books/title", ApiBookByTitleHandler)
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
