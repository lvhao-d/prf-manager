package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prf-manager/interfaces/input"
	usecase "prf-manager/interfaces/usecase"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CoQuanHandler struct {
	u usecase.CoQuanUseCase
}

func NewCoQuanHandler(u usecase.CoQuanUseCase) *CoQuanHandler {
	return &CoQuanHandler{u: u}
}

func (h *CoQuanHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := &input.CreateCoQuanRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	if err := h.u.CreateCoQuan(r.Context(), req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "CoQuan created successfully",
	})
}

func (h *CoQuanHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	coQuans, err := h.u.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(coQuans) == 0 {
		http.Error(w, "No CoQuans found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(coQuans)
}

func (h *CoQuanHandler) Update(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	fmt.Println("idStr =", idStr)

	idInt, err := strconv.Atoi(idStr)

	fmt.Println("ID:", idInt)
	if err != nil {
		http.Error(w, "Invalid CoQuan ID", http.StatusBadRequest)
		return
	}
	coquanId := uint(idInt)
	req := &input.UpdateCoQuanRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.u.UpdateCoQuan(r.Context(), coquanId, req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "CoQuan updated successfully",
	})
}
func (h *CoQuanHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid CoQuan ID", http.StatusBadRequest)
		return
	}
	coquanId := uint(idInt)

	if err := h.u.DeleteCoQuan(r.Context(), coquanId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "CoQuan deleted successfully",
	})
}
