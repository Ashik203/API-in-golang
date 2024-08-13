package handlers

import (
	"app/db"
	"app/web/jwt"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var user db.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	var hash string

	db.InitQueryBuilder()
	query, args, err := db.GetQueryBuilder().Select("password").From("users").Where(squirrel.Eq{"username": user.Username}).ToSql()
	if err != nil {
		fmt.Println("can not build query ")
	}
	err1 := db.Db.QueryRow(query, args...).Scan(&hash)
	if err1 != nil {
		fmt.Fprintln(w, "error selecting password")
	}
	err2 := bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password))

	if err2 != nil {
		fmt.Fprintln(w, "wrong password")

	} else {
		fmt.Fprintln(w, "Successfully logged in")
	}
	expirationTime := time.Now().Add(5 * time.Minute).Unix()

	claims := jwt.Claims{
		Username: user.Username,
		Exp:      expirationTime,
	}

	token := jwt.CreateJwt(claims)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(5 * time.Minute),
	})

	fmt.Fprintln(w, "jwt token is: ", token)

	jwt.LatestToken = token

}
