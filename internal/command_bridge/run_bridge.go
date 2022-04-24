package command_bridge

import (
	"github.com/leddzip/raft/internal/script_folder"
	"github.com/leddzip/raft/internal/service"
	"log"
)

func Run(target string) {

	// ======================================================================
	// Assert we are in a Raft manged folder in order to execute this command
	inRaftManagedFolder, err := service.IsInRaftManagedFolder()
	if err != nil {
		log.Fatal(err.Error())
	}

	if !inRaftManagedFolder {
		log.Fatal("You are not currently inside a raft managed folder.")
	}

	// =====================================================================
	// check if target is a valid candidate
	raftScriptFolderLocation, err := service.GetRaftManagedFolder()
	if err != nil {
		log.Fatal(err.Error())
	}

	isValidCandidate, err := service.CheckCandidateExist(target, raftScriptFolderLocation)
	if err != nil {
		log.Fatal(err.Error())
	}

	if !isValidCandidate {
		log.Fatalf("Candidate '%s' is not a valid target.", target)
	}

	// =====================================================================
	// parse and execute the target
	allTargets, err := service.GetCandidateMap(raftScriptFolderLocation)
	err = service.Run(allTargets[target].Name())
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getCandidate(candidateName string) (*script_folder.CandidateWithContent, error) {
	return nil, nil
}
