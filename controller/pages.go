package controller

import "net/http"

// BEGIN PAGES
func AllBooks(w http.ResponseWriter, r *http.Request) {
	servePage(w, r, "/view/books.html")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	servePage(w, r, "/view/index.html")
}

func BookDetails(w http.ResponseWriter, r *http.Request) {
	servePage(w, r, "/view/bookDetails.html")
}

// END PAGES
