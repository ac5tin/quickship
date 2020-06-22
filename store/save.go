package store

import "encoding/json"

func (s Store) saveStore() error {
	data, err := json.Marshal(s.file)
	if err != nil {
		return err
	}
	if err := write(data, s.path); err != nil {
		return err
	}
	return nil
}

func (s Store) setHookID(hookid int, recordID string) error {
	data, err := read(s.path)
	if err != nil {
		return err
	}
	var f file
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	f[recordID].HookID = &hookid
	s.file = &f
	if err := s.saveStore(); err != nil {
		return err
	}
	return nil
}
