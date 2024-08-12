package handlers

import (
	"app/db"
	"app/web/utils"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("Requesting to create user.")
	var b db.User

	json.NewDecoder(r.Body).Decode(&b)

	db.AddUser(&b)

	utils.SendData(w, b)
}
