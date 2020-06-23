package worker

import (
	"os/exec"
	"strings"
)

// Run - run command
func Run(command string) error {
	splits := strings.Split(command, " ")
	cmd := exec.Command(splits[0], splits[1:]...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
