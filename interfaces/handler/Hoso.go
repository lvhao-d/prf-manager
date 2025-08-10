package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"prf-manager/interfaces/input"
	usecase "prf-manager/interfaces/usecase"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type HoSoHandler struct {
	u usecase.HoSoUseCase
}

func NewHoSoHandler(u usecase.HoSoUseCase) *HoSoHandler {
	return &HoSoHandler{u: u}
}

func (h *HoSoHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := &input.CreateHoSoRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.u.CreateHoSo(r.Context(), req); err != nil {
		log.Println("p", req)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "HoSo created successfully",
	})
}
func (h *HoSoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid HoSo ID", http.StatusBadRequest)
		return
	}

	hoSo, err := h.u.GetHoSoByID(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if hoSo == nil {
		http.Error(w, "HoSo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hoSo)
}
func (h *HoSoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	hoSos, err := h.u.GetAllHoSo(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(hoSos) == 0 {
		http.Error(w, "No HoSos found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hoSos)
}
func (h *HoSoHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid HoSo ID", http.StatusBadRequest)
		return
	}

	req := &input.UpdateHoSoRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.u.UpdateHoSo(r.Context(), uint(id), req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "HoSo updated successfully",
	})
}
func (h *HoSoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid HoSo ID", http.StatusBadRequest)
		return
	}

	if err := h.u.DeleteHoSo(r.Context(), uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "HoSo deleted successfully",
	})
}
