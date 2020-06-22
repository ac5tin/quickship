package worker

import "os/exec"

func run(command string) error {
	cmd := exec.Command(command)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
