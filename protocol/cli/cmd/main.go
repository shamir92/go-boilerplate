package main

import (
	"fmt"
	"log"
	"os"
	commands "simple-invitation/protocol/cli/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	var pingCommands commands.IPingCommands = commands.NewPingCommand()
	app := &cli.App{
		Commands: []*cli.Command{
			&cli.Command{
				Name:     "ping-controller",
				Aliases:  []string{"pc"},
				Category: "ping",
				Usage:    "Ping via controller interface",
				Action:   pingCommands.GetPingController,
			},
			&cli.Command{
				Name:     "ping-usecase",
				Aliases:  []string{"pu"},
				Category: "ping",
				Usage:    "Ping via usecase interface",
				Action:   pingCommands.GetPingUseCase,
				OnUsageError: func(cCtx *cli.Context, err error, isSubcommand bool) error {
					fmt.Fprintf(cCtx.App.Writer, "for shame\n")
					return err
				},
			},
			&cli.Command{
				Name:     "ping-repository",
				Aliases:  []string{"pr"},
				Category: "ping",
				Usage:    "Ping via repository interface",
				Action:   pingCommands.GetPingRepository,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
