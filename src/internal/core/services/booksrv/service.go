package services

import (
	"errors"
	"prakticas/backend-gpsoft/src/internal/core/domain"
	"prakticas/backend-gpsoft/src/internal/core/ports"
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
		return domain.Book{}, errors.New("Book doesn't exist")
	}

	return book, nil
}

func (srv *service) Create(name string) (domain.Book, error){
		id:= domain.
}

func (srv *service) Delete(id string) 
