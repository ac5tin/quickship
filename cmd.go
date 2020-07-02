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
	if *down != "" {
		downCmd()
		return
	}
	if *list {
		lCmd()
		return
	}

	if *addnode != "" {
		addNodeCmd()
		return
	}

	if *delnode != "" {
		rmNodeCmd()
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
	if err := command.AddRecord(dp, *name, fmt.Sprintf("%s:%d", *ms, *port)); err != nil {
		fmt.Println("Failed to add record")
		log.Panic(err.Error())
		return
	}
	return

}

func downCmd() {
	fmt.Println("Removing record")
	if err := command.RmRecord(*down, fmt.Sprintf("%s:%d", *ms, *port)); err != nil {
		fmt.Println("Failed to remove record")
		log.Panic(err.Error())
		return
	}
	return
}

// list
func lCmd() {
	fmt.Println("Listing records")
	reclist, err := command.ListRecords(fmt.Sprintf("%s:%d", *ms, *port))
	if err != nil {
		fmt.Println("Failed to retrieve list")
		log.Panic(err.Error())
		return
	}

	// print list
	for _, r := range reclist {
		fmt.Printf("ID: %s | Name: %s \n", r.ID, r.Name)
	}
	return

}

// add node
func addNodeCmd() {
	fmt.Println("Adding Node")
	if err := command.AddNode(fmt.Sprintf("%s:%d", *ms, *port), *addnode, *rid); err != nil {
		fmt.Println("Failed to add node")
		log.Panic(err.Error())
		return
	}
	fmt.Println("Successfully Added node")
}

// remove node
func rmNodeCmd() {
	fmt.Println("Removing Node")
	if err := command.RmNode(fmt.Sprintf("%s:%d", *ms, *port), *delnode, *rid); err != nil {
		fmt.Println("Failed to remove node")
		log.Panic(err.Error())
		return
	}
	fmt.Println("Failed to remove node")
}
