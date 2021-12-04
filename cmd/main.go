package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// Run launches the commands
func Run() {
	app := &cli.App{
		Name:  "Lodestar",
		Usage: "Run the lodestar application",
		Commands: []*cli.Command{
			commandServer(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
