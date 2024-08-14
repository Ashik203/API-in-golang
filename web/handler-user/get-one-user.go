package handleruser

import (
	"app/db"
	"app/logger"
	"app/web/utils"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request to get user by id.")

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

	user, err := db.ReadOneUser(id)
	if err != nil {
		slog.Error("failed to get book by id", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user,
		}))
		utils.SendError(w, http.StatusInternalServerError, err.Error(), user)
		return
	}

	utils.SendData(w, user)
}
