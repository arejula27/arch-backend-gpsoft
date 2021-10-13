package ports

import "prakticas/backend-gpsoft/src/internal/core/domain"

type BookService interface {
	Get(id string) (domain.Book, error)
	Create(name string) (domain.Book, error)
	Edit(id string, name string)
	Delete(id string) error
}
