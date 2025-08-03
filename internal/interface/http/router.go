package router

import (
	gorminfra "prf-manager/internal/infrastructure/gorm"
	"prf-manager/internal/interface/http/handler"
	"prf-manager/internal/usecase/agency"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	agencyRepo := gorminfra.NewAgencyRepository(db)
	agencyUseCase := agency.NewUseCase(agencyRepo)
	agencyHandler := handler.NewAgencyHandler(agencyUseCase)
	api := r.Group("/api")
	{
		api.GET("/agencies", agencyHandler.GetAll)
	}
}
