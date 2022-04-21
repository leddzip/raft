package script_folder

import (
	"errors"
	"log"
	"path"
	"testing"
)

const AppScriptFolder string = "raft"
const TestSource string = "testdata"
const WithAppManagedFolder string = "with-app-managed-folder"
const WithWantedPathButNotADir string = "with-wanted-path-but-not-a-dir"
const WithoutAppManagedFolder string = "without-app-managed-folder"

func TestGetAppManagedFolder_whenFolderIsManagedByApp(t *testing.T) {
	// Given
	scenarioTestSource := path.Join(TestSource, WithAppManagedFolder)
	from := path.Join(scenarioTestSource, "first-level", "second-level", "third-level")

	// When
	scriptPath, err := GetAppManagedFolder("raft", from, scenarioTestSource)
	log.Printf("Got path: %s", scriptPath)

	// Then
	if scriptPath == "" {
		t.Error("Error while retrieving the path. Should not be empty")
	}

	if err != nil {
		t.Errorf("Got a non nil error while none was expected: %s", err.Error())
	}
}

func TestGetAppManagedFolder_whenFolderIsNotManagedByApp(t *testing.T) {
	// Given
	scenarioTestSource := path.Join(TestSource, WithoutAppManagedFolder)
	from := path.Join(scenarioTestSource, "first-level", "second-level", "third-level")

	// When
	scriptPath, err := GetAppManagedFolder("raft", from, scenarioTestSource)

	// Then
	if scriptPath != "" {
		t.Errorf("Script folder retrieved while none was expected. Got: %s", scriptPath)
	}

	if err == nil {
		t.Error("Got a nil error while one was expected")
	}

	expectedError := NewNotAnAppManagedFolderError("no folder 'raft' found in the folder hierarchy. You are not in a folder managed by this tool")
	if err.Error() != expectedError.Error() {
		t.Errorf("Incorrect expected error text. Expected: '%s', got '%s'", expectedError.Error(), err.Error())
	}
	if !errors.As(err, &expectedError) {
		t.Errorf("Expected: %#v, got %#v", expectedError, err)
	}

}

func TestGetAppManagedFolder_whenSearchedPathIsFindButNotAsAFolder(t *testing.T) {
	// Given
	scenarioTestSource := path.Join(TestSource, WithWantedPathButNotADir)
	from := path.Join(scenarioTestSource, "first-level", "second-level", "third-level")

	// When
	scriptPath, err := GetAppManagedFolder("raft", from, scenarioTestSource)

	// Then
	if scriptPath != "" {
		t.Errorf("Script folder retrieved while none was expected. Got: %s", scriptPath)
	}

	if err == nil {
		t.Error("Got a nil error while one was expected")
	}

	expectedError := NewNotAnAppManagedFolderError("path 'raft' is not a folder")
	if err.Error() != expectedError.Error() {
		t.Errorf("Incorrect expected error text. Expected: '%s', got '%s'", expectedError.Error(), err.Error())
	}
	if !errors.As(err, &expectedError) {
		t.Errorf("Expected: %#v, got %#v", expectedError, err)
	}

}
