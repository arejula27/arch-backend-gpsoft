package ports

import "prakticas/backend-gpsoft/src/internal/core/domain"

type BookRepository interface {
	Get(id string) (domain.Book, error)
	Save(domain.Book) (domain.Book, error)
	Delete(id string) error
}
