package main

import (
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestShortenerAPI(t *testing.T) {
	tests := []struct {
		name          string
		err           error
		httpErrorCode int
	}{
		{"http not found", UrlNotFoundError{pgx.ErrNoRows}, 404},
		{"http unavailable", DatabaseError{pgx.ErrTooManyRows}, 500},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := (shortenerApi{}).NewError(t.Context(), test.err)
			if response.StatusCode != test.httpErrorCode {
				t.Fatalf("got status code %d, want %d", response.StatusCode, test.httpErrorCode)
			}
		})
	}
}
