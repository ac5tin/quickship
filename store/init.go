package store

import (
	"log"
	"quickship/files"
)

// Init - initialise a store
func Init(path string) Store {

	if !files.FileExist(path) {
		// create new file if does not exist
		if err := files.CreateFile(path); err != nil {
			log.Println("Failed to initialise store")
			log.Panic(err.Error())
		}
		newf := make(file)
		writeStore(newf, path)

	}
	s := Store{
		path: path,
	}
	if _, err := s.fetchStore(); err != nil {
		log.Println("Failed to fetch store")
		log.Panic(err.Error())
	}
	return s
}
