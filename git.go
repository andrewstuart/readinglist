package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args[0:]...)

	//Set working directory to home for git
	cmd.Dir = rlHome

	//Run git commands
	output, err := cmd.Output()
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}

func checkGit() error {
	_, err := os.Stat(fmt.Sprintf("%s/.git", rlHome))
	return err
}

func commitGit() error {
	_, err := runGit("add", rlFileName)
	if err != nil {
		return fmt.Errorf("error adding file %s: %v", rlFileName, err)
	}

	_, err = runGit("commit", "-m", "Action at "+time.Now().String())
	if err != nil {
		return fmt.Errorf("error committing: %v", err)
	}

	return nil
}
