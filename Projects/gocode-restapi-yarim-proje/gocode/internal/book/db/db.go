package db

import (
	"database/sql"
	"fmt"

	"github.com/cihanozhan/config"
	"github.com/cihanozhan/internal/model"
)

const dbPrefix = "BOOKS_"

type (
	DB struct {
		db *sql.DB
	}

	BooksDB interface {
		Ping() error
		GetBooks(pagination model.Pagination) ([]model.Book, error)
		GetBook(id string) (model.Book, error)
		CreateBook(book model.Book) (string, error)
		//UpdateBook(book model.Book) (string, error)
		//DeleteBook(id int64) (int64, error)
	}
)

func NewDB(db *sql.DB) *DB {
	postgreSQLDB := DB{
		db: db,
	}
	return &postgreSQLDB
}

func CreateUsersDbConnection() (*sql.DB, error) {
	Global := config.SetConfig(dbPrefix)

	connString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		Global.DBHost,
		Global.DBPort,
		Global.DBName,
		Global.DBUser,
		Global.DBPassword)

	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) Ping() error {
	return db.db.Ping()
}

func (db *DB) GetBooks(pg model.Pagination) ([]model.Book, error) {
	var book model.Book
	var books []model.Book

	// sel := "SELECT * FROM books"
	// if pg.Offset != 0 {
	// 	sel + = " OFFSET $1"
	// }

	stmt, err := db.db.Prepare("SELECT * FROM books OFFSET $1 LIMIT $2")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(pg.Offset, pg.Limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&book.ID, &book.Title, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (db *DB) GetBook(id string) (model.Book, error) {
	var book model.Book

	smtp, err := db.db.Prepare("SELECT * FROM books WHERE id = $1")
	if err != nil {
		return book, err
	}

	defer smtp.Close()

	err = smtp.QueryRow(id).Scan(&book.ID, &book.Title, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return book, err
	}

	return book, err

}

//log tutulmali
func (db *DB) CreateBook(book model.Book) (string, error) {
	var id string

	smtp, err := db.db.Prepare("INSERT INTO books(id, title, created_at, updated_at) VALUES($1,$2,$3,$4) RETURNING id")
	if err != nil {
		return "", err
	}

	defer smtp.Close()

	err = smtp.QueryRow(book.ID, book.Title, book.CreatedAt, book.UpdatedAt).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, err

}
