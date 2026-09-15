package main

import "context"

type apiHandler struct{}

func (api *apiHandler) GetUserById(ctx context.Context, params GetUserByIdParams) (*User, error) {

}
