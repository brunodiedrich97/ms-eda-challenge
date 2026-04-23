package web

import (
	"encoding/json"
	"net/http"

	"github.com.br/brunodiedrich97/ms-balance/internal/usecase/get_balance"
	"github.com/go-chi/chi"
)

type WebBalanceHandler struct {
	GetBalanceUseCase get_balance.GetBalanceUseCase
}

func NewWebBalanceHandler(getBalanceUseCase get_balance.GetBalanceUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		GetBalanceUseCase: getBalanceUseCase,
	}
}

func (h *WebBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "account_id")
	if accountID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.GetBalanceUseCase.Execute(accountID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
