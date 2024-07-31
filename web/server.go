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
		fmt.Println("Can't read data from Config FILE")
		log.Fatal(err)
	}

	defer db.Close()

	formatString := fmt.Sprintf(":%d", cfg.Port)

	mux := mux.NewRouter()
	InitRoutes(mux)

	err = http.ListenAndServe(formatString, mux)

	if err != nil {
		log.Printf("Server failed to start: %v", err)
		log.Fatal(err)
	}

}
