package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prf-manager/interfaces/input"
	usecase "prf-manager/interfaces/usecase"
	"prf-manager/project"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type AgencyHandler struct {
	*project.HandlerProject
	u usecase.AgencyUseCase
}

func NewAgencyHandler(u usecase.AgencyUseCase) *AgencyHandler {
	return &AgencyHandler{u: u}
}

func (h *AgencyHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := &input.CreateAgencyRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		h.JSONError(w, err, http.StatusBadRequest)
		return

	}

	if err := h.u.CreateAgency(r.Context(), req); err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"message": "Agency created successfully",
	// })
	h.JSON(w, map[string]interface{}{
		"message": "Agency created successfully",
	}, http.StatusOK)

	// h.JSON(w, req, http.StatusCreated)
}

func (h *AgencyHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	agency, err := h.u.GetAll(r.Context())
	if err != nil {
		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(agency) == 0 {
		h.JSONError(w, fmt.Errorf("no agency found"), http.StatusNotFound)
		return
	}

	h.JSON(w, agency, http.StatusOK)
}

func (h *AgencyHandler) Update(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)

	fmt.Println("ID:", idInt)
	if err != nil {
		// http.Error(w, "Invalid Agency ID", http.StatusBadRequest)
		h.JSONError(w, err, http.StatusBadRequest)

		return
	}
	agencyId := uint(idInt)
	req := &input.UpdateAgencyRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		h.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.u.UpdateAgency(r.Context(), agencyId, req); err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"message": "Agency updated successfully",
	// })
	h.JSON(w, map[string]interface{}{
		"message": "Agency updated successfully",
	}, http.StatusOK)
}
func (h *AgencyHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		// http.Error(w, "Invalid Agency ID", http.StatusBadRequest)
		h.JSONError(w, err, http.StatusBadRequest)
		return
	}
	agencyId := uint(idInt)

	if err := h.u.DeleteAgency(r.Context(), agencyId); err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// "message": "Agency deleted successfully",
	// })
	h.JSON(w, map[string]interface{}{
		"message": "Agency deleted successfully",
	}, http.StatusOK)
}
