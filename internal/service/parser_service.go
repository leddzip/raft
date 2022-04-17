package service

import (
	"errors"
	"github.com/leddzip/raft/internal/domain"
	"gopkg.in/yaml.v2"
	"os"
)

type YamlJob struct {
	Version     int        `yaml:"version"`
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Tasks       []YamlTask `yaml:"tasks"`
}

type YamlTask struct {
	Type    string `yaml:"type"`
	Name    string `yaml:"name"`
	Runner  string `yaml:"runner"`
	Command string `yaml:"command"`
}

func (task *YamlTask) isSimpleRunner() bool {
	return task.Type == "inline"
}

func (task *YamlTask) ToDomainTask() (domain.Task, error) {
	switch {
	case task.isSimpleRunner():
		newTask := domain.NewInlineTask(*task)
		if err := newTask.Validate(); err != nil {
			return nil, err
		}
		return newTask, nil
	default:
		return nil, errors.New("")
	}
}

func Parse(filePath string) (*YamlJob, error) {

	yamlContent, err := os.ReadFile(filePath)
	if err != nil {
		return &YamlJob{}, err
	}

	var yamlJob YamlJob
	err = yaml.Unmarshal(yamlContent, &yamlJob)
	if err != nil {
		return &YamlJob{}, err
	}
	return &yamlJob, nil
}
