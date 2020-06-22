package deploy

// Deploy - main deploy struct
type Deploy struct {
	Nodes   []string    `json:"nodes"`
	GitRepo string      `json:"gitrepo"`
	Branch  string      `json:"branch"`
	Build   *string     `json:"build"`
	Run     string      `json:"run"`
	Clean   *string     `json:"clean"`
	Health  HealthCheck `json:"health"`
}

// HealthCheck - specify the api apth and check and command to run if fails check
type HealthCheck struct {
	Path     string  `json:"path"`
	Interval uint64  `json:"interval"`
	Checks   uint8   `json:"checks"`
	Run      *string `json:"run"`
}
