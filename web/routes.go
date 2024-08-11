package web

import (
	"app/web/handlers"
	"net/http"
)

func InitRoutes(mux *http.ServeMux) {

	mux.Handle("GET /users", http.HandlerFunc(handlers.GetBooks))
	mux.Handle("POST /users", http.HandlerFunc(handlers.CreateBook))
	mux.Handle("GET /users/{book_id}", http.HandlerFunc(handlers.GetOneBook))
	mux.Handle("PUT /users/{book_id}", http.HandlerFunc(handlers.UpdateBook))
	mux.Handle("DELETE /users/{book_id}", http.HandlerFunc(handlers.DeleteBook))

}
