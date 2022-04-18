package task_runner

import "fmt"

// UnrecognizedTaskTypeError used when parsing a YamlTask. When there is
// no suitable implementation for the given YamlTask, this error is returned.
// This mean that the given type is not yet supported, or this is not a valid type.
type UnrecognizedTaskTypeError struct {
	Message string
}

func (err *UnrecognizedTaskTypeError) Error() string {
	return err.Message
}

func NewUnrecognizedTaskTypeError(message string) error {
	return &UnrecognizedTaskTypeError{Message: message}
}

func NewUnrecognizedTaskTypeErrorf(format string, v ...any) error {
	return &UnrecognizedTaskTypeError{
		Message: fmt.Sprintf(format, v...),
	}
}

// UnrecognizedJobTypeError used when parsing a YamlJob. When there is
// no suitable implementation for the given YamlJob, this is error is returned.
// This mean that the given type is not yet supported, or this is not a valid type.
type UnrecognizedJobTypeError struct {
	Message string
}

func (err *UnrecognizedJobTypeError) Error() string {
	return err.Message
}

func NewUnrecognizedJobTypeError(message string) error {
	return &UnrecognizedJobTypeError{Message: message}
}

func NewUnrecognizedJobTypeErrorf(format string, v ...any) error {
	return &UnrecognizedJobTypeError{
		Message: fmt.Sprintf(format, v...),
	}
}
