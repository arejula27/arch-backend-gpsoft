package booksrv

import (
	"errors"
	"prakticas/backend-gpsoft/src/internal/core/domain"
	"prakticas/backend-gpsoft/src/internal/core/ports"
	"prakticas/backend-gpsoft/src/pkg/apperrors"
)

type service struct {
	bookRepository ports.BookRepository
}

func New(bookRepository ports.BookRepository) *service {
	return &service{bookRepository: bookRepository}
}

func (srv *service) Get(id string) (domain.Book, error) {
	book, err := srv.bookRepository.Get(id)
	if err != nil {
		return domain.Book{}, apperrors.ErrNotFound
	}

	return book, nil
}

func (srv *service) Create(id string, name string) (domain.Book, error) {
	id, err := srv.bookRepository.Save(domain.NewBook(id, name))
	if err != nil {
		return domain.Book{}, errors.New("no se ha creado el libro")
	}
	return domain.NewBook(id, name), nil
}

func (srv *service) Publish(id string) error {
	book, err := srv.bookRepository.Get(id)
	if err != nil {
		return errors.New("no existe el libro")
	}
	book.SetPublished(true)
	_, err = srv.bookRepository.Save(book)
	if err != nil {
		return errors.New("no se ha podido publicar el libro")
	}
	return nil
}

func (srv *service) Delete(id string) error {

	return srv.bookRepository.Delete(id)
}
