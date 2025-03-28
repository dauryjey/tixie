package utils

import (
	"auth/types"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func GetEnv(env string) string {
	val := os.Getenv(env)

	if val == "" {
		log.Fatalf("Missing env: %s", env)
	}

	return val
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := types.ErrorResponse{
		Code:    statusCode,
		Message: message,
	}
	json.NewEncoder(w).Encode(errorResponse)
}
