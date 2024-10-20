package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/tqt1345/Library-Go/model"
)

type Header struct {
	Heading  string
	NavItems []NavItem
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

// HEADER TEMPLATES

func BookCatalogueHeaderTemplate(w http.ResponseWriter, r *http.Request) {
	serveHeaderTemplate(w, "Book Catalogue")
}

func IndexHeaderTemplate(w http.ResponseWriter, r *http.Request) {
	serveHeaderTemplate(w, "Home")
}

// END HEADER TEMPLATES

// CONTENT TEMPLATES

func AllBooksTemplate(w http.ResponseWriter, r *http.Request) {
	books, err := repo.FindAllBooks()
	if err != nil {
		internalServerError(w, err)
		return
	}

	if books == nil {
		books = []model.Book{}
	}

	tmpl, err := template.ParseFiles(wd + "/view/templates/catalogue-table.html")
	if err != nil {
		internalServerError(w, err)
		return
	}

	w.Header().Set(ContentType, Html)
	tmpl.Execute(w, books)
}

func BookDetailsTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		badRequest(w, err)
		return
	}

	book, err := repo.FindBookById(id)
	if err != nil {
		notFound(w, err)
		return
	}

	templ, err := template.ParseFiles(wd + "/view/bookDetails.html")
	if err != nil {
		internalServerError(w, err)
		return
	}

	w.Header().Set(ContentType, Html)
	templ.Execute(w, book)
}

// END CONTENT TEMPLATES
