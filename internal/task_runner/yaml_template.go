package task_runner

type YamlJob struct {
	Version     int        `yaml:"version"`
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Type        *string    `yaml:"type,omitempty"`
	Tasks       []YamlTask `yaml:"tasks"`
}

type YamlTask struct {
	Type        *string `yaml:"type,omitempty"`
	Name        string  `yaml:"name"`
	Description string  `yaml:"runner"`
	Command     *string `yaml:"command,omitempty"`
	Shell       *string `yaml:"shell,omitempty"`
}

func (yamlTask *YamlTask) toTask() (Task, error) {
	switch {
	case isInlineShellRunner(yamlTask):
		newTask := NewInlineShellRunnerTask(yamlTask)
		return newTask, nil
	default:
		return nil, NewUnrecognizedTaskTypeErrorf("the given task can't be map to a valid task implementation: %#v", yamlTask)
	}
}

func (yamlJob *YamlJob) toJob() (Job, error) {

	switch {
	case isSequentialJob(yamlJob):
		return NewSequentialJob(*yamlJob)
	default:
		return nil, NewUnrecognizedJobTypeErrorf("Unrecognized job type: %#v", yamlJob)
	}

}
