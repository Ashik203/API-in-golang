package user

import (
	"app/controller"
	"app/web/utils"
	"database/sql"
	"log"
	"net/http"
)

func GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request to get all users.")

		users, _ := controller.GetUsers(db)

		utils.SendData(w, users)
	}
}
