package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	type args struct {
		id   string
		name string
	}
	tests := []struct {
		name string
		args args
		want Book
	}{
		{"Prueba 1", args{"1", "name"}, NewBook("1", "name")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := NewBook(tt.args.id, tt.args.name)
			assert.Equal(t, tt.want, res)

		})
	}
}
