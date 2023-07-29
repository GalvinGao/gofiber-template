package migrations

import (
	"embed"

	"github.com/uptrace/bun/migrate"
)

// A collection of migrations.
var Migrations = migrate.NewMigrations()

//go:embed sql/*.sql
var sqlMigrations embed.FS

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}
