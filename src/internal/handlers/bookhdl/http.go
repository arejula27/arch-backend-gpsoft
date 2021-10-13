package bookhdl

import (
	"prakticas/backend-gpsoft/src/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	bookService ports.BookService
}

func NewHTTPHandler(bookService ports.BookService) *HTTPHandler {
	return &HTTPHandler{
		bookService: bookService,
	}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	book, err := hdl.bookService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, book)
}

func (hdl *HTTPHandler) Save(c *gin.Context) {
	book, err := hdl.bookService.Create(c.Param("id"), c.Param("name"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return

	}
	c.JSON(200, book)
}
