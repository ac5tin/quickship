package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	nethttp "net/http"
	"os"

	"github.com/buger/jsonparser"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// Clone - git clone a repo
func Clone(repourl string, branch string, path string, token string) error {
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "user",
			Password: token,
		},
		ReferenceName: plumbing.NewBranchReferenceName(branch),
		URL:           repourl,
		Progress:      os.Stdout,
	})
	if err != nil {
		return err
	}
	return nil
}

// Pull - git pull a repo
func Pull(repourl string, branch string, path string, token string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	if err := w.Pull(&git.PullOptions{
		Auth: &http.BasicAuth{
			Username: "user",
			Password: token,
		},
		Progress:      os.Stdout,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
	}); err != nil {
		return err
	}
	return nil
}

// CreateHook - creates a webhook and returns its id
func CreateHook(owner string, reponame string, token string, webhookurl string) (int64, error) {

	cfg := gitHookConfig{
		URL:         webhookurl,
		ContentType: "json",
	}
	hookreq := gitHookReq{
		Config: cfg,
	}

	reqbody, err := json.Marshal(hookreq)
	if err != nil {
		return 0, err
	}
	req, err := nethttp.NewRequest("POST", fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks", owner, reponame), bytes.NewBuffer(reqbody))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	// REQ SETUP SUCCESS
	client := &nethttp.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	hookid, err := jsonparser.GetInt(body, "id")
	if err != nil {
		return 0, err
	}
	return hookid, nil

}

// DeleteHook - deletes a hook from a repo
func DeleteHook(owner string, reponame string, token string, hookid int) error {
	req, err := nethttp.NewRequest("DELETE", fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks/%d", owner, reponame, hookid), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	client := &nethttp.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
