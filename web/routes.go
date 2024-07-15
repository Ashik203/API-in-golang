package web

import (
	"app/web/user"
	"database/sql"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router, dbase *sql.DB) {
	router.HandleFunc("/users", user.GetUsers(dbase)).Methods("GET")
	router.HandleFunc("/users/{id}", user.GetOneUser(dbase)).Methods("GET")
	router.HandleFunc("/users", user.CreateUser(dbase)).Methods("POST")
	router.HandleFunc("/users/{id}", user.DeleteUser(dbase)).Methods("DELETE")
	router.HandleFunc("/users/{id}", user.UpdateUser(dbase)).Methods("PUT")
}
