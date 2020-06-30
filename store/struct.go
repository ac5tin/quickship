package store

import "quickship/structs"

// Record - a deployment record
type Record struct {
	Name   string         `json:"name"`
	Deploy structs.Deploy `json:"deploy"`
	HookID *int           `json:"hook_id"`
}

type file map[string]Record

// Store - a store object contianing path and file
type Store struct {
	master string
	path   string
	file   *file
}

// ListRecord - a deployment list object (for display)
type ListRecord struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
