package store

// low level read write

import (
	"io/ioutil"
	"os"
)

func write(data []byte, path string) error {
	err := os.Truncate(path, 0)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	f.Write(data)
	return nil
}

func read(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
