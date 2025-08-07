package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prf-manager/interfaces/input"
	"prf-manager/interfaces/usecase"
)

type UserHandler struct {
	u usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) *UserHandler {
	return &UserHandler{u: u}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := &input.CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.u.Create(r.Context(), req); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
	})
	fmt.Println("Received CreateUserRequest:", req)
}
