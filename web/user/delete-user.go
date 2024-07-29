package user

// import (
// 	"app/controller"
// 	"app/web/utils"
// 	"database/sql"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func DeleteUser(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Request to delete user.")

// 		vars := mux.Vars(r)
// 		var u controller.User
// 		id := vars["id"]

// 		u, err := controller.Delete(db, id)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusNotFound)
// 			return
// 		}

// 		utils.SendData(w, u)
// 	}
// }
