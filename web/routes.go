package web

import (
	"app/web/handlers"
	"app/web/jwt"
	"net/http"
)

func InitRoutes(mux *http.ServeMux) {
	mux.Handle("POST /users", http.HandlerFunc(handlers.SignUp))
	mux.Handle("POST /userlogin", http.HandlerFunc(handlers.Login))
	mux.Handle("GET /books", http.HandlerFunc(handlers.GetBooks))
	mux.Handle("GET /books/{book_id}", http.HandlerFunc(handlers.GetOneBook))

	mux.Handle("GET /users/{user_id}", jwt.JwtMiddleware(http.HandlerFunc(handlers.GetOneUser)))
	mux.Handle("POST /books", jwt.JwtMiddleware(http.HandlerFunc(handlers.CreateBook)))
	mux.Handle("PUT /users/{book_id}", jwt.JwtMiddleware(http.HandlerFunc(handlers.UpdateBook)))
	mux.Handle("DELETE /users/{book_id}", jwt.JwtMiddleware(http.HandlerFunc(handlers.DeleteBook)))
}
