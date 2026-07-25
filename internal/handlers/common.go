package handlers

import (
	"encoding/json"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
}

func (rw *responseWriter) writeJSON(status int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	_ = json.NewEncoder(rw).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.writeJSON(status, ErrorResponse{
		Error:   http.StatusText(status),
		Message: message,
	})
}
