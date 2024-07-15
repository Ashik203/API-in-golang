package utils

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	str, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Message:Internal Server Error"}`))
		return
	}

	w.WriteHeader(status)
	w.Write(str)
}
