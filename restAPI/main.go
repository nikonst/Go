package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []book{
	{Id: "1", Title: "One", Author: "A1"},
	{Id: "2", Title: "Two", Author: "A2"},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.Run("localhost:9090")
}
