package master

import (
	"fmt"
	"os"
	"quickship/store"
	"quickship/structs"
	"quickship/worker"
	"strings"

	uf "github.com/ac5tin/usefulgo"
)

// AddRecord - adds a new deployment record and hook
func AddRecord(d structs.Deploy, name string, s *store.Store) (string, error) {
	uuid := uf.GenUUIDV4()

	splits := strings.Split(d.GitRepo, "/")
	hookid, err := worker.CreateHook(splits[len(splits)-2], splits[len(splits)-1], os.Getenv("GITHUB_TOKEN"), fmt.Sprintf("%s/api/master/webhook/%s", os.Getenv("SERVER_ADDRESS"), uuid))
	if err != nil {
		return "", err
	}

	if err := s.AddRecord(uuid, d, name, int(hookid)); err != nil {
		return "", err
	}

	return uuid, nil
}

// rmRecord - removes record and removes hook
func rmRecord(id string, s *store.Store) error {

	record := s.GetRecord(id)
	splits := strings.Split(record.Deploy.GitRepo, "/")
	if err := worker.DeleteHook(splits[len(splits)-2], splits[len(splits)-1], os.Getenv("GITHUB_TOKEN"), *record.HookID); err != nil {
		return err
	}
	if err := s.RmRecord(id); err != nil {
		return err
	}
	return nil
}
