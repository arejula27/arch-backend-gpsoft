package domain

type Book struct {
	ID   string `json:"id"`
	name string `json:"name"`
}

func NewBook(id string, name string) Book {
	return Book{
		ID:   id,
		name: name,
	}
}
