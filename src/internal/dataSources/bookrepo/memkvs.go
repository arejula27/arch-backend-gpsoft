package bookrepo

import (
	"encoding/json"
	"prakticas/backend-gpsoft/src/internal/core/domain"
	"prakticas/backend-gpsoft/src/pkg/apperrors"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo *memkvs) Get(id string) (domain.Book, error) {
	if value, ok := repo.kvs[id]; ok {
		book := domain.Book{}
		err := json.Unmarshal(value, &book)
		if err != nil {
			return domain.Book{}, apperrors.ErrInternal
		}

		return book, nil
	}
	return domain.Book{}, apperrors.ErrNotFound
}

func (repo *memkvs) Save(book domain.Book) (string, error) {
	bytes, err := json.Marshal(book)

	if book.ID == "" {

	}
	if err != nil {
		return "", apperrors.ErrInvalidInput
	}

	repo.kvs[book.ID] = bytes
	keys := make([]string, 0, len(repo.kvs))
	for k := range repo.kvs {
		keys = append(keys, k)
	}

	return book.ID, nil
}

func (repo *memkvs) Delete(id string) error {
	return nil
}
