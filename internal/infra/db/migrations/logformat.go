package migrations

import (
	"strings"

	"github.com/uptrace/bun/migrate"
)

// formatMigrations formats migrations for logging.
func formatMigrations(s migrate.MigrationSlice) string {
	var sb strings.Builder
	for i, m := range s {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(m.String())
	}
	return sb.String()
}
