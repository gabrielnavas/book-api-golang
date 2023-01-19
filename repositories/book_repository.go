package repositories

import (
	"book-pai/models"
	"database/sql"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) CreateBook(book *models.Book) error {
	sql := `
		INSERT INTO public.book (
			id, 
			title, 
			isbn, 
			number_of_pages, 
			publishing_company, 
			year_of_edition
		) VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.Exec(
		sql,
		book.Id,
		book.Title,
		book.Isbn,
		book.NumberOfPages,
		book.PublishingCompany,
		book.YearOfEdition,
	)
	return err
}
