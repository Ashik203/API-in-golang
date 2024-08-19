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
		slog.Error("failed to convert id", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": id,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error(), id)
		return
	}

	resultChan := make(chan *db.Books)
	errorChan := make(chan error)

	go func() {
		book, err := db.GetBookRepo().Delete(id)
		if err != nil {
			errorChan <- err
		}
		resultChan <- book

	}()
	select {
	case book := <-resultChan:
		utils.SendData(w, book)
	case err := <-errorChan:
		utils.SendError(w, http.StatusInternalServerError, err.Error(), err)
	}
}
