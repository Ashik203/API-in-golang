package handleruser

import (
	"app/db"
	"app/logger"
	"app/web/middlerware"
	"app/web/utils"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user db.User
	var hash string

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		slog.Error("failed to decode", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error(), user)
		return
	}

	if user.Username == "" {
		http.Error(w, "enter username", http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		http.Error(w, "enter password", http.StatusBadRequest)
		return
	}

	query, args, err := db.GetQueryBuilder().Select("password").From("users").Where(squirrel.Eq{"username": user.Username}).ToSql()
	if err != nil {
		slog.Error("failed to build query", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": query,
		}))
		utils.SendError(w, http.StatusInternalServerError, err.Error(), query)
		return
	}

	err = db.WriteDb.QueryRow(query, args...).Scan(&hash)
	if err != nil {
		slog.Error("error selecting password", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user,
		}))
		utils.SendError(w, http.StatusInternalServerError, err.Error(), user)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password))
	if err != nil {
		slog.Error("wrong password", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": hash,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error(), hash)
		return

	} else {
		fmt.Fprintln(w, "Successfully logged in")
	}
	
	expirationTime := time.Now().Add(5 * time.Minute).Unix()

	claims := middlerware.Claims{
		Username: user.Username,
		Exp:      expirationTime,
	}

	token := middlerware.CreateJwt(claims)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(5 * time.Minute),
	})

	fmt.Fprintln(w, "jwt token is: ", token)
	middlerware.LatestToken = token

}
