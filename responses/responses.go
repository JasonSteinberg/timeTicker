package responses

import (
	"encoding/json"
	"github.com/JasonSteinberg/timeTicker/structs"
	"net/http"
)

func ErrorResponder(w http.ResponseWriter, status int, error structs.ServerMessage) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func Responder(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}
