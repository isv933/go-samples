package main

import (
	"context"
	"strings"

	"github.com/isv933/go-samples/url-shortener/gen/api"
)

type shortenerApi struct{}

func (shortenerApi) CreateShortUrl(ctx context.Context, params api.CreateShortUrlParams) (api.CreateShortUrlOK, error) {

	return api.CreateShortUrlOK{Data: strings.NewReader("http://shortener/go/")}, nil

}

func (shortenerApi) DeleteShortUrl(ctx context.Context, params api.DeleteShortUrlParams) error {

	return nil
}

func (shortenerApi) GetFullUrl(ctx context.Context, params api.GetFullUrlParams) (api.GetFullUrlOK, error) {

	return api.GetFullUrlOK{Data: strings.NewReader("http://new_url")}, nil

}

func (shortenerApi) NewError(ctx context.Context, err error) *api.ErrorStatusCode {

	return &api.ErrorStatusCode{StatusCode: 500, Response: api.Error{Message: "Bad request"}}
}
