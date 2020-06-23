package structs

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
	Port     uint16  `json:"port"`
	Interval uint64  `json:"interval"` // check every x miliseconds
	Checks   uint8   `json:"checks"`   // number of failed checks before running fail command
	Run      *string `json:"run"`
}
