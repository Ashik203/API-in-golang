package handlers

import (
	"app/db"
	"app/web/utils"
	"encoding/json"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	paginationParams := utils.GetPaginationParams(r)

	paginatedBooks, err := db.ReadBooks(paginationParams.Page, paginationParams.Limit, paginationParams.SortBy, paginationParams.SortOrder, paginationParams.FilterBy, paginationParams.FilterValue, paginationParams.SearchBy, paginationParams.SearchValue)
	if err != nil {
		http.Error(w, "Failed to get books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginatedBooks)
}
