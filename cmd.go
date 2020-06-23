package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	command "quickship/cmd"
	"quickship/files"
	"quickship/structs"
)

func greeter() {
	fmt.Println("=== 🚢Quickship ===")
}

func cmd() {
	greeter()
	log.Println(*name)
	if *up != "" {
		upCmd()
		return
	}
}

func upCmd() {
	fmt.Println("Adding new record")
	// check if file exist
	if !files.FileExist(*up) {
		// file does not exist
		fmt.Println("File does not exist")
		return
	}
	// file exist, now try to parse file
	data, err := ioutil.ReadFile(*up)
	if err != nil {
		fmt.Println("Failed to read file")
		log.Panic(err.Error())
		return
	}
	var dp structs.Deploy
	if err := json.Unmarshal(data, &dp); err != nil {
		fmt.Println("Failed to parse file")
		log.Panic(err.Error())
		return
	}
	if err := command.AddRecord(dp, *name, fmt.Sprintf("http://localhost:%d", *addr)); err != nil {
		fmt.Println("Failed to add record")
		log.Panic(err.Error())
		return
	}
	return

}
