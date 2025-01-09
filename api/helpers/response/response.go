package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sugar/globals/types"
)

func BadRequest(w http.ResponseWriter, message string) {

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&types.ApiResponse{
		Code:    http.StatusBadRequest,
		Message: message,
		Success: false,
		Data:    nil,
	})
}

func InternalServerError(w http.ResponseWriter, err error, message string) {

	slog.Error(err.Error())

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&types.ApiResponse{
		Code:    http.StatusInternalServerError,
		Message: message,
		Success: false,
		Data:    nil,
	})
}

func Success(w http.ResponseWriter, message string, data interface{}) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&types.ApiResponse{
		Code:    http.StatusOK,
		Message: message,
		Success: true,
		Data:    data,
	})
}

func Unauthorized(w http.ResponseWriter, message string) {

	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&types.ApiResponse{
		Code:    http.StatusUnauthorized,
		Message: message,
		Success: true,
		Data:    nil,
	})
}

func Conflict(w http.ResponseWriter, message string) {

	w.WriteHeader(http.StatusConflict)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&types.ApiResponse{
		Code:    http.StatusConflict,
		Message: message,
		Success: true,
		Data:    nil,
	})
}
