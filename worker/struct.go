package worker



type gitHookReq struct {
	Config gitHookConfig `json:"config"`
}
type gitHookConfig struct {
	URL string `json:"url"`
	ContentType string `json:"content_type"`
}