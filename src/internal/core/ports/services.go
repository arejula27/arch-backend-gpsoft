package ports

import "prakticas/backend-gpsoft/src/internal/core/domain"

type BookService interface {
	Get(id string) (domain.Book, error)
	Create(id string, name string) (domain.Book, error)
	Publish(id string) error
	Delete(id string) error
}
