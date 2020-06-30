package worker

import (
	"net/http"
)

// Ping a url and return true if status code is 200
func Ping(url string) (bool, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	// server returned something successfully
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK, nil
}
