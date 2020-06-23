package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"quickship/store"
	"quickship/structs"
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
