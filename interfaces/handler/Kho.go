package handler

import (
	"encoding/json"
	"net/http"
	"prf-manager/interfaces/input"
	"prf-manager/interfaces/usecase"
)

type KhoHandler struct {
	u usecase.KhoUseCase
}

func NewKhoHandler(u usecase.KhoUseCase) *KhoHandler {
	return &KhoHandler{u: u}
}
func (h *KhoHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := &input.CreateKhoRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.u.CreateKho(r.Context(), req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Kho created successfully",
	})
}
func (h *KhoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	khos, err := h.u.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(khos) == 0 {
		http.Error(w, "No Khos found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(khos)
}
