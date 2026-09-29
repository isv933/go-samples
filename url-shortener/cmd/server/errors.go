package main

type DatabaseError struct {
	error error
}

type UrlNotFoundError struct {
	error error
}

type UniqueIdConflictError struct {
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

func (e UniqueIdConflictError) Error() string {
	return "could not generate unique id"
}
