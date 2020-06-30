package deploy

import (
	"fmt"
	"log"
	"quickship/store"
	"quickship/structs"
	"sync"
	"time"
)

// Record - deploys or redeploys a record
func Record(d structs.Deploy, id string) {
	log.Println(fmt.Sprintf("Deploying : %s", id))
	for _, server := range d.Nodes {
		go func(server string) {
			runpull(d.Branch, server, id)
			if d.Clean != nil {
				runcmd(*d.Clean, server, id)
			}
			if d.Build != nil {
				runcmd(*d.Build, server, id)
			}
			runcmd(d.Run, server, id)

		}(server)
	}
}

// NewRecord - creates a new record
func NewRecord(d structs.Deploy, id string) {
	log.Println(fmt.Sprintf("Adding New Record : %s", id))

	var wg sync.WaitGroup
	wg.Add(len(d.Nodes))
	for _, server := range d.Nodes {
		go func(server string) {
			defer wg.Done()
			runclone(d.GitRepo, d.Branch, server, id)
			if d.AddFiles != nil && len(*d.AddFiles) > 0 {
				for _, f := range *d.AddFiles {
					runcmd(fmt.Sprintf("curl %s -o %s", f.URL, f.Name), server, id)
				}

			}
		}(server)
	}
	wg.Wait()
	go Record(d, id)
}

// DelRecord - removes a record
func DelRecord(d structs.Deploy, id string) {
	log.Println(fmt.Sprintf("Deleting Record : %s", id))
	for _, server := range d.Nodes {
		go func(server string) {
			if d.Clean != nil {
				runcmd(*d.Clean, server, id)
			}
			runrm(server, id)
		}(server)
	}
}

// KeepAlive - making sure deployment is always alive
func KeepAlive(id string, s *store.Store) {
	log.Println(fmt.Sprintf("Monitoring : %s", id))
	d := s.GetRecordDeploy(id)
	for _, server := range d.Nodes {
		go HealthCheck(server, id, s)
	}
}

// HealthCheck - handles health, check health and run cmd if health check fails, returns true if ended cuz failed checks, false if anything else
func HealthCheck(server, id string, s *store.Store) bool {
	var checks uint8 = 0
	for {
		// record doesnt exist anymore ,exit
		if !s.Exist(id) {
			return false
		}
		d := s.GetRecordDeploy(id)
		if err := runping(server, d.Health); err != nil {
			log.Println(err.Error())
			checks++
			log.Printf("Failed %d out of %d checks\n", checks, d.Health.Checks)
		} else {
			log.Println("Health check success")
			checks = 0
		}

		if checks == d.Health.Checks {
			log.Printf("Record : %s (server : %s) Failed all %d checks, now attempting to redeploy ...\n", id, server, d.Health.Checks)
			if d.Health.Run != nil {
				runcmd(*d.Health.Run, server, id)
			}
			Record(d, id)
			// reset
			checks = 0
		}

		time.Sleep(time.Millisecond * time.Duration(d.Health.Interval))
	}
}
