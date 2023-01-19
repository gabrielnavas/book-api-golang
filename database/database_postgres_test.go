package database_test

import (
	"book-pai/database"
	"testing"
)

func TestDatabasePostgres(t *testing.T) {
	t.Run("postgres database testing", func(t *testing.T) {
		db, err := database.NewDatabasePostgres()

		if err != nil {
			t.Errorf("want nil; got '%s'", err)
		}
		if db == nil {
			t.Errorf("want an instance from db; got nil")
		}
	})
}
