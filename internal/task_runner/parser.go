package task_runner

import "gopkg.in/yaml.v2"

// Parse takes a string that contain the content of a yaml file and return a Job and an error.
// If there is an error while parsing (while converting the parsed YamlJob into a Job or while converting
// a YamlTask into a Task), the Job will be nil and the error non nil.
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
