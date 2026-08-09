package server

import (
	"context"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "start server",
		Action: func(context.Context, *cli.Command) error {
			Run()
			return nil
		},
	}
}
