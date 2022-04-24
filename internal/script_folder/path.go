package script_folder

import (
	backfinder "github.com/leddzip/back-finder"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetAppManagedFolder(folderNameToRetrieve string, fromFolder string, backToFolder string) (string, error) {
	filePresence, err := backfinder.FindFileBetween(folderNameToRetrieve, fromFolder, backToFolder)
	if err != nil {
		return "", NewUnexpectedFolderErrorf("cannot assert if the folder '%s' exist or not", folderNameToRetrieve)
	}

	if !filePresence.IsFilePresent {
		return "", NewNotAnAppManagedFolderErrorf("no folder '%s' found in the folder hierarchy. You are not in a folder managed by this tool", folderNameToRetrieve)
	}

	fileInfo, err := os.Stat(filePresence.FilePathIfExist)
	if err != nil {
		return "", NewUnexpectedFolderErrorf("cannot open path '%s'", filePresence.FilePathIfExist)
	}

	if fileInfo.IsDir() {
		return filePresence.FilePathIfExist, nil
	} else {
		return "", NewNotAnAppManagedFolderErrorf("path '%s' is not a folder", folderNameToRetrieve)
	}
}

// GetAllCandidates retrieve all the script candidates.
// It is assumed that the 'scriptSource' parameter is
// a valid script source for the application (asserted using
// the GetAppManagedFolder function).
// Return a list of all Candidate
func GetAllCandidates(scriptSource string) ([]Candidate, error) {
	var candidates []Candidate

	dirEntries, err := os.ReadDir(scriptSource)
	if err != nil {
		return nil, NewUnexpectedFolderErrorf("unable to read the content of folder '%s", scriptSource)
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}
		if isValidCandidateExtension(entry.Name()) {
			newCandidate := Candidate{
				Path:     path.Join(scriptSource, entry.Name()),
				FileName: entry.Name(),
				Name:     strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())),
			}
			candidates = append(candidates, newCandidate)
		}
	}

	return candidates, nil
}

// isValidCandidateExtension check if the given file name (with extension),
// is a valid candidate. To be a valid candidate, it should be a yaml file
// (either ending with .yaml or .yml).
func isValidCandidateExtension(candidateName string) bool {
	allowedExtensionCandidate := []string{"yaml", "yml"}
	nameSliced := strings.Split(candidateName, ".")
	return sliceContainsString(&allowedExtensionCandidate, nameSliced[len(nameSliced)-1])
}

// sliceContainsString return true if the slice contains the given string.
// false otherwise.
func sliceContainsString(collection *[]string, value string) bool {
	for _, elem := range *collection {
		if elem == value {
			return true
		}
	}
	return false
}
