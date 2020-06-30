package deploy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"quickship/slave"
	"quickship/structs"
)

func runcmd(cmd string, server string, id string) error {
	creq := slave.CmdReq{
		Command: cmd,
	}
	reqBody, err := json.Marshal(creq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/slave/cmd/%s", server, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// REQ SETUP SUCCESS
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func runclone(repo, branch, server, id string) error {
	creq := slave.CloneReq{
		Repo:   repo,
		Branch: branch,
	}

	reqBody, err := json.Marshal(creq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/slave/clone/%s", server, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// REQ SETUP SUCCESS
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func runpull(branch, server, id string) error {
	preq := slave.PullReq{
		Branch: branch,
	}
	reqBody, err := json.Marshal(preq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/slave/pull/%s", server, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// REQ SETUP SUCCESS
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func runrm(server, id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/slave/delete/%s", server, id), nil)
	if err != nil {
		return err
	}
	// REQ SETUP SUCCESS
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func runping(server string, hc structs.HealthCheck) error {
	preq, err := json.Marshal(hc)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/slave/ping", server), bytes.NewBuffer(preq))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// server returned something successfully
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("Ping failed")
	}
	return nil
}
