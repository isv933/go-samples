package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/isv933/go-samples/url-shortener/gen/api"
)

type shortenerApi struct {
	service ShortenerService
}

func (shortener shortenerApi) CreateShortUrl(ctx context.Context, params api.CreateShortUrlParams) (api.CreateShortUrlOK, error) {
	slog.Info("CreateShortUrl", "params", params)

	shortUrl, err := shortener.service.AddShortUrl(ctx, params.URL)

	if err != nil {
		slog.Error("CreateShortUrl", "error", err)
		return api.CreateShortUrlOK{}, err
	}

	return api.CreateShortUrlOK{Data: strings.NewReader(shortUrl)}, nil
}

func (shortener shortenerApi) DeleteShortUrl(ctx context.Context, params api.DeleteShortUrlParams) error {

	return nil
}

func (shortener shortenerApi) GetFullUrl(ctx context.Context, params api.GetFullUrlParams) (api.GetFullUrlOK, error) {
	slog.Info("GetFullUrl", "params", params)

	fullUrl, err := shortener.service.getFullUrl(ctx, params.ID)
	if err != nil {
		slog.Error("GetFullUrl", "error", err)
		return api.GetFullUrlOK{}, err
	}

	return api.GetFullUrlOK{Data: strings.NewReader(fullUrl)}, nil
}

func (shortenerApi) NewError(ctx context.Context, err error) *api.ErrorStatusCode {

	if _, ok := errors.AsType[UrlNotFoundError](err); ok {
		return &api.ErrorStatusCode{StatusCode: http.StatusNotFound,
			Response: api.Error{Message: "URL not found"}}
	}

	return &api.ErrorStatusCode{StatusCode: http.StatusInternalServerError,
		Response: api.Error{Message: "Internal server error"}}
}
