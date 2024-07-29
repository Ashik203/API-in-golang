package user

import (
	"app/controller"
	"app/web/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func CreateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requesting to create book.")

		var b controller.Book

		json.NewDecoder(r.Body).Decode(&b)

		controller.AddBook(db, &b)

		utils.SendData(w, b)

	}
}
