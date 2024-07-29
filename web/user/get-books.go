package user

import (
	"app/controller"
	"encoding/json"
	"net/http"
	"strconv"

	"database/sql"
)

func GetBooksPaginated(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		page, _ := strconv.Atoi(queryParams.Get("page"))
		limit, _ := strconv.Atoi(queryParams.Get("limit"))

		users, err := controller.GetPaginatedBooks(db, page, limit)
		if err != nil {
			http.Error(w, "Failed to get users", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
