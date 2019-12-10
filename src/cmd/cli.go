package cmd

import (
	"os"

	"github.com/hashicorp/consul/api"
	"github.com/sofyan48/dyno/src/config"
	"github.com/sofyan48/dyno/src/libs"
	"github.com/urfave/cli"
)

var app *cli.App

// ArgsMapping object mapping
type ArgsMapping struct {
	EnvPath         string
	TemplatePath    string
	WorkerName      string
	WorkerConcurent string
	UUID            string
}

// Args Glabal Acces args command
var Args ArgsMapping

// Library types
type Library struct {
	Utils   libs.Utils
	Service libs.Service
}

// Init Initialise a CLI app
func Init() *cli.App {
	app = cli.NewApp()
	app.Name = "dyno"
	app.Usage = "dyno [command]"
	app.Author = "meong48"
	app.Email = "meongbego@gmail.com"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config, c",
			Usage:       "Load environtment config from `FILE`",
			Destination: &Args.EnvPath,
		},
	}
	return app
}

// AppCommands All Command line app
func AppCommands() *cli.App {
	app := Init()
	app.Commands = []cli.Command{
		service(),
		list(),
	}
	return app
}

func initConfigConsul() (*api.Client, error) {
	// get consul client
	cfg := config.Config{}
	consulConfig := cfg.CosulConfig.Config()
	consulConfig.Address = os.Getenv("CONSUL_HOST")
	consulConfig.Scheme = "http"
	return cfg.CosulConfig.New(consulConfig)
}
