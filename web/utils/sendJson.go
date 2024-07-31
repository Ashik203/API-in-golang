package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	str, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Error Sending JSON")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Message:Internal Server Error"}`))
		return
	}

	w.WriteHeader(status)
	w.Write(str)
}
