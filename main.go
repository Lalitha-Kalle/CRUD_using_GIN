package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID string `json:id`
	Title string `json:title`
	Author string `json:Author`
}

var books = []Book{
	{ID: "1", Title: "The Alchemist", Author: "Paulo Coelho"},
	{ID : "2", Title: "Atomic Habits", Author: "James Clear"},
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.POST("/books", createBook)

	fmt.Println("Server running at 2000")
	router.Run(":2000")
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBook(c  *gin.Context) {
	id := c.Param("id")
	for _,b := range books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func createBook(c *gin.Context) {
	var newBook Book 
	if err:= c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}