package errors

import (
	"fmt"
)

type NotARaftManagedFolderError struct {
	Message string
}

func (err *NotARaftManagedFolderError) Error() string {
	return err.Message
}

func NewNotARaftManagedFolderError(message string) error {
	return &NotARaftManagedFolderError{Message: message}
}

func NewNotARaftManagedFolderErrorf(format string, v ...any) error {
	return &NotARaftManagedFolderError{
		Message: fmt.Sprintf(format, v...),
	}
}

func NotARaftManagedFolderErrorPointer() *NotARaftManagedFolderError {
	return &NotARaftManagedFolderError{}
}
