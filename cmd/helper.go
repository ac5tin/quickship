package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"quickship/store"
	"quickship/structs"
	"strings"
)

// AddRecord - adds new record
func AddRecord(d structs.Deploy, name string, server string) error {
	recordReq := store.Record{
		Name:   name,
		Deploy: d,
	}
	reqBody, err := json.Marshal(recordReq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/master/record/add", server), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	// REQ SETUP SUCCESS
	// NOW SEND REQ
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

// RmRecord - remove record
func RmRecord(id, server string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/master/record/%s", server, strings.TrimSpace(id)), nil)
	if err != nil {
		return err
	}
	// REQ SETUP SUCCESS
	// NOW SEND REQ
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

// ListRecords - return list of records
func ListRecords(server string) ([]store.ListRecord, error) {
	req, err := http.NewRequest("Get", fmt.Sprintf("%s/api/master/list/all", server), nil)
	if err != nil {
		return make([]store.ListRecord, 0), err
	}
	// REQ SETUP SUCCESS
	// NOW SEND REQ
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return make([]store.ListRecord, 0), err
	}
	defer resp.Body.Close()
	var listresp struct {
		Data []store.ListRecord `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&listresp); err != nil {
		return make([]store.ListRecord, 0), err
	}
	return listresp.Data, nil
}

// GetRecord - returns a single store record
func GetRecord(server, id string) (store.Record, error) {
	var rec store.Record
	req, err := http.NewRequest("Get", fmt.Sprintf("%s/api/master/record/%s", server, id), nil)
	if err != nil {
		return rec, err
	}
	// REQ SETUP SUCCESS
	// NOW SEND REQ
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return rec, err
	}
	defer resp.Body.Close()
	var recresp struct {
		Record store.Record `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&recresp); err != nil {
		return rec, err
	}
	return recresp.Record, nil
}

// AddNode - adds a node to record
func AddNode(server, node, id string) error {
	areq := struct {
		URL string `json:"url"`
	}{
		URL: node,
	}
	reqBody, err := json.Marshal(areq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/master/record/%s/node/add", server, id), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	// REQ SETUP SUCCESS
	// NOW SEND REQ
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

// RmNode - removes a node
func RmNode(server, node, id string) error {
	rreq := struct {
		URL string `json:"url"`
	}{
		URL: node,
	}
	reqBody, err := json.Marshal(rreq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/master/record/%s/node/del", server, id), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	// REQ SETUP SUCCESS
	// NOW SEND REQ
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

// RdRec - redeploys a record
func RdRec(server, id string) error {
	req, err := http.NewRequest("Get", fmt.Sprintf("%s/api/master/record/%s/redeploy", server, id), nil)
	if err != nil {
		return err
	}
	// REQ SETUP SUCCESS
	// NOW SEND REQ
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
