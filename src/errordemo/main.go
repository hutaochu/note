package main

import (
	"errors"
	"fmt"
)

func main() {
	err := testFound()
	var errNotFound NotFoundErr
	if err != nil && errors.As(err, &errNotFound) {
		fmt.Println("is not found error")
	} else {
		fmt.Println("not a not found error")
	}
}

type NotFoundErr interface {
	Error() string
}

type notFoundErr struct {
	message string
}

func (e *notFoundErr) Error() string {
	return e.message

}

func testFound() error {
	err := &notFoundErr{
		message: "not found",
	}
	return fmt.Errorf("wrap error: %w", err)
}
