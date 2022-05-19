package main

import (
	"os"
	"os/exec"
)

func getCmdOutput(cmds ...string) ([]byte, error) {
	command := exec.Command(cmds[0], cmds[1:]...)
	command.Stderr = os.Stderr
	output, err := command.Output()
	if err != nil {
		return nil, err
	}

	return output, nil
}
