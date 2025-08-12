package main

import (
	"log"
	"net/http"
	"os"
	"prf-manager/db"
	repository "prf-manager/infrastructure"
	"prf-manager/interfaces/handler"
	"prf-manager/interfaces/usecase"
	route "prf-manager/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	gdb, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Database connection successfully")
	}

	err = db.Migrate(gdb)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	} else {
		log.Println("Database migration  successfully")
	}
	userRepo := repository.NewUserRepository(gdb)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)
	agencyRepo := repository.NewAgencyRepository(gdb)
	agencyUseCase := usecase.NewAgencyUseCase(agencyRepo)
	agencyHandler := handler.NewAgencyHandler(agencyUseCase)
	wareHouseRepo := repository.NewWareHouseRepository(gdb)
	wareHouseUseCase := usecase.NewWareHouseUseCase(wareHouseRepo)
	wareHouseHandler := handler.NewWareHouseHandler(wareHouseUseCase)
	recordRepo := repository.NewRecordRepository(gdb)
	recordUseCase := usecase.NewRecordUseCase(recordRepo)
	recordHandler := handler.NewRecordHandler(recordUseCase)

	route := &route.Route{
		UserHandler:      userHandler,
		AgencyHandler:    agencyHandler,
		WareHouseHandler: wareHouseHandler,
		RecordHandler:    recordHandler,
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	http.ListenAndServe(":"+port, route.NewRouter())
	log.Println("Server started successfully")
}
