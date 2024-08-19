package web

import (
	"app/web/handlers"
	"app/web/middlerware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middlerware.Manager) {
	mux.Handle(
		"POST /users/signup",
		manager.With(
			http.HandlerFunc(handlers.SignUp),
		),
	)

	mux.Handle(
		"POST /users/verify",
		manager.With(
			http.HandlerFunc(handlers.VerifySignUp),
		),
	)

	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	)

	mux.Handle(
		"GET /books",
		manager.With(
			http.HandlerFunc(handlers.GetBooks),
		),
	)

	mux.Handle(
		"GET /books/{book_id}",
		manager.With(
			http.HandlerFunc(handlers.GetOneBook),
		),
	)

	mux.Handle(
		"GET /users/{user_id}",
		manager.With(
			http.HandlerFunc(handlers.GetOneUser),
			middlerware.JwtMiddleware,
		),
	)

	mux.Handle(
		"POST /books",
		manager.With(
			http.HandlerFunc(handlers.CreateBook),
			middlerware.JwtMiddleware,
		),
	)
	mux.Handle(
		"PUT /books/{book_id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateBook),
			middlerware.JwtMiddleware,
		),
	)

	mux.Handle(
		"DELETE /books/{book_id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteBook),
			middlerware.JwtMiddleware,
		),
	)
}
