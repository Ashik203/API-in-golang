package web

import (
	"app/controller"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func RunServer() {
	config, err := controller.LoadConfig("/home/ashikurrahman/API/config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	formatString := fmt.Sprintf(":%d", config.Port)

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username TEXT, password TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	InitRoutes(router, db)

	log.Fatal(http.ListenAndServe(formatString, router))
}
