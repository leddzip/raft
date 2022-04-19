package script_folder

import "fmt"

// UnexpectedFolderError is a generic error return when there is an issue
// related to the manipulation (read, write, ...) of any app related folder (
// like permission issues for example).
type UnexpectedFolderError struct {
	Message string
}

func (err *UnexpectedFolderError) Error() string {
	return err.Message
}

func NewUnexpectedFolderError(message string) error {
	return &UnexpectedFolderError{Message: message}
}

func NewUnexpectedFolderErrorf(format string, v ...any) error {
	return &UnexpectedFolderError{
		Message: fmt.Sprintf(format, v...),
	}
}

// NotAnAppManagedFolderError is an error used to tell the user he is doing
// a task not in a folder (or sub folder) the app is aware of.
type NotAnAppManagedFolderError struct {
	Message string
}

func (err *NotAnAppManagedFolderError) Error() string {
	return err.Message
}

func NewNotAnAppManagedFolderError(message string) error {
	return &NotAnAppManagedFolderError{Message: message}
}

func NewNotAnAppManagedFolderErrorf(format string, v ...any) error {
	return &NotAnAppManagedFolderError{
		Message: fmt.Sprintf(format, v...),
	}
}
