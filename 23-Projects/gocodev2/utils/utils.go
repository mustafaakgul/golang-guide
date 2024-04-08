package utils

import (
	"encoding/json"
	"net/http"

	"github.com/cihanozhan/models"
)

func SendError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	// fmt.Println(data)
	json.NewEncoder(w).Encode(data)
}
