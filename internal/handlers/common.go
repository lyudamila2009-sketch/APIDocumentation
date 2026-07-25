package handlers

import (
	"encoding/json"
	"net/http"

	"api-doc-example/internal/models"
)

type emptySlice []interface{}

// MarshalJSON returns an empty JSON array instead of null for nil slices
func (s emptySlice) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]interface{}(s))
}

type responseWriter struct {
	http.ResponseWriter
}

func (rw *responseWriter) writeJSON(status int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	_ = json.NewEncoder(rw).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	rw := &responseWriter{ResponseWriter: w}
	rw.writeJSON(status, models.ErrorResponse{
		Error:   http.StatusText(status),
		Message: message,
	})
}
