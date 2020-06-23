package slave

// CmdReq - command request
type CmdReq struct {
	Command string `json:"command"`
}

type cloneReq struct {
	Repo   string `json:"repo"`
	Branch string `json:"branch"`
}
