package web

import (
	"app/config"
	"app/db"
	"fmt"
	"log"
	"net/http"
)

func RunServer() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := db.ConnDb(cfg)
	if err != nil {
		fmt.Println("Can't read data from Config FILE")
		log.Fatal(err)
	}

	defer db.Close()

	formatString := fmt.Sprintf(":%d", cfg.Port)

	mux := http.NewServeMux()

	InitRoutes(mux)
	log.Fatal(http.ListenAndServe(formatString, mux))
}
