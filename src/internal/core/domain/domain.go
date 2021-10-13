package domain

type Book struct {
	ID        string `json:"id"`
	name      string `json:"name"`
	published bool   `json:"published"`
}

func NewBook(id string, name string) Book {
	return newBook(id, name, false)
}

func newBook(id string, name string, published bool) Book {
	return Book{
		ID:        id,
		name:      name,
		published: published}

}

func (b Book) IsPublished() bool {
	return b.published
}
