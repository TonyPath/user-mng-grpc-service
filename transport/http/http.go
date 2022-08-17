package http

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, v any, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(v)
	_, _ = w.Write(b)
}
