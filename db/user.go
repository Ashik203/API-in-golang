package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
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

	InitQueryBuilder()

	countQuery, args, err := GetQueryBuilder().Select("COUNT(*)").From("books").ToSql()
	if err != nil {
		return paginatedBooks, errors.New("failed to execute query")

	}
	err = Db.QueryRow(countQuery, args...).Scan(&totalItems)
	if err != nil {
		return paginatedBooks, errors.New("failed to get total count of books")
	}

	selectQuery := GetQueryBuilder().
		Select("book_id", "title", "author", "publishing_year", "genre", "available_copy").
		From("books")

	if filter != "" && filterValue != "" {
		selectQuery = selectQuery.Where(squirrel.Eq{filter: filterValue})
	}

	if searchValue != "" {
		selectQuery = selectQuery.Where(squirrel.Like{"author": "%" + searchValue + "%"})

	}
	selectQuery = selectQuery.OrderBy(fmt.Sprintf("%s %s", sortBy, sortKey)).Limit(uint64(limit)).Offset(uint64(offset))

	query, args, err := selectQuery.ToSql()
	if err != nil {
		return paginatedBooks, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := Db.Query(query, args...)
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
	InitQueryBuilder()

	addQuery, args, err := GetQueryBuilder().Insert("books").Columns("title", "author", "publishing_year", "genre", "available_copy").Values(b.Tittle, b.Author, b.PublishingYear, b.Genre, b.AvailableCopy).Suffix("Returning book_id").ToSql()

	if err != nil {
		if err != sql.ErrNoRows {
			return errors.New("no user found")
		}
		return err
	}
	err = Db.QueryRow(addQuery, args...).Scan(&b.BookID)
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

	InitQueryBuilder()

	query, args, err := GetQueryBuilder().
		Select("book_id", "title", "author", "publishing_year", "genre", "available_copy").
		From("books").
		Where(squirrel.Eq{"book_id": id}).
		ToSql()

	if err != nil {
		return book, fmt.Errorf("error building query: %w", err)
	}

	err = Db.QueryRow(query, args...).Scan(&book.BookID, &book.Tittle, &book.Author, &book.PublishingYear, &book.Genre, &book.AvailableCopy)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("no book found with the given ID")
		}
		return book, fmt.Errorf("error executing query: %w", err)
	}

	deleteQuery, args, err := GetQueryBuilder().Delete("*").From("books").Where(squirrel.Eq{"book_id": id}).ToSql()
	if err != nil {
		return book, fmt.Errorf("error building query: %w", err)
	}

	_, err = Db.Exec(deleteQuery, args...)
	if err != nil {
		return book, err
	}

	return book, nil
}

func Update(id int, b *Book) (Book, error) {
	var book Book

	InitQueryBuilder()
	updateQuery, args, err := GetQueryBuilder().Update("books").Set("title", b.Tittle).Set("author", b.Author).Set("publishing_year", b.PublishingYear).Set("genre", b.Genre).Set("available_copy", b.AvailableCopy).Where(squirrel.Eq{"book_id": id}).ToSql()

	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("no book found")
		}
		return book, err
	}

	_, err = Db.Exec(updateQuery, args...)
	if err != nil {
		return book, err
	}

	return book, nil
}

func ReadOneBook(id int) (Book, error) {
	var book Book

	InitQueryBuilder()
	selectQuery, args, err := GetQueryBuilder().Select("book_id", "title", "author", "publishing_year", "genre", "available_copy").
		From("books").
		Where(squirrel.Eq{"book_id": id}).
		ToSql()
	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("no book found")
		}
		return book, err

	}
	err = Db.QueryRow(selectQuery, args...).Scan(&book.BookID, &book.Tittle, &book.Author, &book.PublishingYear, &book.Genre, &book.AvailableCopy)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("no book found with the given ID")
		}
		return book, fmt.Errorf("error executing query: %w", err)
	}
	_, err = Db.Exec(selectQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("no book found")
		}
		return book, err
	}

	return book, nil
}
