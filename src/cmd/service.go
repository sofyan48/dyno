package cmd

import (
	"log"
	"os"

	"github.com/sofyan48/dyno/src/config"
	"github.com/sofyan48/dyno/src/libs"
	"github.com/sofyan48/dyno/src/libs/entity"
	"github.com/urfave/cli"
)

// Library types
type Library struct {
	Utils   libs.Utils
	Service libs.Service
}

func service() cli.Command {
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
		library := Library{}
		argsFile := Args.TemplatePath
		templates, err := library.Utils.CheckTemplateFile(argsFile)
		err = initServiceRegister(templates)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		return nil
	}

	return command
}

func initServiceRegister(path string) error {
	library := Library{}

	library.Utils.LoadEnvirontment(Args.EnvPath)

	ymlRegis, err := library.Utils.ServiceRegisterYML(path)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	regis := entity.ServiceRegister{}
	regis.Host = ymlRegis.Service.Host
	regis.Port = ymlRegis.Service.Port
	regis.ID = ymlRegis.Service.ID
	regis.Name = ymlRegis.Service.Name

	regis.HealthCheck = ymlRegis.HealthCheck.Endpoint
	regis.Interval = ymlRegis.HealthCheck.Interval
	regis.Timeout = ymlRegis.HealthCheck.Timeout
	// get consul client
	cfg := config.Config{}
	consulConfig := cfg.CosulConfig.Config()
	consulConfig.Address = os.Getenv("CONSUL_HOST")
	consulConfig.Scheme = "http"
	client, err := cfg.CosulConfig.New(consulConfig)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	err = library.Service.RegisterConsul(client, regis)
	if err != nil {
		library.Utils.LogFatal("Error: ", err)
	}
	return err

}
