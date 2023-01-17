package models_test

import (
	"book-pai/models"
	"testing"
	"time"
)

func TestTitle(t *testing.T) {

	makeLongString := func(howLong int) string {
		str := ""
		for i := 0; i < howLong; i++ {
			str += "a"
		}
		return str
	}

	t.Run("testing title", func(t *testing.T) {
		casesTitle := map[string]error{
			"a":                 models.ErrTitleIsShort,
			"":                  models.ErrTitleIsShort,
			makeLongString(256): models.ErrTitleIsLong,
			"ab":                nil,
		}

		for title, want := range casesTitle {
			_, got := models.NewBook(
				title,
				"any_isbn_12345",
				500,
				"any publishing place",
				"any publishing company",
				time.Now().UTC(),
			)
			if got != want {
				t.Errorf("got '%s', want: '%s'\n", got, want)
			}
		}
	})
}

func TestIsbn(t *testing.T) {
	t.Run("testing isbn", func(t *testing.T) {
		casesIsbn := map[string]error{
			"":               models.ErrIsbnDoesNotHaveCharactersEnoght,
			"123456789123":   models.ErrIsbnDoesNotHaveCharactersEnoght,
			"1231231231234":  nil,
			"12312312312345": models.ErrIsbnDoesNotHaveCharactersEnoght,
		}

		for isbn, want := range casesIsbn {
			_, got := models.NewBook(
				"any_title",
				isbn,
				500,
				"any publishing place",
				"any publishing company",
				time.Now().UTC(),
			)
			if got != want {
				t.Errorf("got '%s', want: '%s'\n", got, want)
			}
		}
	})
}
