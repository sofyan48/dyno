package cmd

import (
	"log"
	"os"

	"github.com/hashicorp/consul/api"
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
		library.Utils.LoadEnvirontment(Args.EnvPath)
		ymlRegis, err := library.Utils.ServiceRegisterYML(templates)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		cmd := c.Args()[0]
		if cmd == "register" {
			err = initServiceRegister(ymlRegis)
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
		}
		if cmd == "check" {
			err = initCheckService(ymlRegis)
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
			library.Utils.LogInfo("Service Register ", "OK")
		}
		return nil
	}

	return command
}

func initConfigConsul() (*api.Client, error) {
	// get consul client
	cfg := config.Config{}
	consulConfig := cfg.CosulConfig.Config()
	consulConfig.Address = os.Getenv("CONSUL_HOST")
	consulConfig.Scheme = "http"
	return cfg.CosulConfig.New(consulConfig)
}

func initCheckService(ymlRegis entity.ServiceRegisterYML) error {
	library := libs.Service{}
	regis := entity.ServiceRegister{}
	regis.Host = ymlRegis.Service.Host
	regis.Port = ymlRegis.Service.Port
	regis.ID = ymlRegis.Service.ID
	regis.Name = ymlRegis.Service.Name
	client, err := initConfigConsul()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return library.CheckServiceConsul(client, regis)
}

func initServiceRegister(ymlRegis entity.ServiceRegisterYML) error {
	library := Library{}

	regis := entity.ServiceRegister{}
	regis.Host = ymlRegis.Service.Host
	regis.Port = ymlRegis.Service.Port
	regis.ID = ymlRegis.Service.ID
	regis.Name = ymlRegis.Service.Name

	regis.HealthCheck = ymlRegis.HealthCheck.Endpoint
	regis.Interval = ymlRegis.HealthCheck.Interval
	regis.Timeout = ymlRegis.HealthCheck.Timeout
	client, err := initConfigConsul()
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
