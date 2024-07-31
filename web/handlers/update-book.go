package handlers

import (
	"app/db"
	"app/web/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request to update book.")

	var b db.Book
	json.NewDecoder(r.Body).Decode(&b)

	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	db.Update(id, &b)

	utils.SendData(w, b)
}
