package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/alonsofritz/cep-service-go/internal/usecase"
)

type AddressHandler struct {
	uc *usecase.AddressUseCase
}

func NewAddressHandler(uc *usecase.AddressUseCase) *AddressHandler {
	return &AddressHandler{uc: uc}
}

func (h *AddressHandler) GetAddressHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP is required", http.StatusBadRequest)
		return
	}
	addr, err := h.uc.GetAddress(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addr)
}
