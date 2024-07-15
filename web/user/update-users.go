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

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request to update user.")

		var u controller.User
		json.NewDecoder(r.Body).Decode(&u)

		vars := mux.Vars(r)
		id := vars["id"]

		controller.Update(db, id, &u)

		utils.SendData(w, u)
	}
}
