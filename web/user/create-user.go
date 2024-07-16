package user

import (
	"app/controller"
	"app/web/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request to add user.")

		var u controller.User

		json.NewDecoder(r.Body).Decode(&u)

		controller.Add(db, &u)

		utils.SendData(w, u)
	}
}
