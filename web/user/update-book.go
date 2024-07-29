package user

import (
	"app/controller"
	"app/web/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request to update book.")

		var b controller.Book
		json.NewDecoder(r.Body).Decode(&b)

		vars := mux.Vars(r)
		id := vars["book_id"]

		controller.Update(db, id, &b)

		utils.SendData(w, b)
	}
}
