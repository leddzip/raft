package service

import "github.com/leddzip/raft/internal/domain"

func Run(pathTarget string) error {
	job, err := Parse(pathTarget)
	if err != nil {
		return err
	}

	var tasks []domain.Task
	for _, task := range job.Tasks {
		domainTask, err := task.ToDomainTask()
		if err != nil {
			return err
		}

		tasks = append(tasks, domainTask)
	}

	for _, task := range tasks {
		task.Execute()
	}

	return nil
}
