package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/ramezanius/octoql/cmd"
	"github.com/ramezanius/octoql/pkg/config"
)

func main() {
	app := &cli.App{
		Name: "OctoQL",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "env",
				Value:   "DEV",
				Usage:   "set environment",
				EnvVars: []string{"ENV"},
				Aliases: []string{"e"},
			},
			&cli.StringFlag{
				Name:    "cfg",
				Value:   config.DefaultPath,
				Usage:   "set configuration path",
				Aliases: []string{"c"},
			},
		},
		Usage: "OctoQL’s command-line utility for administrative tasks",
		Commands: []*cli.Command{
			{
				Name:    "runserver",
				Action:  cmd.RunServer,
				Aliases: []string{"r"},
				Usage:   "Starts a Web server on the machine",
			},
			{
				Name:    "generate",
				Action:  cmd.Generate,
				Aliases: []string{"g"},
				Usage:   "Generates new code based on the changes detected to your schema",
			},
			{
				Name:    "migrate",
				Action:  cmd.Migrate,
				Aliases: []string{"m"},
				Usage:   "Synchronizes the database state with the current set of models and migrations",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
