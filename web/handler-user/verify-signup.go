package handleruser

import (
	"app/db"
	"app/logger"
	"app/web/utils"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type OTP struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

func VerifySignUp(w http.ResponseWriter, r *http.Request) {

	var otp OTP
	if err := json.NewDecoder(r.Body).Decode(&otp); err != nil {
		slog.Error("failed to get user data to verify", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": otp,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error(), otp)
		return
	}

	actualOtp := db.GetRedisClient().Get(context.Background(), otp.Email).Val()

	if actualOtp != otp.Otp {
		slog.Error("OTP doesnt match", logger.Extra(map[string]any{
			"payload": actualOtp,
		}))
		utils.SendError(w, http.StatusBadRequest, "OTP doesnt match", actualOtp)
		return
	}

	utils.SendData(w, "Succesfully verified")

}
