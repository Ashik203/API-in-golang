package web

import (
	"app/web/handlers"
	"app/web/jwt"
	"net/http"
)

func InitRoutes(mux *http.ServeMux) {

	mux.Handle("GET /users", http.HandlerFunc(handlers.GetBooks))
	mux.Handle("POST /books", http.HandlerFunc(handlers.CreateBook))

	mux.Handle("GET /books/{book_id}", http.HandlerFunc(handlers.GetOneBook))
	mux.Handle("PUT /users/{book_id}", http.HandlerFunc(handlers.UpdateBook))
	// mux.Handle("DELETE /users/{book_id}",http.HandlerFunc(handlers.DeleteBook))

	mux.Handle("DELETE /users/{book_id}", jwt.JwtMiddleware(http.HandlerFunc(handlers.DeleteBook)))

	mux.Handle("POST /users", http.HandlerFunc(handlers.CreateUser))
	mux.Handle("GET /users/{user_id}", http.HandlerFunc(handlers.GetOneUser))
	mux.Handle("POST /userlogin", http.HandlerFunc(handlers.UserLogin))

}
