package domain

type Task interface {
	Execute()
	Validate() error
}
