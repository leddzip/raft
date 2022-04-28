package service

import (
	"os"
	"path"
	"regexp"
	"strings"
)

func GetCandidateMap(rootPath string) (map[string]os.DirEntry, error) {
	dirEntries, err := os.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}

	candidates := make(map[string]os.DirEntry)

	for _, entry := range dirEntries {
		if isValidCandidate(entry) {
			trimmedName := strings.TrimSuffix(entry.Name(), path.Ext(entry.Name()))
			candidates[trimmedName] = entry
		}
	}

	return candidates, nil
}

func CheckCandidateExist(candidate string, rootPath string) (bool, error) {
	candidateMap, err := GetCandidateMap(rootPath)
	if err != nil {
		return false, err
	}

	_, isCandidatePresent := candidateMap[candidate]

	return isCandidatePresent, nil
}

func isValidCandidate(candidate os.DirEntry) bool {
	if candidate.IsDir() {
		return false
	}

	match, _ := regexp.MatchString("[a-zA-Z-_.@]*\\.ya?ml", candidate.Name())
	return match
}
