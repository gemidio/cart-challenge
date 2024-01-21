package sqlite3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeed(t *testing.T) {
	setUp()
	defer tearDown()

	t.Run("successfully seeded", func(t *testing.T) {
		err := Seed()

		assert.NoError(t, err)
	})

	t.Run("seeding fails to load script file", func(t *testing.T) {
		t.Setenv("SQLITE_DATABASE_SEED_FILE", "")
		err := Seed()

		assert.Error(t, err)
	})

	t.Run("seeding fails in database querying", func(t *testing.T) {
		t.Setenv("SQLITE_DATABASE_SEED_FILE", "fixtures/invalid-seed.sql")
		err := Seed()

		assert.Error(t, err)
	})
}
