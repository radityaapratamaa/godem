package transaction

import (
	"bcg-test/domain/models/transaction"
	"encoding/json"
	"net/http"
)

type purchase interface {
	CreateNew(w http.ResponseWriter, r *http.Request)
}

func (h *Handler) CreateNew(w http.ResponseWriter, r *http.Request) {
	var requestData []*transaction.Purchase
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		response := map[string]interface{}{
			"code":          400,
			"message":       err.Error(),
			"error_message": "Bad Request",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := h.transactionUc.CreateNew(r.Context(), requestData)
	if err != nil {
		response := map[string]interface{}{
			"code":          500,
			"error_message": "Internal Server Error",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]interface{}{
		"code":    200,
		"message": "SUCCESS",
		"data":    result,
	}
	json.NewEncoder(w).Encode(response)
}
