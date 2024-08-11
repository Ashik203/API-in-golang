package web

import (
	"app/config"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func RunServer(wg *sync.WaitGroup) {
	mux := http.NewServeMux()
	InitRoutes(mux)

	wg.Add(1)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	// db, err := db.ConnDb(cfg)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	// defer db.Close()

	// go func() {

	formatString := fmt.Sprintf(":%d", cfg.Port)

	log.Fatal(http.ListenAndServe(formatString, mux))
	defer wg.Done()

	// }()
}
