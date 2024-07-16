package web

import (
	"app/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer() {
	cfg, err := config.LoadConfig("/home/ashikurrahman/API/config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := config.ConnDb(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := mux.NewRouter()
	InitRoutes(router, db)

	formatString := fmt.Sprintf(":%d", cfg.Port)
	log.Fatal(http.ListenAndServe(formatString, router))
}
