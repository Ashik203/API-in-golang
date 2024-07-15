package user

import (
	"app/controller"
	"app/web/utils"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetOneUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request to get user by id.")

		vars := mux.Vars(r)
		id := vars["id"]

		us, _ := controller.GetOneUser(db, id)

		utils.SendData(w, us)
	}
}
