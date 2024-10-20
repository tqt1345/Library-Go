package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/tqt1345/Library-Go/controller"
	"github.com/tqt1345/Library-Go/model"
)

const (
	JsonMime    = "application/json"
	HtmlMime    = "text/html"
	ContentType = "Content-Type"
)

func main() {
	db := model.NewDB()
	repo := &model.Repository{Db: db}
	server := controller.Server{Repo: repo, Port: ":8080"}
	server.Start()
}
