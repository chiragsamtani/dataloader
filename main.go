package main

import (
	config "datamerge/internal/config"
	handlers "datamerge/internal/handler"
	"datamerge/internal/repository"
	service "datamerge/internal/service"
	"datamerge/internal/utils"
	"log"
	"net/http"
)

func main() {
	config, err := config.GetConfigFromEnv()
	if err != nil {
		panic("unable to read config, please verify app.<env>.env exists")
	}
	logger := utils.NewLogger(config.GetLogLevel())

	repo := repository.NewInMemoryHotelRepository()

	dataLoaderService := service.NewDirectDataLoaderService(config.GetSupplierConfig(), repo, logger)
	dataLoaderService.LoadData()

	svc := service.NewHotelService(repo)
	hotelHandler := handlers.NewHotelHandler(svc)

	hotelHandler.SetupHandlers()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
