package domain

type Job struct {
	Version     int
	Name        string
	Description string
	Tasks       []Task
}
