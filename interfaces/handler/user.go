package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prf-manager/interfaces/input"
	"prf-manager/interfaces/usecase"
	"prf-manager/project"
)

type UserHandler struct {
	*project.HandlerProject
	u usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) *UserHandler {
	return &UserHandler{u: u}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := &input.UserLoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	resp, err := h.u.Login(r.Context(), req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to login: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	h.JSON(w, resp, http.StatusOK)
}
