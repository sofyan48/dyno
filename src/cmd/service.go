package cmd

import (
	"log"
	"os"

	"github.com/sofyan48/dyno/src/config"
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
	cfg := config.Config{}
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
		library.Utils.LoadEnvirontment(Args.EnvPath)
		consulConfig := cfg.CosulConfig.Config()
		consulConfig.Address = os.Getenv("CONSUL_HOST")
		consulConfig.Scheme = "http"
		client, err := cfg.CosulConfig.New(consulConfig)
		if err != nil {
			log.Fatalln(err)
		}
		argsFile := Args.TemplatePath
		var templates string
		if argsFile == "" {
			templates = library.Utils.GetCurrentPath() + "/dyno.yml"
		} else {
			templates = argsFile
		}
		if !library.Utils.CheckFile(templates) {
			return cli.NewExitError("No Templates Parse", 1)
		}
		dataRegister, err := library.Utils.ServiceRegisterYML(argsFile)
		if err != nil {
			log.Fatalln(err)
		}
		library.Service.RegisterConsul(client, dataRegister)
		return nil
	}

	return command
}
