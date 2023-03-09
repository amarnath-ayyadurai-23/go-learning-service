// Package main Golang Learning Service REST API.
//
// Documentation for Learning REST API
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

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

// swagger:response BooksRes
//
//in:body
type _ struct {
	// ID
	//
	// example: 1
	ID string `json:"id"`
	// Title
	//
	// example: Book Name
	Title string `json:"title"`
	// Author
	//
	// example: Author Name
	Author string `json:"author"`
	// schema:
	//   type: application/json
}

// swagger:model BooksRes
// in:body
type Book struct {
	// ID
	//
	// example: 1
	ID string `json:"id"`
	// Title
	//
	// example: Book Name
	Title string `json:"title"`
	// Author
	//
	// example: Author Name
	Author string `json:"author"`
	// schema:
	//   type: application/json
}

// swagger:operation GET /books Books GettingAllBooks
//
// # Getting a List of all the Books
//
// ---
// produces:
// - application/json
// responses:
//
//	  "200":
//		   description: Books response (default)
//		   "$ref": "#/responses/BooksRes"
func getB(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

// swagger:operation POST /books Books CreatingBook
//
// # Creating a Book based on the info
//
// ---
// produces:
// - application/json
// responses:
//
//	  "200":
//		   "$ref": "#/responses/BooksRes"

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

// swagger:parameters DeletingBook UpdatingBook
type _ struct {
	// ID
	// in:path
	// required: true
	// schema:
	//   type: string
	ID string `json:"id"`
}

// swagger:operation DELETE /books/{id} Books DeletingBook
//
// # Deleting the records of a Book by ID
//
// ---
// produces:
// - application/json
// responses:
//
//	"204": HTTPResponse
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

// swagger:operation PATCH /books/{id} Books UpdatingBook
//
// # Updating a Book's Information
//
// ---
// produces:
// - application/json
// responses:
//
//	  "200":
//		   "$ref": "#/responses/BooksRes"

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
