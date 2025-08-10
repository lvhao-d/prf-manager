package handler

import (
	"encoding/json"
	"net/http"
	"prf-manager/interfaces/input"
	"prf-manager/interfaces/usecase"
	"strconv"

	"github.com/go-chi/chi/v5"
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
func (h *KhoHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kho ID", http.StatusBadRequest)
		return
	}

	req := &input.UpdateKhoRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.u.UpdateKho(r.Context(), uint(id), req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Kho updated successfully",
	})
}

func (h *KhoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kho ID", http.StatusBadRequest)
		return
	}

	if err := h.u.DeleteKho(r.Context(), uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Kho deleted successfully",
	})
}
