package main

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pashagolub/pgxmock/v5"
	"github.com/stretchr/testify/require"
)

func TestPostgresRepository_ConvertError(t *testing.T) {
	repository := PostgresRepository{}

	require.ErrorAs(t, repository.convertError(pgx.ErrNoRows), new(UrlNotFoundError))
	require.ErrorIs(t, repository.convertError(pgx.ErrNoRows), pgx.ErrNoRows)

	pgErr := &pgconn.PgError{Code: "42P01"}
	require.ErrorIs(t, repository.convertError(pgErr), pgErr)
}

func TestPostgresRepository_GetFullUrl(t *testing.T) {
	mock, err := pgxmock.NewPool(
		pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
	)

	require.NoError(t, err)
	defer mock.Close()

	const myId = "super_id"
	const url = "http://example.com"
	const query = "SELECT url FROM shortener_url WHERE id = $1"
	mock.ExpectQuery(query).
		WithArgs(myId).
		WillReturnRows(pgxmock.NewRows([]string{"url"}).AddRow(url)).
		Times(1)
	repository := PostgresRepository{pool: mock}
	full, err := repository.GetFullUrl(context.Background(), myId)
	require.NoError(t, err)
	require.Equal(t, url, full)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresRepository_AddShortUrlCollisionLimit(t *testing.T) {
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
	require.Empty(t, id)
	require.ErrorAs(t, err, &UniqueIdConflictError{})
	require.NoError(t, mock.ExpectationsWereMet())
}
