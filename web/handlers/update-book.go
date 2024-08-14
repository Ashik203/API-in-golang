package handlers

import (
	"app/db"
	"app/logger"
	"app/web/utils"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request to update book.")

	var book db.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		slog.Error("Can't decode body for updatebook", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": book,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error(), book)
		return
	}

	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("Can't get id", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": id,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error(), id)
		return
	}

	book, err = db.Update(id, &book)
	if err != nil {
		slog.Error("Can't update book", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": book,
		}))
		utils.SendError(w, http.StatusInternalServerError, err.Error(), book)
		return

	}

	utils.SendData(w, book)
}
