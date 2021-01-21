package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Book ...
type Book struct {
	Title  string
	Author string
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")

	books := make([]Book, 0)
	books = append(books, Book{
		Title:  "It",
		Author: "Stephen King",
	})
	books = append(books, Book{
		Title:  " Harry Potter and the Philosopher's Stone",
		Author: "J. K. Rowling",
	})
	books = append(books, Book{
		Title:  "The Monk Who Sold His Ferrari",
		Author: "Robin Sharma",
	})
	books = append(books, Book{
		Title:  "Five Point Someone",
		Author: "Chetan Bhagat",
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"books": books,
		})
	})
	log.Fatal(r.Run())
}
