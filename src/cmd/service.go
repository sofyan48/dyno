package cmd

import (
	"log"

	"github.com/sofyan48/dyno/src/libs"
	"github.com/sofyan48/dyno/src/libs/entity"
	"github.com/urfave/cli"
)

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

		cli.StringFlag{
			Name:        "ID, i",
			Usage:       "Service Identity",
			Destination: &Args.ID,
		},
	}
	command.Action = func(c *cli.Context) error {
		library := Library{}

		cmd := c.Args()[0]
		if cmd == "add" {
			ymlRegis, err := checkTemplate()
			err = initServiceRegister(ymlRegis)
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
		}
		if cmd == "check" {

			ymlRegis, err := checkTemplate()
			err = initCheckService(ymlRegis)
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
			library.Utils.LogInfo("Service Register ", "OK")
		}

		if cmd == "delete" {
			ymlRegis, err := checkTemplate()
			err = initDeregister(ymlRegis.Service.ID)
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
		}

		if cmd == "lookup" {
			client, err := initConfigConsul()
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
			var ID string
			if Args.ID == "" {
				ymlRegis, err := checkTemplate()
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				ID = ymlRegis.Service.ID
			} else {
				ID = Args.ID
			}
			address, err := library.Service.ServiceLookupConsul(client, ID)
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
			library.Utils.LogInfo("Result", address)
		}

		if cmd == "health" {
			client, err := initConfigConsul()
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
			var ID string
			if Args.ID == "" {
				ymlRegis, err := checkTemplate()
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				ID = ymlRegis.Service.ID
			} else {
				ID = Args.ID
			}
			status, _, _ := library.Service.GetHealthByIDConsul(client, ID)
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
			library.Utils.LogInfo("Health", status)
		}
		return nil
	}

	return command
}

func checkTemplate() (entity.ServiceRegisterYML, error) {
	argsFile := Args.TemplatePath
	library := Library{}
	templates, err := library.Utils.CheckTemplateFile(argsFile)
	ymlRegis, err := library.Utils.ServiceRegisterYML(templates)
	if err != nil {
		return ymlRegis, err
	}
	return ymlRegis, nil
}

func initDeregister(ID string) error {
	library := Library{}
	library.Utils.LoadEnvirontment(Args.EnvPath)
	client, err := initConfigConsul()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return library.Service.UnRegisterConsul(client, ID)
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
	regis.Tags = ymlRegis.Service.Tags
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
