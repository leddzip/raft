package task_runner

type InlineShellRunnerTask struct {
	Name        string
	ShellRunner string
	Command     string
}

// isInlineShellRunner verify if the given YamlTask is a InlineShellRunnerTask candidate.
// To be a candidate, the YamlTask should have its type to "shell-runner" and a non nil Command.
func isInlineShellRunner(task *YamlTask) bool {
	return task.Type == "shell-runner" && task.Command != nil
}

// NewInlineShellRunnerTask return a new InlineShellRunnerTask, which is an instance of Task.
// Before using this function, it is advised to assert the YamlTask is a valid InlineShellRunnerTask
// candidate by using the function 'isInlineShellRunner' (to make sure expected field such as 'Command'
// are present and non nil in the YamlTask.
func NewInlineShellRunnerTask(task *YamlTask) *InlineShellRunnerTask {
	var shell string
	if task.Shell == nil {
		shell = "/usr/bin/bash"
	} else {
		shell = *task.Shell
	}
	return &InlineShellRunnerTask{
		Name:        task.Name,
		ShellRunner: shell,
		Command:     *task.Command,
	}
}

func (task *InlineShellRunnerTask) Execute() {

}

func (task *InlineShellRunnerTask) Validate() error {
	return nil
}

func (task *InlineShellRunnerTask) Type() RunnerType {
	return Shell
}
