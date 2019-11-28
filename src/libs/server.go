package libs

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

// StartServer start server scheduler
func StartServer(cnf *config.Config) (*machinery.Server, error) {
	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}
	return server, nil
}
