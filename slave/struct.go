package slave

// CmdReq - command request
type CmdReq struct {
	Command string `json:"command"`
}

// CloneReq - clone request
type CloneReq struct {
	Repo   string `json:"repo"`
	Branch string `json:"branch"`
}

// PullReq - pull request
type PullReq struct {
	Branch string `json:"branch"`
}

// PingReq - ping request
type PingReq struct {
	Port uint16 `json:"port"`
	Path string `json:"path"`
}
