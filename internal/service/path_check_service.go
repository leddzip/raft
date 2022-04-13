package service

import (
	"errors"
	backfinder "github.com/leddzip/back-finder"
	"github.com/leddzip/raft/internal/config"
	rafterror "github.com/leddzip/raft/internal/errors"
	"os"
)

func GetRaftManagedFolder() (string, error) {
	filePresence, err := backfinder.FindFileBetween(config.RaftScriptFolder, config.PathBoundaries.From, config.PathBoundaries.BackTo)
	if err != nil {
		return "", rafterror.NewUnexpectedErrorf("cannot assert if folder '%s' exist or not", config.RaftScriptFolder)
	}

	if !filePresence.IsFilePresent {
		return "", rafterror.NewNotARaftManagedFolderErrorf("no folder '%s' found in folder hierarchy. You are not in a raft managed folder", config.RaftScriptFolder)
	}

	fileInfo, err := os.Stat(filePresence.FilePathIfExist)
	if err != nil {
		return "", rafterror.NewUnexpectedErrorf("cannot open path '%s'", filePresence.FilePathIfExist)
	}

	if fileInfo.IsDir() {
		return filePresence.FilePathIfExist, nil
	} else {
		return "", rafterror.NewNotARaftManagedFolderErrorf("path '%s' is not a folder", filePresence.FilePathIfExist)
	}

}

func IsInRaftManagedFolder() (bool, error) {
	_, err := GetRaftManagedFolder()
	unexpectedError := rafterror.NewUnexpectedError("")
	notARaftManagedFolderError := rafterror.NewNotARaftManagedFolderError("")

	switch {
	case errors.As(err, &unexpectedError):
		return false, err
	case errors.As(err, &notARaftManagedFolderError):
		return false, nil
	case err == nil:
		return true, nil
	default:
		return false, rafterror.NewUnexpectedError("unable to determine if inside a raft managed folder")
	}
}
