package main

import (
	"context"
	"fmt"
)

type ShortenerService interface {
	AddShortUrl(ctx context.Context, url string) (string, error)
	getFullUrl(ctx context.Context, id string) (string, error)
}

type ShortenerServiceImpl struct {
	repository  Repository
	redirectUrl string
}

func (s ShortenerServiceImpl) AddShortUrl(ctx context.Context, url string) (string, error) {
	id, err := s.repository.AddShortUrl(ctx, url)

	if err != nil {
		return "", DatabaseError{err}
	}

	return fmt.Sprintf("%s/%s", s.redirectUrl, id), nil
}

func (s ShortenerServiceImpl) getFullUrl(ctx context.Context, id string) (string, error) {
	url, err := s.repository.GetFullUrl(ctx, id)
	if err != nil {
		return "", err
	}

	return url, nil
}

func NewShortenerService(repository Repository, redirectUrl string) (ShortenerService, error) {
	return ShortenerServiceImpl{repository, redirectUrl}, nil
}
