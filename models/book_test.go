package models_test

import (
	"book-pai/models"
	"testing"
	"time"
)

func makeLongString(howLong int) string {
	str := ""
	for i := 0; i < howLong; i++ {
		str += "a"
	}
	return str
}

func TestBookModel(t *testing.T) {

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
