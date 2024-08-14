package db

import "github.com/jmoiron/sqlx"

var (
	readDb  *sqlx.DB
	WriteDb *sqlx.DB
)

func GetWriteDB() *sqlx.DB {
	return WriteDb
}

func GetReadDB() *sqlx.DB {
	return readDb
}
