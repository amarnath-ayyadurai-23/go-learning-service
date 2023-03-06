// Package main Golang Learning Service Microservice API.
//
// Documentation for Learning API
//
//	Schemes: http, https
//	Host: localhost
//	Description: This is a basic Learning APIs based on VCS service
//	BasePath: /books
//	Version: 0.0.1
//	License: CC https://creativecommons.org/licenses/
//	Contact: Amar<amarnath.ayyadurai@pantheon.io>
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func getB(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func postB(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	books = append(books, book)

	c.JSON(http.StatusCreated, book)
}

func deleteB(c *gin.Context) {
	id := c.Param("id")

	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}

func patchB(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, a := range books {
		if a.ID == book.ID {
			books[i] = book
			break
		}
	}
	c.JSON(http.StatusCreated, book)
}

func main() {
	r := gin.New()

	r.GET("/books", getB)
	r.POST("/books", postB)
	r.DELETE("/books/:id", deleteB)
	r.PATCH("/books/:id", patchB)

	r.Run(":3000")
}
