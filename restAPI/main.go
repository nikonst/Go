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
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

var books = []Book{}

// postAlbums adds an album from JSON received in the request body.
func postBooks(c *gin.Context) {
	var newBook Book

	// Call BindJSON to bind the received JSON to
	// newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new album to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
	fmt.Print(books)
}

func readBookData() []Book {
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
	return books

}

func main() {
	users := readUsersFromFile()
	fmt.Print(users)
	router := gin.Default()
	books := readBookData()
	fmt.Print(books)
	// Print the data
	// for _, book := range books {
	// 	fmt.Printf("Name: %s, Age: %s, Email: %s\n", book.Id, book.Title, book.Author)
	// }

	router.GET("/books", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, books)
	})
	router.POST("/books", postBooks)

	router.Run("localhost:9090")

}
