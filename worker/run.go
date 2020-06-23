package worker

import "os/exec"

// Run - run command
func Run(command string) error {
	cmd := exec.Command(command)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
