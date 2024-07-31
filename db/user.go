package db

import (
	"app/config"
	"database/sql"
	"errors"
	"fmt"
)

type Book struct {
	BookID         int    `json:"id" `
	Tittle         string `json:"tittle"`
	Author         string `json:"author"`
	PublishingYear int    `json:"publishing_year"`
	Genre          string `json:"genre"`
	AvailableCopy  int    `json:"available_copy"`
}

type PaginatedBooks struct {
	Books       []Book `json:"users"`
	TotalPage   int    `json:"total_page"`
	TotalItems  int    `json:"total_items"`
	Limit       int    `json:"limit"`
	CurrentPage int    `json:"current_page"`
}

func ReadBooks(page int, limit int, sortBy string, sortKey string, filter string, filterValue string, search string, searchValue string) (PaginatedBooks, error) {
	var paginatedBooks PaginatedBooks
	var books []Book
	var totalItems int

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	countQuery := "SELECT COUNT(*) FROM books"
	err := config.Db.QueryRow(countQuery).Scan(&totalItems)
	if err != nil {
		return paginatedBooks, errors.New("failed to get total count of books")
	}

	selectQuery := ("SELECT book_id, title, author, publishing_year, genre, available_copy FROM books")
	if filter != "" {
		selectQuery += fmt.Sprintf(" WHERE  %s='%s'", filter, filterValue)
	}

	if searchValue != "" {
		selectQuery += fmt.Sprintf(" WHERE author LIKE '%%%s%%'", searchValue)

	}
	selectQuery += fmt.Sprintf(" ORDER BY %s %s LIMIT $1 OFFSET $2", sortBy, sortKey)

	rows, err := config.Db.Query(selectQuery, limit, offset)
	if err != nil {
		return paginatedBooks, errors.New("failed to execute query")
	}

	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.BookID, &book.Tittle, &book.Author, &book.PublishingYear, &book.Genre, &book.AvailableCopy); err != nil {
			return paginatedBooks, errors.New("failed to scan rows")
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return paginatedBooks, errors.New("error iterating rows")
	}

	totalPages := (totalItems + limit - 1) / limit

	paginatedBooks = PaginatedBooks{
		Books:       books,
		TotalItems:  totalItems,
		TotalPage:   totalPages,
		CurrentPage: page,
		Limit:       limit,
	}

	return paginatedBooks, nil
}

func AddBook(b *Book) error {
	err := config.Db.QueryRow("INSERT INTO books(title,author, publishing_year,genre,available_copy) VALUES ($1,$2,$3,$4,$5) returning book_id", b.Tittle, b.Author, b.PublishingYear, b.Genre, b.AvailableCopy).Scan(&b.BookID)
	if err != nil {
		if err != sql.ErrNoRows {
			return errors.New("no user found")
		}
		return err
	}

	return nil
}

func Delete(id int) (Book, error) {
	var book Book

	err := config.Db.QueryRow("SELECT * FROM books WHERE book_id=$1", id).Scan(&book.BookID, &book.Tittle, &book.Author, &book.PublishingYear, &book.Genre, &book.AvailableCopy)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("no user found")
		}
		return book, err
	}

	_, err = config.Db.Exec("DELETE FROM books WHERE book_id = $1", id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func Update(id int, b *Book) (Book, error) {
	var book Book

	_, err := config.Db.Exec("UPDATE books SET title = $1, author = $2, publishing_year=$3,genre=$4,available_copy=$5 WHERE book_id = $6", b.Tittle, b.Author, b.PublishingYear, b.Genre, b.AvailableCopy, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("no book found")
		}
		return book, err
	}

	return book, nil
}

func ReadOneBook(id int) (Book, error) {
	var book Book

	err := config.Db.QueryRow("SELECT * FROM books WHERE book_id = $1", id).Scan(&book.BookID, &book.Tittle, &book.Author, &book.PublishingYear, &book.Genre, &book.AvailableCopy)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("no book found")
		}
		return book, err
	}

	return book, nil
}
