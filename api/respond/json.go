package respond

import (
	"encoding/json"
	"net/http"
)

func WithJSON(w http.ResponseWriter, v interface{}) {
	enc := json.NewEncoder(w)
	if err := enc.Encode(v); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
	}
}

type ErrorResponse struct {

	// message
	// Example: missing search field
	Message string `json:"message,omitempty"`

	// status code
	// Example: 400
	StatusCode int64 `json:"statusCode,omitempty"`
}

func WithError(w http.ResponseWriter, msg string, status int) {
	enc := json.NewEncoder(w)
	w.WriteHeader(status)
	err := enc.Encode(ErrorResponse{
		Message:    msg,
		StatusCode: int64(status),
	})
	if err != nil {
		http.Error(w, msg, status)
	}
}
