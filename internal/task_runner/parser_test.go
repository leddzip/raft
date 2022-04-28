package task_runner

import "testing"

func TestParser_SequentialJob_noTask(t *testing.T) {
	// Given
	yamlContent := `version: 0
name: myName
description: myDescription
type: sequential
`
	// When
	sequentialJob, err := Parse(yamlContent)

	// Then
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if _, isSequentialJob := sequentialJob.(*SequentialJob); !isSequentialJob {
		t.Errorf("Unexpected type for Job. Got %T, expected '%s'", sequentialJob, "SequentialJob")
	}

	expectedVersion := 0
	if sequentialJob.Version() != expectedVersion {
		t.Errorf("Invalid version. Expected %d, got %d", expectedVersion, sequentialJob.Version())
	}

	expectedName := "myName"
	if sequentialJob.Name() != expectedName {
		t.Errorf("Invalid name. Expected '%s', got '%s'", expectedName, sequentialJob.Name())
	}

	expectedDescription := "myDescription"
	if sequentialJob.Description() != expectedDescription {
		t.Errorf("Invalid description. Expected '%s', got '%s'", expectedDescription, sequentialJob.Description())
	}

	if len(sequentialJob.Tasks()) != 0 {
		t.Errorf("Invalid task. Expected no task, got '%d'", len(sequentialJob.Tasks()))
	}
}
