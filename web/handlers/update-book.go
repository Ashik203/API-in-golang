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

	var book db.Books
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		slog.Error("Can't decode body for updatebook", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": book,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error(), book)
		return
	}

	if err := utils.Validator(book); err != nil {
		slog.Error("Failed to validate new book data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": book,
		}))
		utils.SendError(w, http.StatusExpectationFailed, err.Error(), err)
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

	updatedBook, err := db.GetBookRepo().Update(id, &book)
	if err != nil {
		slog.Error("Can't update book", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": updatedBook,
		}))
		utils.SendError(w, http.StatusInternalServerError, err.Error(), updatedBook)
		return

	}

	utils.SendData(w, book)
}
