package main

type DatabaseError struct {
	error error
}

type UrlNotFoundError struct {
	error error
}

func (e DatabaseError) Error() string {
	return e.error.Error()
}

func (e DatabaseError) Unwrap() error {
	return e.error
}

func (e UrlNotFoundError) Error() string {
	return e.error.Error()
}

func (e UrlNotFoundError) Unwrap() error {
	return e.error
}
