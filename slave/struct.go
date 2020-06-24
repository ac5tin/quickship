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
