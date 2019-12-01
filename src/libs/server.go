package libs

import (
	"os"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/sofyan48/duck/src/configs"
)

// StartServer start server scheduler
func StartServer(cnf *config.Config) (*machinery.Server, error) {
	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// InitServer call server function
func InitServer(path string) (*machinery.Server, error) {
	cfg, err := configs.LoadMachineryConfig(path)
	if err != nil {
		os.Exit(3)
	}
	srv, err := StartServer(cfg)
	if err != nil {
		return srv, err
	}
	return srv, nil
}
