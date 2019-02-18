package migrations

import (
	"github.com/go-pg/migrations"
)

// We need zero migration .go file for go-pg/migrations
func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		return nil
	})
}
