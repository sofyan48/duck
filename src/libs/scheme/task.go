package scheme

// RegisterTask mapping register task worker
type RegisterTask struct {
	Worker struct {
		Register []struct {
			Action string `yaml:"action"`
			Name   string `yaml:"name"`
		} `yaml:"register"`
	} `yaml:"worker"`
}
