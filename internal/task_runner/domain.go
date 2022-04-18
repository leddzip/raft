package task_runner

// Job is the representation of what the user want to execute.
// it consists of one or multiple Task which are the fundamental execution unit.
// A Job is mostly a bag that group Tasks together and hold basic information
// such as the purpose of the Job.
type Job interface {
	Version() int
	Name() string
	Description() string
	Tasks() []Task
}

// Task is the fundamental execution unit. Whatever the actual runner (shell, docker
// or something else).
type Task interface {
	Type() RunnerType
	Execute()
	Validate() error
}

// RunnerType describe what kind of runner the Task will use
type RunnerType string

const (
	Shell  RunnerType = "shell-runner"
	Docker            = "docker"
)
