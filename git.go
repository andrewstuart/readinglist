package main

import (
	"os"
	"os/exec"
)

func runGit(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)

	//Set working directory to home for git
	cmd.Dir = rlHome

	output, err := cmd.Output()
	if err != nil {
		return err
	}

	os.Stdout.Write(output)

	return nil
}
