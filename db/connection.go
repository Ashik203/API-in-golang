package db

import (
	"app/config"
	"fmt"
	"log"
)

func ConnectDb() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("can noy load config")
	}

	db, err := ConnDb(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

}
