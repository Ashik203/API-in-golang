package handlers

import (
	"app/db"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	page, _ := strconv.Atoi(queryParams.Get("page"))
	limit, _ := strconv.Atoi(queryParams.Get("limit"))

	filter := r.URL.Query().Get("filter")
	filterValue := r.URL.Query().Get("filterValue")
	
	search := r.URL.Query().Get("search")
	searchValue := r.URL.Query().Get("searchValue")

	sortBy := r.URL.Query().Get("sortBy")
	if sortBy == "" {
		sortBy = "book_id"
	}
	sortKey := r.URL.Query().Get("sortKey")
	if sortKey == "" {
		sortKey = "desc"
	}

	paginatedBooks, err := db.ReadBooks(page, limit, sortBy, sortKey, filter, filterValue, search, searchValue)
	if err != nil {
		http.Error(w, "Failed to get books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginatedBooks)
}
