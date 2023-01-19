package repositories_test

import (
	"book-pai/database"
	"book-pai/models"
	"book-pai/repositories"
	"database/sql"
	"testing"
	"time"
)

func TestBookRepository(t *testing.T) {
	var db *sql.DB

	beforeAll := func() {
		db, _ = database.NewDatabasePostgres()
	}

	afterAll := func() {
		db.Exec("delete from public.book")
	}

	t.Run("testing repository", func(t *testing.T) {
		beforeAll()
		defer afterAll()

		book, _ := models.NewBook(
			"any_title",
			"any_isbn_1234",
			500,
			"any publishing company",
			time.Now().UTC(),
		)

		bookRepository := repositories.NewBookRepository(db)
		err := bookRepository.CreateBook(book)

		if err != nil {
			t.Errorf("want nil; got '%s'", err)
		}
	})

}
