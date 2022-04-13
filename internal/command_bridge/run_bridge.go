package command_bridge

import (
	"fmt"
	"github.com/leddzip/raft/internal/service"
	"log"
)

func Run() {

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
	// do something since the first assertion succeeded
	raftScriptFolderLocation, err := service.GetRaftManagedFolder()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Raft script folder location: %s", raftScriptFolderLocation)

}
