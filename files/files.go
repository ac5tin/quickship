package files

import (
	"os"
)

// FileExist returns bool value of whether file exists
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// CreateFile creates empty file
func CreateFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
