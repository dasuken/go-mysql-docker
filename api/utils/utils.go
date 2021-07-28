package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func WriteAsJson(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	WriteAsJson(w, struct {
		Error string `json:"error"`
	}{Error: err.Error()})
}

func Debug(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(bytes))
	}
}
