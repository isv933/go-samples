package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var errIDAttemptsExhausted = errors.New("could not generate unique id")

type shortenerDB interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...any) pgx.Row
}

type PostgresRepository struct {
	pool shortenerDB
}

func (s PostgresRepository) AddShortUrl(ctx context.Context, url string) (string, error) {
	for range 3 {
		uniqueId, err := s.getUniqueId()
		if err != nil {
			return "", fmt.Errorf("could not generate unique id: %w", err)
		}
		tag, err := s.pool.Exec(ctx, "INSERT INTO shortener_url(id, url) VALUES($1, $2) ON CONFLICT (id) DO NOTHING", uniqueId, url)

		if err != nil {
			return "", s.convertError(err)
		}
		if tag.RowsAffected() == 1 {
			return uniqueId, nil
		}
	}
	return "", errIDAttemptsExhausted
}

func (s PostgresRepository) GetFullUrl(ctx context.Context, id string) (string, error) {
	var url string
	if err := s.pool.QueryRow(ctx, "SELECT url FROM shortener_url WHERE id = $1", id).Scan(&url); err != nil {
		return "", s.convertError(err)
	}
	return url, nil
}

func (s PostgresRepository) getUniqueId() (string, error) {
	var uniqueId [12]byte

	res, err := rand.Read(uniqueId[:])
	if (err != nil) || (res != len(uniqueId)) {
		return "", s.convertError(err)
	}

	return base64.URLEncoding.EncodeToString(uniqueId[:]), nil
}

func (s PostgresRepository) convertError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return UrlNotFoundError{err}
	}

	return DatabaseError{err}
}
