package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

func list() cli.Command {
	command := cli.Command{}
	command.Name = "ls"
	command.Usage = "ls [command]"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "File Template Path",
			Destination: &Args.TemplatePath,
		},
	}
	command.Action = func(c *cli.Context) error {
		library := Library{}
		cmd := c.Args()[0]
		if cmd == "service" {
			client, err := initConfigConsul()
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}

			data, err := library.Service.ListServiceConsul(client)
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
			for n, i := range data {
				fmt.Println("#####################################################")
				fmt.Println("List Service", n)
				fmt.Println("#####################################################")
				library.Utils.LogInfo("ID", i.ID)
				library.Utils.LogInfo("Name", i.Service)
				library.Utils.LogInfo("Addres", i.Address)
				library.Utils.LogInfo("Port", i.Port)
				library.Utils.LogInfo("Tags", i.Tags)
			}
		}

		return nil
	}

	return command
}
