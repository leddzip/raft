package task_runner

import "gopkg.in/yaml.v2"

func Parse(yamlContent string) (Job, error) {
	var yamlJob YamlJob
	err := yaml.Unmarshal([]byte(yamlContent), &yamlJob)
	if err != nil {
		return nil, err
	}

	job, err := yamlJob.toJob()
	if err != nil {
		return nil, err
	}

	return job, nil
}
