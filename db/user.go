package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
)

type Book struct {
	BookID         int    `json:"id"`
	Tittle         string `json:"tittle"`
	Author         string `json:"author"`
	PublishingYear int    `json:"publishing_year"`
	Genre          string `json:"genre"`
	AvailableCopy  int    `json:"available_copy"`
}

type User struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PaginatedBooks struct {
	Books       []Book `json:"books"`
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

	countQuery, args, err := GetQueryBuilder().Select("COUNT(*)").From("books").ToSql()
	if err != nil {
		return paginatedBooks, errors.New("failed to execute query")
	}

	err = WriteDb.QueryRow(countQuery, args...).Scan(&totalItems)
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

	rows, err := WriteDb.Query(query, args...)
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
	addQuery, args, err := GetQueryBuilder().
		Insert("books").
		Columns("title", "author", "publishing_year", "genre", "available_copy").
		Values(b.Tittle, b.Author, b.PublishingYear, b.Genre, b.AvailableCopy).
		Suffix("Returning book_id").
		ToSql()

	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		return err
	}

	err = WriteDb.QueryRow(addQuery, args...).Scan(&b.BookID)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}

		return err
	}

	return err
}

func SignUp(b *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	addQuery, args, err := GetQueryBuilder().
		Insert("users").
		Columns("username", "email", "password").
		Values(b.Username, b.Email, string(hashedPassword)).
		Suffix("Returning user_id").
		ToSql()
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}

		return err
	}

	err = WriteDb.QueryRow(addQuery, args...).Scan(&b.UserID)
	if err != nil {
		return err
	}

	return err
}

func ReadOneBook(id int) (Book, error) {
	var book Book

	selectQuery, args, err := GetQueryBuilder().
		Select("book_id", "title", "author", "publishing_year", "genre", "available_copy").
		From("books").
		Where(squirrel.Eq{"book_id": id}).
		ToSql()
	if err != nil {
		if err == sql.ErrNoRows {
			return book, err
		}

		return book, err
	}

	err = WriteDb.QueryRow(selectQuery, args...).
		Scan(&book.BookID, &book.Tittle, &book.Author, &book.PublishingYear, &book.Genre, &book.AvailableCopy)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, err
		}

		return book, err
	}

	_, err = WriteDb.Exec(selectQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, err
		}
		return book, err
	}

	return book, err
}

func Delete(id int) (Book, error) {
	var book Book

	query, args, err := GetQueryBuilder().
		Select("book_id", "title", "author", "publishing_year", "genre", "available_copy").
		From("books").
		Where(squirrel.Eq{"book_id": id}).
		ToSql()

	if err != nil {
		return book, err
	}

	err = WriteDb.QueryRow(query, args...).
		Scan(&book.BookID, &book.Tittle, &book.Author, &book.PublishingYear, &book.Genre, &book.AvailableCopy)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, err
		}
		return book, err
	}

	deleteQuery, args, err := GetQueryBuilder().
		Delete("*").
		From("books").
		Where(squirrel.Eq{"book_id": id}).
		ToSql()
	if err != nil {
		return book, err
	}

	_, err = WriteDb.Exec(deleteQuery, args...)
	if err != nil {
		return book, err
	}

	return book, err
}

func Update(id int, b *Book) (Book, error) {
	var book Book
	query, args, err := GetQueryBuilder().
		Select("book_id", "title", "author", "publishing_year", "genre", "available_copy").
		From("books").
		Where(squirrel.Eq{"book_id": id}).
		ToSql()

	if err != nil {
		return book, err
	}

	err = WriteDb.QueryRow(query, args...).Scan(&book.BookID, &book.Tittle, &book.Author, &book.PublishingYear, &book.Genre, &book.AvailableCopy)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, err
		}
		return book, err
	}

	updateQuery, args, err := GetQueryBuilder().
		Update("books").
		Set("title", b.Tittle).
		Set("author", b.Author).
		Set("publishing_year", b.PublishingYear).
		Set("genre", b.Genre).
		Set("available_copy", b.AvailableCopy).
		Where(squirrel.Eq{"book_id": id}).
		ToSql()

	if err != nil {
		if err == sql.ErrNoRows {
			return book, err
		}
		return book, err
	}

	_, err = WriteDb.Exec(updateQuery, args...)
	if err != nil {
		return book, err
	}

	return book, err
}

func ReadOneUser(id int) (User, error) {
	var user User

	selectQuery, args, err := GetQueryBuilder().
		Select("user_id", "username", "email", "password").
		From("users").
		Where(squirrel.Eq{"user_id": id}).
		ToSql()

	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}

	err = WriteDb.QueryRow(selectQuery, args...).Scan(&user.UserID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}

	_, err = WriteDb.Exec(selectQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}

		return user, err
	}

	return user, nil
}
