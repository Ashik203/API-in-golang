package handlers

import (
	"app/db"
	"app/web/utils"
	"encoding/json"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	defaultSortBy := "book_id"
	defaultSortOrder := "asc"

	paginationParams := utils.GetPaginationParams(r, defaultSortBy, defaultSortOrder)

	paginatedBooks, err := db.ReadBooks(paginationParams.Page, paginationParams.Limit, paginationParams.SortBy, paginationParams.SortOrder, paginationParams.FilterBy, paginationParams.FilterValue, paginationParams.SearchBy, paginationParams.SearchValue)
	if err != nil {
		http.Error(w, "Failed to get books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginatedBooks)

	// queryParams := r.URL.Query()

	// page, _ := strconv.Atoi(queryParams.Get("page"))
	// limit, _ := strconv.Atoi(queryParams.Get("limit"))

	// filter := r.URL.Query().Get("filter")
	// filterValue := r.URL.Query().Get("filterValue")

	// search := r.URL.Query().Get("search")
	// searchValue := r.URL.Query().Get("searchValue")

	// sortBy := r.URL.Query().Get("sortBy")
	// if sortBy == "" {
	// 	sortBy = "book_id"
	// }
	// sortKey := r.URL.Query().Get("sortKey")
	// if sortKey == "" {
	// 	sortKey = "asc"
	// }

}
