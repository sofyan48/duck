package cmd

import (
	"os"

	"github.com/RichardKnop/machinery/v1"
	"github.com/sofyan48/duck/src/configs"
	"github.com/sofyan48/duck/src/libs"
)

// InitServer call server function
func InitServer() (*machinery.Server, error) {
	cfg, err := configs.LoadMachineryConfig(Args.EnvPath)
	if err != nil {
		os.Exit(3)
	}
	srv, err := libs.StartServer(cfg)
	if err != nil {
		return srv, err
	}
	return srv, nil
}
