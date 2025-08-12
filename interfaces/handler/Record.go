package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"prf-manager/interfaces/input"
	usecase "prf-manager/interfaces/usecase"
	"prf-manager/project"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RecordHandler struct {
	*project.HandlerProject
	u usecase.RecordUseCase
}

func NewRecordHandler(u usecase.RecordUseCase) *RecordHandler {
	return &RecordHandler{u: u}
}

func (h *RecordHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := &input.CreateRecordRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		h.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.u.CreateRecord(r.Context(), req); err != nil {
		log.Println("p", req)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"message": "Record created successfully",
	// })
	h.JSON(w, map[string]interface{}{
		"message": "Record created successfully",
	}, http.StatusCreated)
}
func (h *RecordHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// http.Error(w, "Invalid Record ID", http.StatusBadRequest)
		h.JSONError(w, err, http.StatusBadRequest)
		return
	}

	record, err := h.u.GetRecordByID(r.Context(), uint(id))
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	if record == nil {
		// http.Error(w, "Record not found", http.StatusNotFound)
		h.JSONError(w, err, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(record)
	h.JSON(w, record, http.StatusOK)
}
func (h *RecordHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	record, err := h.u.GetAllRecord(r.Context())
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	if len(record) == 0 {
		// http.Error(w, "No record found", http.StatusNotFound)
		h.JSONError(w, fmt.Errorf("no record found"), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(record)
	h.JSON(w, record, http.StatusOK)
}
func (h *RecordHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// http.Error(w, "Invalid Record ID", http.StatusBadRequest)
		h.JSONError(w, err, http.StatusBadRequest)
		return
	}

	req := &input.UpdateRecordRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		h.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.u.UpdateRecord(r.Context(), uint(id), req); err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"message": "Record updated successfully",
	// })
	h.JSON(w, map[string]interface{}{
		"message": "Record updated successfully",
	}, http.StatusOK)
}
func (h *RecordHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// http.Error(w, "Invalid Record ID", http.StatusBadRequest)
		h.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.u.DeleteRecord(r.Context(), uint(id)); err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"message": "Record deleted successfully",
	// })
	h.JSON(w, map[string]interface{}{
		"message": "Record deleted successfully",
	}, http.StatusOK)
}
