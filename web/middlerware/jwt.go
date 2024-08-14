package middlerware

import (
	"app/logger"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte("secretKey")
var LatestToken, WhoseToken string

type Claims struct {
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
}

type Header struct {
	Alg string `json:"alg,omitempty"`
	Typ string `json:"typ,omitempty"`
}

func CreateJwt(claims Claims) string {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	headerBytes, err := json.Marshal(header)
	if err != nil {
		slog.Error("Can not convert header", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": headerBytes,
		}))
	}

	claimsByte, err := json.Marshal(claims)
	if err != nil {
		slog.Error("Can not convert claims", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": claimsByte,
		}))

	}

	headerBase64 := base64.RawURLEncoding.EncodeToString(headerBytes)
	claimsBase64 := base64.RawURLEncoding.EncodeToString(claimsByte)

	signature := CreateSignature(headerBase64 + "." + claimsBase64)

	token := headerBase64 + "." + claimsBase64 + "." + signature

	return token
}

func CreateSignature(data string) string {
	h := hmac.New(sha256.New, JwtKey)
	h.Write([]byte(data))

	return base64.RawStdEncoding.EncodeToString(h.Sum(nil))
}

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie := r.Header.Get("token")

		if !ValidateJWT(w, cookie) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} else if cookie != LatestToken {
			fmt.Fprint(w, "Token doesn't match")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ValidateJWT(w http.ResponseWriter, token string) bool {

	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		fmt.Fprintln(w, "Format is incorrect")
		return false
	}

	headerBase64 := parts[0]
	claimsBase64 := parts[1]
	signature := parts[2]

	expectedSignature := CreateSignature(headerBase64 + "." + claimsBase64)

	if signature != expectedSignature {
		fmt.Fprintln(w, "Signature doesnt match")
		return false
	}

	claimsByte, err := base64.RawURLEncoding.DecodeString(claimsBase64)
	if err != nil {
		slog.Error("Can not decode string", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": claimsByte,
		}))

	}

	var claims Claims

	json.Unmarshal(claimsByte, &claims)

	if claims.Exp < time.Now().Unix() {
		fmt.Fprintln(w, "Token Expired")
		return false
	}

	WhoseToken = claims.Username
	return true
}

func JwtLogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenstr := r.Header.Get("token")
		if tokenstr == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(tokenstr, ".")
		if len(parts) != 3 {
			fmt.Fprintln(w, "Format is incorrect")
			return
		}

		headerBase64 := parts[0]
		claimsBase64 := parts[1]
		signature := parts[2]

		expectedSignature := CreateSignature(headerBase64 + "." + claimsBase64)
		if signature != expectedSignature {
			fmt.Fprintln(w, "Signature doesnt match")
			return
		}

		claimsByte, err := base64.RawURLEncoding.DecodeString(claimsBase64)
		if err != nil {
			slog.Error("Can not decode string", logger.Extra(map[string]any{
				"error":   err.Error(),
				"payload": claimsByte,
			}))

		}

		var claims Claims

		json.Unmarshal(claimsByte, &claims)

		if claims.Exp > time.Now().Unix() {
			claims.Exp = time.Now().Add(-10 * time.Minute).Unix()
			fmt.Fprintln(w, "The new expiration", claims.Exp)
			LatestToken = ""
			WhoseToken = ""
		}
		next.ServeHTTP(w, r)

	})

}
