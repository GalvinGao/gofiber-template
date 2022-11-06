package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/GalvinGao/gofiber-template/pkg/app/appconfig"
)

func New(conf *appconfig.Config) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(conf.BunPostgresDSN)))
	db := bun.NewDB(sqldb, pgdialect.New())
	hook := bundebug.NewQueryHook(bundebug.WithEnabled(conf.BunDebug >= 1), bundebug.WithVerbose(conf.BunDebug >= 2))
	db.AddQueryHook(hook)
	return db
}
