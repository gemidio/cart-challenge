package sqlite3

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	t.Run("successful connection", func(t *testing.T) {
		assert.NoError(t, Connect())
	})

	t.Run("connection fail", func(t *testing.T) {
		t.Setenv("SQLITE_DATABASE", "user:password@/dbname")

		assert.Error(t, Connect())
	})
}

func setUp() {
	err := godotenv.Load("../../.env.testing")

	if err != nil {
		log.Println(err)
	}

	Connect()
}

func tearDown() {
	database.Close()
}
