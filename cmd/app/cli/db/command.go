package db

import (
	"context"
	"strings"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v3"
	"go.uber.org/fx"

	cliapp "github.com/GalvinGao/gofiber-template/cmd/app/cli"
	"github.com/GalvinGao/gofiber-template/internal/infra/db/migrations"
)

type dbCommandDeps struct {
	fx.In

	DB *bun.DB
}

func Command() *cli.Command {
	migrator := func() *migrations.Migrator {
		var deps dbCommandDeps
		cliapp.Start(fx.Populate(&deps))

		m := migrate.NewMigrator(deps.DB, migrations.Migrations)
		return migrations.NewMigrator(m)
	}

	return &cli.Command{
		Name:  "db",
		Usage: "manage database migrations",
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create migration tables",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return migrator().Init(ctx)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return migrator().Migrate(ctx)
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback the last migration group",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return migrator().Rollback(ctx)
				},
			},
			{
				Name:  "lock",
				Usage: "lock migrations",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return migrator().Lock(ctx)
				},
			},
			{
				Name:  "unlock",
				Usage: "unlock migrations",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return migrator().Unlock(ctx)
				},
			},
			{
				Name:  "create_go",
				Usage: "create Go migration",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					name := strings.Join(cmd.Args().Slice(), "_")
					return migrator().CreateGoMigration(ctx, name)
				},
			},
			{
				Name:  "create_sql",
				Usage: "create up and down SQL migrations",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					name := strings.Join(cmd.Args().Slice(), "_")
					return migrator().CreateSQLMigrations(ctx, name)
				},
			},
			{
				Name:  "status",
				Usage: "print migrations status",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return migrator().Status(ctx)
				},
			},
			{
				Name:  "mark_applied",
				Usage: "mark migrations as applied without actually running them",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return migrator().MarkApplied(ctx)
				},
			},
		},
	}
}
