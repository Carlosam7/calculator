package handlers

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error errorBody `json:"error"`
}

type errorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeError(w http.ResponseWriter, status int, code string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	req := errorResponse{
		Error: errorBody{
			Code:    code,
			Message: message,
		},
	}

	_ = json.NewEncoder(w).Encode(req)
}
