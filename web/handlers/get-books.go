package handlers

import (
	"app/db"
	"app/logger"
	"app/web/utils"
	"log/slog"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	paginationParams := utils.GetPaginationParams(r)

	paginatedBooks, err := db.ReadBooks(paginationParams.Page, paginationParams.Limit, paginationParams.SortBy, paginationParams.SortOrder, paginationParams.FilterBy, paginationParams.FilterValue, paginationParams.SearchBy, paginationParams.SearchValue)
	if err != nil {
		slog.Error("Can't get all books", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": paginatedBooks,
		}))
		utils.SendError(w, http.StatusInternalServerError, err.Error(), paginatedBooks)
		return
	}

	utils.SendData(w, paginatedBooks)
}
