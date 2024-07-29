package controller

import (
	"database/sql"
	"errors"
	"strconv"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Book struct {
	BookID         int    `json:"id"`
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

func GetPaginatedBooks(db *sql.DB, page int, limit int) (PaginatedBooks, error) {
	var paginatedBooks PaginatedBooks
	var books []Book

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	var totalItems int
	err := db.QueryRow("SELECT COUNT(*) from books").Scan(&totalItems)
	if err != nil {
		return paginatedBooks, errors.New("failed to get total count of books")

	}

	rows, err := db.Query("SELECT book_id, title,author,publishing_year,genre,available_copy FROM books LIMIT $1 OFFSET $2", limit, offset)
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

// func GetPaginatedUsers(db *sql.DB, page int, limit int) (PaginatedUsers, error) {
// 	var paginatedUsers PaginatedUsers
// 	var users []User

// 	if page < 1 {
// 		page = 1
// 	}
// 	if limit < 1 {
// 		limit = 10
// 	}
// 	offset := (page - 1) * limit

// 	var totalItems int
// 	err := db.QueryRow("SELECT COUNT(*) from users").Scan(&totalItems)
// 	if err != nil {
// 		return paginatedUsers, errors.New("failed to get total count of users")

// 	}

// 	rows, err := db.Query("SELECT id, username, password FROM users LIMIT $1 OFFSET $2", limit, offset)
// 	if err != nil {
// 		return paginatedUsers, errors.New("failed to execute query")
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var user User
// 		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
// 			return paginatedUsers, errors.New("failed to scan rows")
// 		}
// 		users = append(users, user)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return paginatedUsers, errors.New("error iterating rows")
// 	}

// 	totalPages := (totalItems + limit - 1) / limit

// 	paginatedUsers = PaginatedUsers{
// 		Users:       users,
// 		TotalItems:  totalItems,
// 		TotalPage:   totalPages,
// 		CurrentPage: page,
// 		Limit:       limit,
// 	}

// 	return paginatedUsers, nil
// }

func AddBook(db *sql.DB, b *Book) error {
	err := db.QueryRow("INSERT INTO books(title,author, publishing_year,genre,available_copy) VALUES ($1,$2,$3,$4,$5) returning book_id", b.Tittle, b.Author, b.PublishingYear, b.Genre, b.AvailableCopy).Scan(&b.BookID)
	if err != nil {
		if err != sql.ErrNoRows {
			return errors.New("no user found")
		}
		return err
	}
	return nil
}

// func Add(db *sql.DB, u *User) error {
// 	err := db.QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id", u.Username, u.Password).Scan(&u.ID)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return errors.New("no user found")
// 		}

// 		return err
// 	}

// 	return nil
// }

func Delete(db *sql.DB, id string) (Book, error) {
	var book Book

	bookID, err := strconv.Atoi(id)
	if err != nil {
		return book, errors.New("invalid book_id")
	}

	err = db.QueryRow("SELECT * FROM books WHERE book_id=$1", bookID).Scan(&book.BookID, &book.Tittle, &book.Author, &book.PublishingYear, &book.Genre, &book.AvailableCopy)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("no user found")
		}

		return book, err
	}

	_, err = db.Exec("DELETE FROM books WHERE book_id = $1", bookID)
	if err != nil {
		return book, err
	}

	return book, nil

}

// func Delete(db *sql.DB, id string) (User, error) {
// 	var user User

// 	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Password)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return user, errors.New("no user found")
// 		}

// 		return user, err
// 	}

// 	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

func Update(db *sql.DB, id string, b *Book) (Book, error) {
	var book Book

	_, err := db.Exec("UPDATE books SET title = $1, author = $2, publishing_year=$3,genre=$4,available_copy=$5 WHERE book_id = $6", b.Tittle, b.Author, b.PublishingYear, b.Genre, b.AvailableCopy, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("no user found")
		}
		return book, err
	}

	return book, nil
}

// func Update(db *sql.DB, id string, b *Book) (Book, error) {
// 	var book Book

// 	_, err := db.Exec("UPDATE users SET username = $1, password = $2 WHERE id = $3", u.Username, u.Password, id)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return user, errors.New("no user found")
// 		}
// 		return user, err
// 	}

// 	return user, nil
// }

func GetOneBook(db *sql.DB, id string) (Book, error) {
	var book Book

	err := db.QueryRow("SELECT * FROM books WHERE book_id = $1", id).Scan(&book.BookID, &book.Tittle, &book.Author, &book.PublishingYear, &book.Genre, &book.AvailableCopy)

	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("no user found")
		}
		return book, err
	}

	return book, nil
}

// func GetOneUser(db *sql.DB, id string) (User, error) {
// 	var user User

// 	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Password)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return user, errors.New("no user found")
// 		}
// 		return user, err
// 	}

// 	return user, nil
// }
