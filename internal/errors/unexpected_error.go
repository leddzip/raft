package errors

import (
	"fmt"
)

type UnexpectedError struct {
	Message string
}

func (err *UnexpectedError) Error() string {
	return err.Message
}

func NewUnexpectedError(message string) error {
	return &UnexpectedError{Message: message}
}

func NewUnexpectedErrorf(format string, v ...any) error {
	return &UnexpectedError{
		Message: fmt.Sprintf(format, v...),
	}
}
