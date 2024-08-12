package db

import (
	"app/config"
	"database/sql"
	"fmt"
	"log"
)

var Db *sql.DB

func ConnDb(cfg config.Config) (*sql.DB, error) {
	var err error
	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Db.Host, cfg.Db.Port, cfg.Db.User, cfg.Db.Pass, cfg.Db.Name)
	Db, err = sql.Open("postgres", dbConnStr)
	if err != nil {
		fmt.Println("Error in Initialing the POSTGRES")
		log.Fatal(err)


	}

	
	_, err = Db.Exec("CREATE TABLE IF NOT EXISTS users (user_id SERIAL primary key,username TEXT, email TEXT, password TEXT)")
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	_, err = Db.Exec("CREATE TABLE IF NOT EXISTS books(book_id SERIAL primary key,title varchar(255) ,author varchar(255) , publishing_year int,genre varchar(255),available_copy int)")
	if err != nil {
		log.Fatal("failed to create table")
	}

	return Db, err
}
