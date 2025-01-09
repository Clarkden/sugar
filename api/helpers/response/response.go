package response

import (
	"encoding/json"
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
