package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

	// Open the JSON file
	jsonFile, err := os.Open("data/books.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// Read the JSON file
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON data into a slice of Person structs
	var books []Book
	if err := json.Unmarshal(byteValue, &books); err != nil {
		log.Fatal(err)
	}

	// Print the data
	for _, book := range books {
		fmt.Printf("Name: %s, Age: %s, Email: %s\n", book.Id, book.Title, book.Author)
	}

	router.GET("/books", getBooks)
	router.Run("localhost:9090")
}
