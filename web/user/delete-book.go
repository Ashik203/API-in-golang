package user

import (
	"app/controller"
	"app/web/utils"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request to delete book.")

		vars := mux.Vars(r)
		var b controller.Book
		id := vars["book_id"]

		b, err := controller.Delete(db, id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		utils.SendData(w, b)
	}
}
