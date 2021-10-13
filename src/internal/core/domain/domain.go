package domain

type Book struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Published bool   `json:"published"`
}

func NewBook(id string, name string) *Book {
	return &Book{
		ID:        id,
		Name:      name,
		Published: false}
}

func (b Book) IsPublished() bool {
	return b.Published
}

func (b *Book) SetId(id string) {
	b.ID = id

}

func (b *Book) SetName(name string) {
	b.Name = name

}

func (b *Book) SetPublished(published bool) {
	b.Published = published

}
