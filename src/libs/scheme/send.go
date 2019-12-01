package scheme

// SendTask mapping send task to worker
type SendTask struct {
	Duck struct {
		Task   string `yaml:"task"`
		Action struct {
			URL     string `yaml:"url"`
			Method  string `yaml:"method"`
			Trigger string `yaml:"trigger"`
		} `yaml:"action"`
		Headers []struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
			Type  string `yaml:"type"`
		} `yaml:"headers"`
		Params []struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
			Type  string `yaml:"type"`
		} `yaml:"parameter"`
		Body []struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
			Type  string `yaml:"type"`
		} `yaml:"body"`
		Setting struct {
			ETA                         int    `yaml:"eta"`
			RetryCount                  int    `yaml:"retryCount"`
			RoutingKey                  string `yaml:"routingKey"`
			IgnoreWhenTaskNotRegistered bool   `yaml:"ignoreWhenTaskNotRegistered"`
		} `yaml:"setting"`
	} `yaml:"duck"`
}
