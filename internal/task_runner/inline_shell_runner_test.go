package task_runner

import (
	"testing"
)

// isInlineShellRunner ============================================================

func TestIsInlineShellRunner_success(t *testing.T) {
	// Given
	command := "command placeholder"
	shell := "shell placeholder"
	yamlTask := &YamlTask{
		Type:        "shell-runner",
		Name:        "name",
		Description: "description",
		Command:     &command,
		Shell:       &shell,
	}

	// When
	result := isInlineShellRunner(yamlTask)

	// Then
	if result != true {
		t.Error("result is false, expected true")
	}
}

func TestIsInlineShellRunner_invalidType(t *testing.T) {
	// Given
	command := "command placeholder"
	shell := "shell placeholder"
	yamlTask := &YamlTask{
		Type:        "NOT A VALID TYPE",
		Name:        "name",
		Description: "description",
		Command:     &command,
		Shell:       &shell,
	}

	// when
	result := isInlineShellRunner(yamlTask)

	// Then
	if result != false {
		t.Error("result is true, expected false")
	}
}

func TestIsInlineShellRunner_nilCommand(t *testing.T) {
	// Given
	shell := "shell placeholder"
	yamlTask := &YamlTask{
		Type:        "NOT A VALID TYPE",
		Name:        "name",
		Description: "description",
		Command:     nil,
		Shell:       &shell,
	}

	// when
	result := isInlineShellRunner(yamlTask)

	// Then
	if result != false {
		t.Error("result is true, expected false")
	}
}

// NewInlineShellRunnerTask =======================================================

func TestNewInlineShellRunnerTask_withShell(t *testing.T) {
	// Given
	command := "command placeholder"
	shell := "my shell"
	yamlTask := &YamlTask{
		Type:        "shell-runner",
		Name:        "name",
		Description: "description",
		Command:     &command,
		Shell:       &shell,
	}

	// When
	task := NewInlineShellRunnerTask(yamlTask)

	// Then
	if task.Name != yamlTask.Name {
		t.Errorf("invalid name. Got '%s', expected '%s'", task.Name, yamlTask.Name)
	}

	if task.ShellRunner != *yamlTask.Shell {
		t.Errorf("invalid shell. Got '%s', expected '%s'", task.ShellRunner, *yamlTask.Shell)
	}

	if task.Command != *yamlTask.Command {
		t.Errorf("invalid command. Got '%s', expected '%s'", task.Command, *yamlTask.Command)
	}
}

func TestNewInlineShellRunnerTask_withoutShell(t *testing.T) {
	// Given
	command := "command placeholder"
	yamlTask := &YamlTask{
		Type:        "shell-runner",
		Name:        "name",
		Description: "description",
		Command:     &command,
		Shell:       nil,
	}

	// When
	task := NewInlineShellRunnerTask(yamlTask)

	// Then
	if task.Name != yamlTask.Name {
		t.Errorf("invalid name. Got '%s', expected '%s'", task.Name, yamlTask.Name)
	}

	if task.ShellRunner != "/usr/bin/bash" {
		t.Errorf("invalid shell. Got '%s', expected '%s'", task.ShellRunner, "/usr/bin/bash")
	}

	if task.Command != *yamlTask.Command {
		t.Errorf("invalid command. Got '%s', expected '%s'", task.Command, *yamlTask.Command)
	}
}
