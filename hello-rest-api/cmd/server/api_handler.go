package main

import (
	"context"
	"strconv"

	"github.com/isv933/go-samples/hello-rest-api/gen/api"
)

type apiHandler struct{}

func (apiHandler) GetUserById(_ context.Context, params api.GetUserByIdParams) (*api.User, error) {
	// Пример данных. Здесь можно подключить хранилище пользователей.
	user := &api.User{Name: "User " + strconv.Itoa(params.ID)}
	if params.ID == 1 {
		user.Name = "Alice"
		user.SetEmail(api.NewOptString("alice@example.com"))
	}
	return user, nil
}

func (apiHandler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {

	return &api.ErrorStatusCode{StatusCode: 500}
}
