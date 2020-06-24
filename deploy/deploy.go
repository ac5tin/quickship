package deploy

import (
	"fmt"
	"log"
	"quickship/structs"
	"sync"
)

// Record - deploys or redeploys a record
func Record(d structs.Deploy, id string) {
	log.Println(fmt.Sprintf("Deploying : %s", id))
	for _, server := range d.Nodes {
		go func(server string) {
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
			log.Println(d.AddFiles)
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
	for _, server := range d.Nodes {
		go func(server string) {
			if d.Clean != nil {
				runcmd(*d.Clean, server, id)
			}
			runrm(server, id)
		}(server)
	}
}
