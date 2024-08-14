package db

import (
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var psql squirrel.StatementBuilderType

func InitQueryBuilder(writeDb *sqlx.DB) {
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	createTables(writeDb)
}

func createTables(writeDb *sqlx.DB) {
	usersTable := `
	CREATE TABLE IF NOT EXISTS
	 users (user_id SERIAL primary key,username TEXT, email TEXT, password TEXT);`

	_, err := writeDb.Exec(usersTable)
	if err != nil {
		log.Printf("Failed to create table 'users': %v", err)
	}

	booksTable := `CREATE TABLE IF NOT EXISTS 
	books(book_id SERIAL primary key,title varchar(255) ,author varchar(255) , publishing_year int,genre varchar(255),available_copy int);`

	_, err1 := writeDb.Exec(booksTable)
	if err1 != nil {
		log.Printf("Failed to create table 'books': %v", err1)
	}
}

func GetQueryBuilder() squirrel.StatementBuilderType {
	return psql
}
