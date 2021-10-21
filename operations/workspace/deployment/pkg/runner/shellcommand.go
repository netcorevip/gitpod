package runner

import (
	"bytes"
	"os/exec"
)

// ShellRun is a simple wrapper to run commands in shell
func ShellRun(shellCmd string, shellArgs []string) (string, string, error) {
	cmd := exec.Command(shellCmd, shellArgs...)
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		return stdOut.String(), stdErr.String(), err
	}
	return stdOut.String(), stdErr.String(), err
}
