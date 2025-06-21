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

	fmt.Println("Server running at 2000")
	router.Run(":2000")
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}