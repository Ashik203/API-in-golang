package web

import (
	"app/web/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	r.HandleFunc("/users", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/users", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/users/{book_id}", handlers.GetOneBook).Methods("GET")
	r.HandleFunc("/users/{book_id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/users/{book_id}", handlers.DeleteBook).Methods("DELETE")

}
