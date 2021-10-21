package runner

import (
	"bytes"
	"os"
	"os/exec"
)

// ShellRun is a simple wrapper to run commands in shell and return the
// Stdout, Stdout as strings
func ShellRun(shellCmd string, shellArgs []string) (string, string, error) {
	cmd := exec.Command(shellCmd, shellArgs...)
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	return stdOut.String(), stdErr.String(), err
}

// ShellRun is a simple wrapper to run commands in shell
// and redirect Stdout and Stderr to OS's Stdout and Stderr
func ShellRunWithDefaultConfig(shellCmd string, shellArgs []string) error {
	cmd := exec.Command(shellCmd, shellArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return err
}
