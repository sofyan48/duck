package scheme

// SendTask mapping send task to worker
type SendTask struct {
	Duck struct {
		Task   string `yaml:"task" json:"tasks"`
		Action struct {
			URL     string `yaml:"url" json:"url"`
			Method  string `yaml:"method" json:"method"`
			Trigger string `yaml:"trigger" json:"trigger"`
		} `yaml:"action" json:"action"`
		Headers []struct {
			Name  string `yaml:"name" json:"name"`
			Value string `yaml:"value" json:"value"`
			Type  string `yaml:"type" json:"type"`
		} `yaml:"headers" json:"headers"`
		Params []struct {
			Name  string `yaml:"name" json:"name"`
			Value string `yaml:"value" json:"value"`
			Type  string `yaml:"type" json:"type"`
		} `yaml:"parameter" json:"parameter"`
		Body []struct {
			Name  string `yaml:"name" json:"name"`
			Value string `yaml:"value" json:"value"`
			Type  string `yaml:"type" json:"type"`
		} `yaml:"body" json:"body"`
		Setting struct {
			ETA                         int    `yaml:"eta" json:"eta"`
			RetryCount                  int    `yaml:"retryCount" json:"retryCount"`
			RoutingKey                  string `yaml:"routingKey" json:"routingKey"`
			IgnoreWhenTaskNotRegistered bool   `yaml:"ignoreWhenTaskNotRegistered" json:"ignoreWhenTaskNotRegistered"`
		} `yaml:"setting" json:"setting"`
	} `yaml:"duck" json:"duck"`
}

// Query ...
type Query struct {
	Name  string `yaml:"name" json:"name"`
	Value string `yaml:"value" json:"value"`
	Type  string `yaml:"type" json:"type"`
}

// ResultsSend ...
type ResultsSend struct {
	Result interface{} `json:"result"`
	UUID   string      `json:"uuid"`
}
