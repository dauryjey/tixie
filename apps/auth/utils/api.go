package utils

import (
	"auth/types"
	"encoding/json"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := types.ErrorResponse{
		Code:    statusCode,
		Message: message,
	}
	json.NewEncoder(w).Encode(errorResponse)
}
