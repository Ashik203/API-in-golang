package handlers

import (
	"app/db"
	"app/web/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request to delete book.")
	
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	var b db.Book
	b, err = db.Delete(id)
	if err != nil {
		fmt.Println("Can't Delete Book")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.SendData(w, b)
}
