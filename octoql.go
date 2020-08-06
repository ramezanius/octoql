package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/ramezanius/octoql/cmd"
)

func main() {
	app := &cli.App{
		Name:  "OctoQL",
		Usage: "OctoQL’s command-line utility for administrative tasks",
		Commands: []*cli.Command{
			{
				Name:   "runserver",
				Action: cmd.RunServer,
				Usage:  "Starts a Web server on the machine",
			},
			{
				Name:   "generate",
				Action: cmd.Generate,
				Usage:  "Generates new code based on the changes detected to your schema",
			},
			{
				Name:   "migrate",
				Action: cmd.Migrate,
				Usage:  "Synchronizes the database state with the current set of models and migrations",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
