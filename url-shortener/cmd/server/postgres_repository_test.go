package main

import (
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pashagolub/pgxmock/v5"
)

func TestPostgresRepositoryConvertError(t *testing.T) {
	repository := PostgresRepository{}

	notFound := repository.convertError(pgx.ErrNoRows)
	if _, ok := errors.AsType[UrlNotFoundError](notFound); !ok {
		t.Fatalf("expected UrlNotFoundError, got %T", notFound)
	}
	if !errors.Is(notFound, pgx.ErrNoRows) {
		t.Fatalf("expected pgx.ErrNoRows in error chain, got %v", notFound)
	}

	pgErr := &pgconn.PgError{Code: "42P01"}
	databaseErr := repository.convertError(pgErr)
	if _, ok := errors.AsType[DatabaseError](databaseErr); !ok {
		t.Fatalf("expected DatabaseError, got %T", databaseErr)
	}
	if !errors.Is(databaseErr, pgErr) {
		t.Fatalf("expected PostgreSQL error in error chain, got %v", databaseErr)
	}
}

func TestPostgresRepositoryCollisionLimit(t *testing.T) {
	mock, err := pgxmock.NewPool(
		pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	const url = "https://example.com"
	const insert = "INSERT INTO shortener_url(id, url) VALUES($1, $2) ON CONFLICT (id) DO NOTHING"

	mock.ExpectExec(insert).
		WithArgs(pgxmock.AnyArg(), url).
		WillReturnResult(pgxmock.NewResult("INSERT", 0)).
		Times(3)

	repository := PostgresRepository{pool: mock}
	id, err := repository.AddShortUrl(t.Context(), url)
	if id != "" {
		t.Errorf("id = %q; want empty string", id)
	}
	if !errors.Is(err, errIDAttemptsExhausted) {
		t.Errorf("error = %v; want errIDAttemptsExhausted", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock expectations: %v", err)
	}
}
