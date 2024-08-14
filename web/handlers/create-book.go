package handlers

import (
	"app/db"
	"app/logger"
	"app/web/utils"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("Requesting to create book.")
	var book db.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		slog.Error("failed to decode header for create book", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": book,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error(), book)
		return
	}

	err = db.AddBook(&book)
	if err != nil {
		slog.Error("failed to create book", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": book,
		}))
		utils.SendError(w, http.StatusInternalServerError, err.Error(), book)
		return
	}

	utils.SendData(w, book)
}
