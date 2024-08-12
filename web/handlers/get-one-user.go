package handlers

import (
	"app/db"
	"app/web/utils"
	"log"
	"net/http"
	"strconv"
)

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request to get user by username.")

	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	us, _ := db.ReadOneUser(id)
	utils.SendData(w, us)
}
