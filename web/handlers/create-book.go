package handlers

import (
	"app/db"
	"app/web/utils"
	"encoding/json"
	"log"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("Requesting to create book.")

	var b db.Book

	json.NewDecoder(r.Body).Decode(&b)

	db.AddBook(&b)

	utils.SendData(w, b)
}
