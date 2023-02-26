package handler

import (
	"datamerge/internal/service"
	"encoding/json"
	"net/http"
)

type HotelHandler struct {
	service service.IHotelService
}

func NewHotelHandler(service service.IHotelService) *HotelHandler {
	return &HotelHandler{
		service: service,
	}
}

func (h *HotelHandler) GetAllHotels(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(200)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	hotels, err := h.service.GetHotels()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotels)
}
