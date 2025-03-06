package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ParseJSON(r *http.Request, value any) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("Content-Type header is not json")
	}
	decode := json.NewDecoder(r.Body)
	return decode.Decode(value)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Write([]byte("{\"error\":\"" + err.Error() + "\"}"))
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		log.Println("something is wrong with your code, fix it")
	}
	w.WriteHeader(status)
	w.Write(jsonData)
}
