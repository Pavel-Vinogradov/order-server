package response

import (
	"encoding/json"
	"net/http"
)

type JSONResponse[T any] struct {
	Data T `json:"data"`
}

type ErrorsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteJSON[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := JSONResponse[T]{Data: data}
	_ = json.NewEncoder(w).Encode(resp)
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	WriteJSON(w, status, map[string]string{"error": msg})
}
