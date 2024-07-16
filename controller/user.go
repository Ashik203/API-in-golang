package controller

import (
	"database/sql"
	"errors"
	"log"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUsers(db *sql.DB) ([]User, error) {
	users := []User{}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		if err == sql.ErrNoRows {
			return users, errors.New("no user found")
		}
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func Add(db *sql.DB, u *User) error {
	err := db.QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id", u.Username, u.Password).Scan(&u.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("no user found")
		}

		return err
	}

	return nil
}

func Delete(db *sql.DB, id string) (User, error) {
	var user User

	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("no user found")
		}

		return user, err
	}

	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func Update(db *sql.DB, id string, u *User) (User, error) {
	var user User

	_, err := db.Exec("UPDATE users SET username = $1, password = $2 WHERE id = $3", u.Username, u.Password, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("no user found")
		}
		return user, err
	}

	return user, nil
}

func GetOneUser(db *sql.DB, id string) (User, error) {
	var user User

	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("no user found")
		}
		return user, err
	}

	return user, nil
}
