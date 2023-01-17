package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTitleIsShort = errors.New("title must be more than 2 or more characters")
	ErrTitleIsLong  = errors.New("title must be less than 255 characters")
)

type Book struct {
	Id                uuid.UUID
	Title             string
	Isbn              string
	NumberOfPages     int
	PublishingPlace   string
	PublishingCompany string
	YearOfEdition     time.Time
}

func NewBook(
	title string,
	isbn string,
	numberOfPages int,
	publishingPlace string,
	publishingCompany string,
	yearOfEdition time.Time,
) (*Book, error) {
	book := &Book{
		Title:             title,
		Isbn:              isbn,
		NumberOfPages:     numberOfPages,
		PublishingCompany: publishingCompany,
		YearOfEdition:     yearOfEdition,
	}
	err := book.validate()
	return book, err
}

func (b *Book) validate() error {
	if len(b.Title) < 2 {
		return ErrTitleIsShort
	}
	if len(b.Title) > 255 {
		return ErrTitleIsLong
	}
	return nil
}
