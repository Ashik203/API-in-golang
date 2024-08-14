package handlers

import (
	"app/db"
	"app/logger"
	"app/web/utils"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request to delete book.")

	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("failed to get id for delete book", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": id,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error(), id)
		return
	}

	var book db.Book
	book, err = db.Delete(id)
	if err != nil {
		slog.Error("Can't Delete Book", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": book,
		}))
		utils.SendError(w, http.StatusInternalServerError, err.Error(), book)
		return
	}

	utils.SendData(w, book)
}
