package cmd

import (
	"github.com/sofyan48/dyno/src/libs"
	"github.com/urfave/cli"
)

// Library types
type Library struct {
	Utils libs.Libs
}

func service() cli.Command {
	util := Library{}
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
		util.Utils.LoadEnvirontment(Args.EnvPath)
		util.Utils.LogInfo("OK ", "CUK")
		return nil
	}

	return command
}
