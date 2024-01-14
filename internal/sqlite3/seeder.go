package sqlite3

import (
	"os"
)

func Seed() error {
	contents, err := os.ReadFile(os.Getenv("SQLITE_DATABASE_SEED_FILE"))

	if err != nil {
		return err
	}

	_, err = database.Exec(string(contents))

	if err != nil {
		return err
	}

	return nil
}
