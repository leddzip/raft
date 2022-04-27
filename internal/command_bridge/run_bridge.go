package command_bridge

import (
	"github.com/leddzip/raft/internal/script_folder"
	"github.com/leddzip/raft/internal/task_runner"
)

func Run(targetName string) error {

	// Assert we are in a folder managed by the application
	scriptFolder, err := script_folder.GetAppManagedFolder("raft", ".", "/")
	if err != nil {
		return err // to replace with an actual error that illustrate an issue while running the application
	}

	// Retrieve candidate
	allCandidates, err := script_folder.GetAllCandidates(scriptFolder)
	if err != nil {
		return err
	}

	runnerCandidate, err := script_folder.GetCandidateWithContentIfExist(allCandidates, targetName)
	if err != nil {
		return err
	}

	// Infer correct Job from candidate
	job, err := task_runner.Parse(runnerCandidate.Content)
	if err != nil {
		return err
	}

	err = job.Run()
	if err != nil {
		return err
	}

	// Run the job
	return nil
}

func getCandidate(candidateName string) (*script_folder.CandidateWithContent, error) {
	return nil, nil
}
