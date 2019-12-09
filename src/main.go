package main

import (
	"os"

	"github.com/sofyan48/dyno/src/cmd"
)

func main() {
	app := cmd.AppCommands()
	app.Run(os.Args)
}
