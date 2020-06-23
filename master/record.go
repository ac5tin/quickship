package master

import (
	"quickship/store"
	"quickship/structs"

	uf "github.com/ac5tin/usefulgo"
)

// AddRecord - adds a new deployment record
func AddRecord(d structs.Deploy, name string, s *store.Store) error {
	uuid := uf.GenUUIDV4()
	if err := s.AddRecord(uuid, d, name); err != nil {
		return err
	}
	if err := s.AddRecord(uuid, d, name); err != nil {
		return err
	}
	return nil
}
