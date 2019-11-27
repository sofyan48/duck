package configs

import (
	"github.com/RichardKnop/machinery/v1/config"
)

// LoadMachineryConfig load config Function
func LoadMachineryConfig(path string) (*config.Config, error) {
	if path == ""{
		return config.NewFromYaml(path, true)
	}
	return config.NewFromEnvironment(true)
}
