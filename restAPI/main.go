package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()

	dataFile, err := os.Open("data/books.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var response map[string][]Book
	err = json.NewDecoder(dataFile).Decode(&response)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer dataFile.Close()

	router.GET("/books", getBooks)
	router.Run("localhost:9090")
}
