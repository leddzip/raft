package script_folder

import (
	"errors"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"testing"
)

var GetAppManagedFolderTestSource string = path.Join("testdata", "GetAppManagedFolder")
var GetAllCandidatesTestSource string = path.Join("testdata", "GetAllCandidates")

const AppScriptFolder string = "raft"

func TestGetAppManagedFolder_whenFolderIsManagedByApp(t *testing.T) {
	// Given
	const WithAppManagedFolder string = "with-app-managed-folder"
	scenarioTestSource := path.Join(GetAppManagedFolderTestSource, WithAppManagedFolder)
	from := path.Join(scenarioTestSource, "first-level", "second-level", "third-level")

	// When
	scriptPath, err := GetAppManagedFolder(AppScriptFolder, from, scenarioTestSource)
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
	const WithoutAppManagedFolder string = "without-app-managed-folder"
	scenarioTestSource := path.Join(GetAppManagedFolderTestSource, WithoutAppManagedFolder)
	from := path.Join(scenarioTestSource, "first-level", "second-level", "third-level")

	// When
	scriptPath, err := GetAppManagedFolder(AppScriptFolder, from, scenarioTestSource)

	// Then
	if scriptPath != "" {
		t.Errorf("Script folder retrieved while none was expected. Got: %s", scriptPath)
	}

	if err == nil {
		t.Error("Got a nil error while one was expected")
	}

	expectedError := NewNotAnAppManagedFolderErrorf("no folder '%s' found in the folder hierarchy. You are not in a folder managed by this tool", AppScriptFolder)
	if err.Error() != expectedError.Error() {
		t.Errorf("Incorrect expected error text. Expected: '%s', got '%s'", expectedError.Error(), err.Error())
	}
	if !errors.As(err, &expectedError) {
		t.Errorf("Expected: %#v, got %#v", expectedError, err)
	}

}

func TestGetAppManagedFolder_whenSearchedPathIsFindButNotAsAFolder(t *testing.T) {
	// Given
	const WithWantedPathButNotADir string = "with-wanted-path-but-not-a-dir"
	scenarioTestSource := path.Join(GetAppManagedFolderTestSource, WithWantedPathButNotADir)
	from := path.Join(scenarioTestSource, "first-level", "second-level", "third-level")

	// When
	scriptPath, err := GetAppManagedFolder(AppScriptFolder, from, scenarioTestSource)

	// Then
	if scriptPath != "" {
		t.Errorf("Script folder retrieved while none was expected. Got: %s", scriptPath)
	}

	if err == nil {
		t.Error("Got a nil error while one was expected")
	}

	expectedError := NewNotAnAppManagedFolderErrorf("path '%s' is not a folder", AppScriptFolder)
	if err.Error() != expectedError.Error() {
		t.Errorf("Incorrect expected error text. Expected: '%s', got '%s'", expectedError.Error(), err.Error())
	}
	if !errors.As(err, &expectedError) {
		t.Errorf("Expected: %#v, got %#v", expectedError, err)
	}

}

func TestSliceContainsString_whenSliceContainsString(t *testing.T) {
	// Given
	s := []string{"one", "two", "three"}

	// When
	result := sliceContainsString(&s, "two")

	// Then
	if result != true {
		t.Errorf("Invalid result. Expected '%t', got '%t'", true, result)
	}
}

func TestSliceContainsString_whenSliceIsEmpty(t *testing.T) {
	// Given
	s := []string{}

	// When
	result := sliceContainsString(&s, "anything")

	// Then
	if result != false {
		t.Errorf("Invalid result. Expected '%t', got '%t'", false, result)
	}
}

func TestSliceContainsString_whenSliceDoesNotContainsString(t *testing.T) {
	// Given
	s := []string{"one", "two", "three"}

	// When
	result := sliceContainsString(&s, "four")

	// Then
	if result != false {
		t.Errorf("Invalid result. Expected '%t', got '%t'", false, result)
	}
}

func TestIsValidCandidateExtension_whenIsValid(t *testing.T) {
	// Given
	candidateNames := []string{
		"this_is_a_test.yaml",
		"this_is_a_test.yml",
		"this.is.a.test.yaml",
		"this.is.a.test.yml",
	}

	for _, candidate := range candidateNames {
		t.Run("assert that each candidate is valid", func(t *testing.T) {
			// When
			result := isValidCandidateExtension(candidate)

			// Then
			if result != true {
				t.Errorf("invalid result for candidate '%s'. Expected '%t', got '%t'", candidate, true, result)
			}
		})
	}
}

func TestIsValidCandidateExtension_whenIsNotValid(t *testing.T) {
	// Given
	candidateNames := []string{
		"this_is_a_test.Yaml",
		"this_is_a_test",
		"this_is_a_test.Yml",
		"this.is.a.test.yamle",
		"this.is.a.test",
	}

	for _, candidate := range candidateNames {
		fmt.Printf("candidate: %s, ext: %s", candidate, filepath.Ext(candidate))
		t.Run("assert that each candidate is not valid", func(t *testing.T) {
			// When
			result := isValidCandidateExtension(candidate)

			// Then
			if result != false {
				t.Errorf("invalid result for candidate '%s'. Expected '%t', got '%t'", candidate, false, result)
			}
		})
	}
}

func TestGetAllCandidates_whenOnlyActualCandidatesExists(t *testing.T) {
	// Given
	const WithOnlyActualCandidates string = "with-only-actual-candidates"
	scenarioTestSource := path.Join(GetAllCandidatesTestSource, WithOnlyActualCandidates)

	// When
	candidates, err := GetAllCandidates(scenarioTestSource)

	// Then
	if err != nil {
		t.Errorf("Unexpected error occured: %s", err.Error())
	}

	if len(candidates) != 3 {
		t.Errorf("Invalid number of retrieved candidates. Expected %d, got %d", 3, len(candidates))
	}
}

func TestGetAllCandidates_whenMixOfCandidatesAndOthers(t *testing.T) {
	// Given
	const WithMixOfCandidatesAndOthers string = "with-mix-of-candidates-and-others"
	scenarioTestSource := path.Join(GetAllCandidatesTestSource, WithMixOfCandidatesAndOthers)

	// When
	candidates, err := GetAllCandidates(scenarioTestSource)

	// Then
	if err != nil {
		t.Errorf("Unexpected error occured: %s", err.Error())
	}

	if len(candidates) != 3 {
		t.Errorf("Invalid number of retrieved candidates. Expected %d, got %d", 3, len(candidates))
	}
}

func TestGetAllCandidates_whenNoCandidateExists(t *testing.T) {
	// Given
	const WithoutCandidate string = "without-any-candidate"
	scenarioTestSource := path.Join(GetAllCandidatesTestSource, WithoutCandidate)

	// When
	candidates, err := GetAllCandidates(scenarioTestSource)

	// Then
	if err != nil {
		t.Errorf("Unexpected error orccured: %s", err.Error())
	}

	if len(candidates) != 0 {
		t.Errorf("Invalid number of retrieved candidates. Expected %d, got %d", 0, len(candidates))
	}
}
