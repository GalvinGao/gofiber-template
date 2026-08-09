package app

import (
	"context"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"

	"github.com/GalvinGao/gofiber-template/cmd/app/cli/db"
	"github.com/GalvinGao/gofiber-template/cmd/app/server"
)

func Run() {
	if err := Command().Run(context.Background(), os.Args); err != nil {
		log.Fatal().Err(err).Msg("failed to run app")
	}
}

func Command() *cli.Command {
	return &cli.Command{
		Name: "app",
		Commands: []*cli.Command{
			server.Command(),
			db.Command(),
		},
	}
}
