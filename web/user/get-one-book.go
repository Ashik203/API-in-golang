package user

import (
	"app/controller"
	"app/web/utils"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetOneBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request to get book by id.")

		vars := mux.Vars(r)
		id := vars["book_id"]

		us, _ := controller.GetOneBook(db, id)

		utils.SendData(w, us)
	}
}
