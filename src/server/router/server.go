package router

import (
	"fmt"
	"os"

	controller "github.com/sofyan48/duck/app/controller"
	"github.com/sofyan48/duck/app/libs"
)

// Init Create Server
func Init() {
	r := Routes()
	controller.RoutesController(r)
	// Add Env PORT and HOST
	port := libs.GetEnvVariabel("APP_PORT", os.Getenv("APP_PORT"))
	host := libs.GetEnvVariabel("APP_HOST", os.Getenv("APP_HOST"))
	// Create Servers
	r.Run(fmt.Sprintf("%s:%s", host, port))
}
