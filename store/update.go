package store

import (
	"encoding/json"
	"quickship/structs"
)

func (s *Store) saveStore() error {
	data, err := json.Marshal(s.file)
	if err != nil {
		return err
	}
	if err := write(data, s.path); err != nil {
		return err
	}
	if _, err := s.fetchStore(); err != nil {
		return err
	}
	return nil
}

// AddRecord - adds a record to store
func (s *Store) AddRecord(id string, d structs.Deploy, name string) error {
	data, err := read(s.path)
	if err != nil {
		return err
	}
	var f file
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	r := Record{
		Name:   name,
		Deploy: d,
	}
	f[id] = r
	s.file = &f
	if err := s.saveStore(); err != nil {
		return err
	}
	return nil
}

// RmRecord - removes a record by id
func (s *Store) RmRecord(id string) error {
	data, err := read(s.path)
	if err != nil {
		return err
	}
	var f file
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	delete(f, id)
	s.file = &f
	if err := s.saveStore(); err != nil {
		return err
	}
	return nil
}

// SetHookID - sets hook id to a record
func (s *Store) SetHookID(hookid int, recordID string) error {
	data, err := read(s.path)
	if err != nil {
		return err
	}
	var f file
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	if fr, ok := f[recordID]; ok {
		fr.HookID = &hookid
	}

	s.file = &f
	if err := s.saveStore(); err != nil {
		return err
	}
	return nil
}

func writeStore(f file, path string) error {
	data, err := json.Marshal(f)
	if err != nil {
		return err
	}
	if err := write(data, path); err != nil {
		return err
	}
	return nil
}
