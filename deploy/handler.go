package deploy

import (
	"fmt"
	"log"
	"quickship/structs"
)

func handleServer(d structs.Deploy, server, id string) {
	runpull(d.Branch, server, id)
	if d.Clean != nil {
		runcmd(*d.Clean, server, id)
	}
	if d.Build != nil {
		runcmd(*d.Build, server, id)
	}
	runcmd(d.Run, server, id)
}

func handleReDeployServer(d structs.Deploy, server, id string) {
	runpull(d.Branch, server, id)
	if d.Clean != nil {
		runcmd(*d.Clean, server, id)
	}
	if d.AddFiles != nil && len(*d.AddFiles) > 0 {
		for _, f := range *d.AddFiles {
			runcmd(fmt.Sprintf("curl %s -o %s", f.URL, f.Name), server, id)
		}
	}
	if d.Build != nil {
		runcmd(*d.Build, server, id)
	}
	runcmd(d.Run, server, id)
}

func handleNewServer(d structs.Deploy, id, server string) {
	log.Printf("Setting up : %s \n", id)
	runclone(d.GitRepo, d.Branch, server, id)
	if d.AddFiles != nil && len(*d.AddFiles) > 0 {
		for _, f := range *d.AddFiles {
			runcmd(fmt.Sprintf("curl %s -o %s", f.URL, f.Name), server, id)
		}
	}
}

func handleDelServer(d structs.Deploy, id, server string) {
	if d.Clean != nil {
		runcmd(*d.Clean, server, id)
	}
	runrm(server, id)
}
