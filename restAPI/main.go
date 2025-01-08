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

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

var books = []Book{}

func main() {
	users := readUsersFromFile()
	fmt.Print(users)

	router := gin.Default()
	books := readBookData()
	//fmt.Print(books)
	// Print the data
	// for _, book := range books {
	// 	fmt.Printf("Name: %s, Age: %s, Email: %s\n", book.Id, book.Title, book.Author)
	// }

	router.GET("/books", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, books)
	})
	router.POST("/books", postBooks)
	router.POST("/login", userLogin)

	router.Run("localhost:9090")

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

func userLogin(c *gin.Context) {
	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var creds Credentials

	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": "hello"})
	// for _, user := range users {
	//     if user.Username == creds.Username && user.Password == creds.Password {
	//         token, err := generateJWT(creds.Username)
	//         if err != nil {
	//             c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
	//             return
	//         }
	//         c.JSON(http.StatusOK, gin.H{"token": token})
	//         return
	//     }
	// }

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}
