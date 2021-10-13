package main

import (
	"prakticas/backend-gpsoft/src/internal/core/services/booksrv"
	"prakticas/backend-gpsoft/src/internal/dataSources/bookrepo"
	"prakticas/backend-gpsoft/src/internal/handlers/bookhdl"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	bookRepository := bookrepo.NewMemKVS()
	bookService := booksrv.New(bookRepository)
	bookHandler := bookhdl.NewHTTPHandler(bookService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("book/:id", bookHandler.Get)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
