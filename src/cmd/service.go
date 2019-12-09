package cmd

import (
	"github.com/sofyan48/dyno/src/libs"
	"github.com/urfave/cli"
)

// Library types
type Library struct {
	Utils   libs.Utils
	Service libs.Service
}

func service() cli.Command {
	library := Library{}
	command := cli.Command{}
	command.Name = "service"
	command.Usage = "service start, service configure"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "File Template Path",
			Destination: &Args.TemplatePath,
		},
	}
	command.Action = func(c *cli.Context) error {
		// init template file
		// argsFile := Args.TemplatePath
		// var templates string

		// if argsFile == "" {
		// 	templates = util.Utils. + "/duck.yml"
		// } else {
		// 	templates = argsFile
		// }
		// if !utils.Utils.CheckFile(templates) {
		// 	return cli.NewExitError("No Templates For Send", 1)
		// }
		// ymlData, err := libs.ReadYMLSend(templates)
		// libs.Check(err)

		//load environtment
		library.Utils.LoadEnvirontment(Args.EnvPath)
		library.Utils.LogInfo("OK ", "CUK")
		library.Service.Register()
		return nil
	}

	return command
}
