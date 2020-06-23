package store

import (
	"encoding/json"
	"quickship/structs"
)

func (s *Store) fetchStore() (file, error) {
	data, err := read(s.path)
	if err != nil {
		return nil, err
	}

	var f file
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, err
	}

	s.file = &f
	return f, nil
}

// GetList - get list of records
func (s *Store) GetList() []ListRecord {
	retme := make([]ListRecord, 0)
	if s.file != nil && len(*s.file) > 0 {
		for uid, rec := range *s.file {
			retme = append(retme, ListRecord{ID: uid, Name: rec.Name})
		}
	}

	return retme
}

// GetRecordDeploy - get record deploy info
func (s *Store) GetRecordDeploy(id string) structs.Deploy {
	return (*s.file)[id].Deploy
}
