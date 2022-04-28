package domain

import (
	"bytes"
	"github.com/leddzip/raft/internal/service"
	"log"
	"os/exec"
	"strings"
)

type InlineTask struct {
	Name    string
	Runner  string
	Command string
}

func NewInlineTask(task service.YamlTask) *InlineTask {
	return &InlineTask{
		Name:    task.Name,
		Runner:  task.Runner,
		Command: task.Command,
	}
}

func (inline *InlineTask) Execute() {
	cmd := exec.Command(inline.Runner)
	cmd.Stdin = strings.NewReader(inline.Command)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error while executing task '%s':\n%s", inline.Name, err.Error())
	}
}

func (inline *InlineTask) Validate() error {
	return nil
}
