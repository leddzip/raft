package task_runner

type SequentialJob struct {
	version     int
	name        string
	description string
	tasks       []Task
}

// isSequentialJob check if the given YamlJob is a valid SequentialJob candidate.
// To be a valid SequentialJob, a YamlJob need to have its Type to nil or
// this Type should be set to "sequential".
// Because of the nil check, this Job is the default supported job (any job will
// be a SequentialJob if the Type is not provided).
func isSequentialJob(job *YamlJob) bool {
	return job.Type == nil || *job.Type == "sequential"
}

// Version gets the version of the Job
func (job *SequentialJob) Version() int {
	return job.version
}

// Name gets the name of the Job
func (job *SequentialJob) Name() string {
	return job.name
}

// Description get the description of the Job
func (job *SequentialJob) Description() string {
	return job.description
}

// Tasks gets the tasks of the Job
func (job *SequentialJob) Tasks() []Task {
	return job.tasks
}

func (job *SequentialJob) Run() error {
	for _, task := range job.tasks {
		if err := task.Execute(); err != nil {
			return err
		}
	}

	return nil
}

// NewSequentialJob create a new SequentialJob based on a YamlJob.
// It takes care of the mapping of its tasks implicitly (as long as the
// YamlTask can be mapped to a valid Task using the '(yamlTask *YamlTask) toTask()'
// method.
// Return a pointer to SequentialJob and an error. The error is non nil if there
// is an issue while mapping the inner YamlTask to a valid Task.
func NewSequentialJob(yamlJob YamlJob) (*SequentialJob, error) {
	job := &SequentialJob{
		version:     yamlJob.Version,
		name:        yamlJob.Name,
		description: yamlJob.Description,
		tasks:       []Task{},
	}
	for _, yamlTask := range yamlJob.Tasks {
		task, err := yamlTask.toTask()
		if err != nil {
			return nil, err
		}
		job.tasks = append(job.tasks, task)
	}
	return job, nil
}
