package script_folder

import (
	backfinder "github.com/leddzip/back-finder"
	"os"
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
