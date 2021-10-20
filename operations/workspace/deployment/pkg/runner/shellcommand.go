package runner

import (
	"os"
	"os/exec"
)

func ShellRun(shellCmd string, shellArgs []string) error {
	cmd := exec.Command(shellCmd, shellArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
