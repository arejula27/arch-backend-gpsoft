package booksrv_test

import (
	"prakticas/backend-gpsoft/src/internal/core/domain"
	"prakticas/backend-gpsoft/src/internal/core/services/booksrv"
	"prakticas/backend-gpsoft/src/mocks/mockups"
	"prakticas/backend-gpsoft/src/pkg/apperrors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	bookRepository *mockups.MockBookRepository
}

func TestGet(t *testing.T) {

	// · Mocks · //
	book := simplemockBook("1", "Quijote")

	// · Test · //
	type args struct {
		id string
	}
	type want struct {
		result domain.Book
		err    error
	}
	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mocks)
	}{
		{
			name: "Should return a book sccesfully",
			args: args{id: "1"},
			want: want{result: book},
			mocks: func(m mocks) {
				m.bookRepository.EXPECT().Get("1").Return(book, nil)
			},
		},
		{
			name: "Should return error - book not found",
			args: args{id: "2"},
			want: want{err: apperrors.ErrNotFound},
			mocks: func(m mocks) {
				m.bookRepository.EXPECT().Get("2").Return(domain.Book{}, apperrors.ErrNotFound)
			},
		},
	}
	// · Runner · //
	for _, tt := range tests {
		//Prepare

		m := mocks{
			bookRepository: mockups.NewMockBookRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := booksrv.New(m.bookRepository)

		//Execute
		result, err := service.Get(tt.args.id)

		//Verify
		if tt.want.err != nil && err != nil {
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)

	}

}

//TODO
func TestCreate(t *testing.T) {

	// · Mocks · //
	book := simplemockBook("1", "Quijote")

	// · Test · //
	type args struct {
		id   string
		name string
	}
	type want struct {
		result domain.Book
		err    error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mocks)
	}{
		{
			name: "Should save a book sccesfully",
			args: args{id: "1", name: "quijote"},
			want: want{result: book},
			mocks: func(m mocks) {
				m.bookRepository.EXPECT().Save(domain.NewBook("1", "quijote")).Return(book, nil)
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {
		//Prepare

		m := mocks{
			bookRepository: mockups.NewMockBookRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := booksrv.New(m.bookRepository)

		//Execute
		result, err := service.Create(tt.args.id, tt.args.name)

		//Verify
		if tt.want.err != nil && err != nil {
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)

	}

}

//TODO
func TestPublish(t *testing.T) {}

//TODO
func TestDelete(t *testing.T) {}

func simplemockBook(id string, name string) domain.Book {
	book := domain.NewBook(id, name)
	return book
}
