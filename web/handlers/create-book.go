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
	var newBook db.Books

	// decodeChan:=make(chan error)

	// go func ()  {
	// 	err := json.NewDecoder(r.Body).Decode(&newBook)
	// 	if err!=nil{

	// 		decodeChan<

	// 	}

	// }()

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		slog.Error("Failed to decode new user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newBook,
		}))
		utils.SendError(w, http.StatusPreconditionFailed, err.Error(), newBook)
		return
	}

	if err := utils.Validator(newBook); err != nil {
		slog.Error("Failed to validate new book data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newBook,
		}))
		utils.SendError(w, http.StatusExpectationFailed, err.Error(), err)
		return
	}

	var insertedBook *db.Books
	var err error

	if insertedBook, err = db.GetBookRepo().InsertBook(&newBook); err != nil {
		slog.Error("failed to create book", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newBook,
		}))

		utils.SendError(w, http.StatusInternalServerError, err.Error(), newBook)
		return
	}

	utils.SendData(w, insertedBook)
}
