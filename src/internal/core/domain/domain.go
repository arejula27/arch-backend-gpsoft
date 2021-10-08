package domain

type Book struct {
	ID   string `json:"id"`
	name string `json:"name"`
}

func NewBook(name string) *Book {
	return &Book{bookRepository: bookRepository}
}
