package bookrepository

import (
	"database/sql"

	"github.com/go_course/2-rest-api/book-list/driver"
	"github.com/go_course/2-rest-api/book-list/models"
)

// BookRepository define structure of Book Repository
type BookRepository struct{}

// GetBooks DB structure for get all books
func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query("select * from books")
	driver.LogFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		driver.LogFatal(err)

		books = append(books, book)
	}

	return books
}

// GetBook return single BOOK
func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book {
	rows := db.QueryRow("select * from books where id=$1", id)

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	driver.LogFatal(err)

	return book
}

// AddBook adding a book to DB
func (b BookRepository) AddBook(db *sql.DB, book models.Book) int {
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&book.ID)

	driver.LogFatal(err)

	return book.ID
}

// UpdateBook updating book in database
func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)

	driver.LogFatal(err)

	rowsUpdated, err := result.RowsAffected()
	driver.LogFatal(err)

	return rowsUpdated
}

// RemoveBook removing book from DB
func (b BookRepository) RemoveBook(db *sql.DB, id int) int64 {
	result, err := db.Exec("delete from books where id = $1", id)
	driver.LogFatal(err)

	rowsDeleted, err := result.RowsAffected()
	driver.LogFatal(err)

	return rowsDeleted
}
