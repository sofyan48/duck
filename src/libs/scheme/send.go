package scheme

import "time"

// GetTask mapping
type GetTask struct {
	Duck struct {
		UUID   string `yaml:"uuid" json:"uuid"`
		Action struct {
			Trigger string `yaml:"trigger" json:"trigger"`
			Worker  string `yaml:"worker" json:"worker"`
		} `yaml:"action" json:"action"`
	} `yaml:"duck" json:"duck"`
}

// SendTask mapping send task to worker
type SendTask struct {
	Duck struct {
		Task   string `yaml:"task" json:"tasks"`
		Action struct {
			URL     string `yaml:"url" json:"url"`
			Method  string `yaml:"method" json:"method"`
			Trigger string `yaml:"trigger" json:"trigger"`
			Worker  string `yaml:"worker" json:"worker"`
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
			Loop bool `yaml:"loop" json:"loop"`
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

// SendResponse ...
type SendResponse struct {
	UUID      string      `json:"uuid"`
	TaskName  string      `json:"task_name"`
	QueueName string      `json:"queue_name"`
	Args      interface{} `json:"args"`
	CreatedAt time.Time   `json:"created_at"`
}

// ActionScheme ...
type ActionScheme struct {
	URL     string
	Method  string
	Trigger string
	Worker  string
}

// ActionScheme ...
type SettingScheme struct {
	Loop       bool
	TimeResult int64
}
