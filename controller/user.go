package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"os"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	Port int `json:"port"`
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
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Password); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
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
	var u User

	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Username, &u.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return u, errors.New("no user found")
		}

		return u, err
	}

	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return u, err
	}

	return u, nil
}

func Update(db *sql.DB, id string, u *User) (User, error) {
	var us User

	_, err := db.Exec("UPDATE users SET username = $1, password = $2 WHERE id = $3", u.Username, u.Password, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return us, errors.New("no user found")
		}
		return us, err
	}

	return us, nil
}

func GetOneUser(db *sql.DB, id string) (User, error) {
	var u User

	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Username, &u.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return u, errors.New("no user found")
		}
		return u, err
	}

	return u, nil
}

func LoadConfig(filename string) (Config, error) {
	var config Config

	configFile, err := os.Open(filename)
	if err != nil {
		return config, err
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}
