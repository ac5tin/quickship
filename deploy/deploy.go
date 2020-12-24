package deploy

import (
	"fmt"
	"log"
	"quickship/store"
	"quickship/structs"
	"sync"
	"time"

	uf "github.com/ac5tin/usefulgo"
)

// Record - deploys or redeploys a record
func Record(d structs.Deploy, id string) {
	log.Println(fmt.Sprintf("Deploying : %s", id))
	for _, server := range d.Nodes {
		go handleServer(d, server, id)
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
			handleNewServer(d, id, server)
		}(server)
	}
	wg.Wait()
	go Record(d, id)
}

// DelRecord - removes a record
func DelRecord(d structs.Deploy, id string) {
	log.Println(fmt.Sprintf("Deleting Record : %s", id))
	for _, server := range d.Nodes {
		go handleDelServer(d, id, server)
	}
}

// ReDeploy - full rebuild,redeploys a record
func ReDeploy(d structs.Deploy, id string) {
	log.Println(fmt.Sprintf("Redeploying Record : %s", id))
	for _, server := range d.Nodes {
		handleReDeployServer(d, server, id)
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
		// check if node still exist
		if !uf.ArrContains(d.Nodes, server) {
			// node not exist
			return false
		}
		// ping
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

// AddNode - adds a new node
func AddNode(server, id string, s *store.Store) {
	log.Println("Adding new node")
	if err := s.AddSNode(server, id); err != nil {
		log.Println("Failed to add new node")
		log.Println(err.Error())
	}
	d := s.GetRecordDeploy(id)
	handleNewServer(d, id, server)
	handleServer(d, server, id)
}

// RemoveNode - removes a node
func RemoveNode(server, id string, s *store.Store) {
	log.Println("Removing node")
	if err := s.DelSNode(server, id); err != nil {
		log.Println("Failed to add new node")
		log.Println(err.Error())
	}
	handleDelServer(s.GetRecordDeploy(id), id, server)
}

// KeepStoreAlive - monitor all records existing in store
func KeepStoreAlive(s *store.Store) {
	reclist := s.GetList()
	for _, rec := range reclist {
		go KeepAlive(rec.ID, s)
	}
}
