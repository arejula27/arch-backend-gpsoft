package ports

import "prakticas/backend-gpsoft/src/internal/core/domain"

type BookRepository interface {
	Get(id string) (domain.Book, error)
	Save(domain.Book) error
}

type GameService interface {
	Get(id string) (domain.Book, error)
	Create(name string) (domain.Book, error)
	Delete(id string)
}
