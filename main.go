package main

import (
	handlers "datamerge/internal/handler"
	service "datamerge/internal/service"
	"log"
	"net/http"
)

func main() {

	svc := service.NewHotelService()
	handler := handlers.NewHotelHandler(svc)

	http.HandleFunc("/hotels", handler.SearchHotels)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
