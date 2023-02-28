package main

import (
	handlers "datamerge/internal/handler"
	"datamerge/internal/repository"
	service "datamerge/internal/service"
	"log"
	"net/http"
)

func main() {

	//logger := utils.NewLogger("warn")

	repo := repository.NewInMemoryHotelRepository()
	svc := service.NewHotelService(repo)
	handler := handlers.NewHotelHandler(svc)

	http.HandleFunc("/hotels", handler.SearchHotels)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
