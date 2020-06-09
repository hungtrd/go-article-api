package ulti

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseError struct {
	StatusCode int
	Message    string
}

type Error struct {
	Body []string `json:"body"`
}

type ResponseFormatError struct {
	Errors Error `json:"errors"`
}

func SendResponseError(w http.ResponseWriter, e ResponseError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.StatusCode)
	response := ResponseFormatError{
		Errors: Error{
			Body: []string{e.Message},
		},
	}

	jsonResponse, err := json.Marshal(response)
	log.Println("Error:", err)

	w.Write(jsonResponse)
}

func SendResponseData(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	response, _ := json.Marshal(payload)
	w.Write(response)
}
