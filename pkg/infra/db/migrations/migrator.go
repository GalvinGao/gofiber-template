package migrations

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun/migrate"
)

type Migrator struct {
	m *migrate.Migrator
}

func NewMigrator(m *migrate.Migrator) *Migrator {
	return &Migrator{m: m}
}

func (m *Migrator) Init(ctx context.Context) error {
	return m.m.Init(ctx)
}

func (m *Migrator) Migrate(ctx context.Context) error {
	group, err := m.m.Migrate(ctx)
	if err != nil {
		return err
	}

	if group.IsZero() {
		log.Info().Msg("no migrations to apply")
		return nil
	}

	log.Info().
		Str("migrations", formatMigrations(group.Migrations)).
		Msg("applied migrations")

	return nil
}

func (m *Migrator) Rollback(ctx context.Context) error {
	group, err := m.m.Rollback(ctx)
	if err != nil {
		return err
	}

	if group.IsZero() {
		log.Info().Msg("no migrations to rollback")
		return nil
	}

	log.Info().
		Interface("migrations", formatMigrations(group.Migrations)).
		Msg("rolled back migrations")
	return nil
}

func (m *Migrator) Lock(ctx context.Context) error {
	return m.m.Lock(ctx)
}

func (m *Migrator) Unlock(ctx context.Context) error {
	return m.m.Unlock(ctx)
}

func (m *Migrator) CreateGoMigration(ctx context.Context, name string) error {
	mf, err := m.m.CreateGoMigration(ctx, name)
	if err != nil {
		return err
	}

	log.Info().
		Str("name", mf.Name).
		Str("path", mf.Path).
		Msg("created go migration file")

	return nil
}

func (m *Migrator) CreateSQLMigrations(ctx context.Context, name string) error {
	files, err := m.m.CreateSQLMigrations(ctx, name)
	if err != nil {
		return err
	}

	for _, mf := range files {
		log.Info().
			Str("name", mf.Name).
			Str("path", mf.Path).
			Msg("created sql migration file")
	}

	return nil
}

func (m *Migrator) Status(ctx context.Context) error {
	ms, err := m.m.MigrationsWithStatus(ctx)
	if err != nil {
		return err
	}

	log.Info().
		Interface("migrations", formatMigrations(ms)).
		Msg("migrations status")

	log.Info().
		Interface("unapplied", formatMigrations(ms.Unapplied())).
		Msg("unapplied migrations")

	log.Info().
		Stringer("last_group", ms.LastGroup()).
		Msg("last migration group")

	return nil
}

func (m *Migrator) MarkApplied(ctx context.Context) error {
	group, err := m.m.Migrate(ctx, migrate.WithNopMigration())
	if err != nil {
		return err
	}

	if group.IsZero() {
		log.Info().Msg("no migrations to apply")
		return nil
	}

	log.Info().
		Interface("migrations", formatMigrations(group.Migrations)).
		Msg("marked migrations as applied")
	return nil
}
