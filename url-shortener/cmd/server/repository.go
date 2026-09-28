package main

import "context"

type Repository interface {
	AddShortUrl(ctx context.Context, url string) (string, error)
	GetFullUrl(ctx context.Context, id string) (string, error)
}
