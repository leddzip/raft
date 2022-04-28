package task_runner

import "testing"

func TestIsSequentialJob_nilType_success(t *testing.T) {
	// Given
	yamlJob := &YamlJob{
		Version:     0,
		Name:        "name",
		Description: "description",
		Type:        nil,
		Tasks:       nil,
	}

	// When
	result := isSequentialJob(yamlJob)

	// Then
	if result != true {
		t.Error("result is false, expected true")
	}
}

func TestIsSequentialJob_validType_success(t *testing.T) {
	// Given
	seq := "sequential"
	yamlJob := &YamlJob{
		Version:     0,
		Name:        "name",
		Description: "description",
		Type:        &seq,
		Tasks:       nil,
	}

	// When
	result := isSequentialJob(yamlJob)

	// Then
	if result != true {
		t.Error("result is false, expected true")
	}
}

func TestIsSequentialJob_invalidType_failure(t *testing.T) {
	// Given
	seq := "non_sequential"
	yamlJob := &YamlJob{
		Version:     0,
		Name:        "name",
		Description: "description",
		Type:        &seq,
		Tasks:       nil,
	}

	// When
	result := isSequentialJob(yamlJob)

	// Then
	if result != false {
		t.Error("result is true, expected false")
	}
}

func TestNewSequentialJob_with_2_Tasks(t *testing.T) {
	// Given
	seq := "sequential"
	yamlJob := YamlJob{
		Version:     0,
		Name:        "name",
		Description: "description",
		Type:        &seq,
		Tasks:       nil,
	}

	for i := 0; i < 2; i++ {
		command := "command placeholder"
		shell := "shell placeholder"
		yamlTask := YamlTask{
			Type:        "shell-runner",
			Name:        "name",
			Description: "description",
			Command:     &command,
			Shell:       &shell,
		}
		yamlJob.Tasks = append(yamlJob.Tasks, yamlTask)
	}

	// When
	jobResult, err := NewSequentialJob(yamlJob)

	// Then
	if err != nil {
		t.Errorf("Unexpected error occurred: %s", err.Error())
	}

	if jobResult.Version() != yamlJob.Version {
		t.Errorf("incorrect version. Expected '%d', got '%d'", yamlJob.Version, jobResult.Version())
	}

	if jobResult.Name() != yamlJob.Name {
		t.Errorf("incorrect name. Expected '%s', got '%s'", yamlJob.Name, jobResult.Name())
	}

	if jobResult.Description() != yamlJob.Description {
		t.Errorf("incorrect description. Expected '%s', got '%s'", yamlJob.Description, jobResult.Description())
	}

	if len(jobResult.Tasks()) != 2 {
		t.Errorf("incorrect task size. Expected '%d', got '%d'", len(yamlJob.Tasks), len(jobResult.Tasks()))
	}

	for _, task := range jobResult.Tasks() {
		t.Run("assert task is correctly mapped to InlineShellRunnerTask", func(t *testing.T) {
			ans := task.Type()
			if ans != Shell {
				t.Errorf("incorrect mapping of yamltask to InlineShellRunnerTask. Expected type '%s', got '%s'", Shell, ans)
			}
		})
	}

}

func TestNewSequentialJob_withoutTasks(t *testing.T) {
	// Given
	seq := "sequential"
	yamlJob := YamlJob{
		Version:     0,
		Name:        "name",
		Description: "description",
		Type:        &seq,
		Tasks:       nil,
	}

	// When
	jobResult, err := NewSequentialJob(yamlJob)

	// Then
	if err != nil {
		t.Errorf("Unexpected error occurred: %s", err.Error())
	}

	if jobResult.Version() != yamlJob.Version {
		t.Errorf("incorrect version. Expected '%d', got '%d'", yamlJob.Version, jobResult.Version())
	}

	if jobResult.Name() != yamlJob.Name {
		t.Errorf("incorrect name. Expected '%s', got '%s'", yamlJob.Name, jobResult.Name())
	}

	if jobResult.Description() != yamlJob.Description {
		t.Errorf("incorrect description. Expected '%s', got '%s'", yamlJob.Description, jobResult.Description())
	}

	if len(jobResult.Tasks()) != 0 {
		t.Errorf("incorrect task size. Expected '%d', got '%d'", len(yamlJob.Tasks), len(jobResult.Tasks()))
	}

}

func TestNewSequentialJob_with_invalid_Tasks(t *testing.T) {
	// Given
	seq := "sequential"
	yamlJob := YamlJob{
		Version:     0,
		Name:        "name",
		Description: "description",
		Type:        &seq,
		Tasks:       nil,
	}

	command := "command placeholder"
	shell := "shell placeholder"
	yamlTask := YamlTask{
		Type:        "invalid-type-that-trigger-an-error",
		Name:        "name",
		Description: "description",
		Command:     &command,
		Shell:       &shell,
	}
	yamlJob.Tasks = append(yamlJob.Tasks, yamlTask)

	// When
	jobResult, err := NewSequentialJob(yamlJob)

	// Then
	if err == nil {
		t.Error("No error detected while one was expected")
	}

	if jobResult != nil {
		t.Errorf("SequentialJov should be nil since an error was expected. Got %#v", jobResult)
	}
}
