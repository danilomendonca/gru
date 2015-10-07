package service

type Service struct {
	Name          string      `json:"name"`
	Type          string      `json:"type"`
	Image         string      `json:"image"`
	Constraints   Constraints `json:"constraints"` //Needed?
	Configuration Config      `json: "configuration"`
}

// This can maybe be eliminated
type Constraints struct {
	CpuMax    float64 `json:"cpumax"`
	CpuMin    float64 `json:"cpumin"`
	MinActive int     `json:"minactive"`
	MaxActive int     `json:"maxactive"`
}

type Config struct {
	Cmd          []string               `json:"cmd"`
	Volumes      map[string]struct{}    `json:"volumes"`
	Entrypoint   string                 `json:"entrypoint"`
	Memory       string                 `json:"memory"`
	CpuShares    int64                  `json:"cpushares"`
	CpuSet       int64                  `json:"cpuset"`
	PortBindings map[string]PortBinding `json:"portbindings"`
	Links        []string               `json:"links"`
}

type PortBinding struct {
	HostIp   string `json:"hostip"`
	HostPort string `json:"hostport"`
}