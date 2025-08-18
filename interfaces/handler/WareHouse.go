package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prf-manager/interfaces/input"
	"prf-manager/interfaces/usecase"
	"prf-manager/project"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WareHouseHandler struct {
	*project.HandlerProject
	u usecase.WareHouseUseCase
}

func NewWareHouseHandler(u usecase.WareHouseUseCase) *WareHouseHandler {
	return &WareHouseHandler{u: u}
}
func (h *WareHouseHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := &input.CreateWareHouseRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.u.CreateWareHouse(r.Context(), req); err != nil {
		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	h.JSON(w, map[string]interface{}{
		"message": "WareHouse created successfully",
	}, http.StatusCreated)
}

// get all
func (h *WareHouseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	wareHouse, err := h.u.GetAll(r.Context())
	if err != nil {
		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	if len(wareHouse) == 0 {
		h.JSONError(w, fmt.Errorf("no WareHouse found"), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	h.JSON(w, wareHouse, http.StatusOK)
}

// update kho
func (h *WareHouseHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.JSONError(w, err, http.StatusBadRequest)
		return
	}

	req := &input.UpdateWareHouseRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.u.UpdateWareHouse(r.Context(), uint(id), req); err != nil {
		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	h.JSON(w, map[string]interface{}{
		"message": "WareHouse updated successfully",
	}, http.StatusOK)
}

func (h *WareHouseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {

		h.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.u.DeleteWareHouse(r.Context(), uint(id)); err != nil {

		h.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	h.JSON(w, map[string]interface{}{
		"message": "WareHouse deleted successfully",
	}, http.StatusOK)
}
