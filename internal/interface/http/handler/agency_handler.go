package handler

import (
	"net/http"
	"prf-manager/internal/usecase/agency"

	"github.com/gin-gonic/gin"
)

type AgencyHandler struct {
	useCase agency.UseCase
}

func NewAgencyHandler(useCase *agency.UseCase) *AgencyHandler {
	return &AgencyHandler{useCase: *useCase}
}
func (h *AgencyHandler) GetAll(c *gin.Context) {
	agencies, err := h.useCase.GetAllAgencies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch agencies"})
		return
	}
	c.JSON(http.StatusOK, agencies)
}
