package domain

type Book struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Published bool   `json:"published"`
}

func NewBook(id string, name string) Book {
	return Book{
		ID:        id,
		Name:      name,
		Published: false}
}

func (b Book) IsPublished() bool {
	return b.Published
}
