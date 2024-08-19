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

func GetOneBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request to get book by id.")

	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("Can't get user id for get book by id", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": id,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error(), id)
		return
	}

	user, err := db.GetBookRepo().GetBook(id)
	if err != nil {
		slog.Error("Can't get book by id", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user,
		}))
		utils.SendError(w, http.StatusInternalServerError, err.Error(), user)
		return
	}

	utils.SendData(w, user)
}
