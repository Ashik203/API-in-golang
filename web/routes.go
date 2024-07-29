package web

import (
	"app/web/user"
	"database/sql"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router, dbase *sql.DB) {
	//router.HandleFunc("/users", user.GetUsers(dbase)).Methods("GET")
	router.HandleFunc("/users/{book_id}", user.GetOneBook(dbase)).Methods("GET")
	// router.HandleFunc("/users", user.CreateUser(dbase)).Methods("POST")
	router.HandleFunc("/users/{book_id}", user.DeleteBook(dbase)).Methods("DELETE")
	router.HandleFunc("/users/{book_id}", user.UpdateBook(dbase)).Methods("PUT")
	router.HandleFunc("/users", user.GetBooksPaginated(dbase)).Methods("GET")
	router.HandleFunc("/users", user.CreateBook(dbase)).Methods("POST")
}
