package main

import (
	"os"

	"github.com/sofyan48/duck/src/cmd"
)

func main() {
	app := cmd.AppCommands()
	app.Run(os.Args)
}
