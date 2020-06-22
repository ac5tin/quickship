package store

import "quickship/deploy"

type record struct {
	Name   string        `json:"name"`
	Deploy deploy.Deploy `json:"deploy"`
	HookID *int          `json:"hook_id"`
}

type file map[string]*record

// Store - a store object contianing path and file
type Store struct {
	path string
	file *file
}

// ListRecord - a deployment list object (for display)
type ListRecord struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
