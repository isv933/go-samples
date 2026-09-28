package main

import (
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestShortenerApiNotFound(t *testing.T) {
	response := (shortenerApi{}).NewError(t.Context(), UrlNotFoundError{pgx.ErrNoRows})
	if response.StatusCode != 404 {
		t.Fatalf("expected HTTP 404, got %d", response.StatusCode)
	}
}

func TestShortenerApiServiceUnavailable(t *testing.T) {
	response := (shortenerApi{}).NewError(t.Context(), DatabaseError{pgx.ErrTooManyRows})
	if response.StatusCode != 500 {
		t.Fatalf("expected HTTP 500, got %d", response.StatusCode)
	}
}
