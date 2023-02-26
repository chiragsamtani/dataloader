package handler

import (
	"datamerge/internal/model"
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

func (h *HotelHandler) SearchHotels(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
	default:
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		sendErrorResponse(w, "Request body must be in JSON format", http.StatusBadRequest)
		return
	}

	var requestHotelDTO model.HotelRequestDTO
	err := json.NewDecoder(r.Body).Decode(&requestHotelDTO)
	if err != nil {
		sendErrorResponse(w, "Please specify at least one hotel ID(s) or a single destination ID", http.StatusBadRequest)
		return
	}

	if requestHotelDTO.HotelId != nil && len(requestHotelDTO.HotelId) > 0 {
		hotels, err := h.service.SearchHotelsByHotelId(requestHotelDTO.HotelId)
		if err != nil {
			sendErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hotels)
	} else if requestHotelDTO.DestinationId != 0 {
		hotels, err := h.service.SearchHotelsByDestinationId(requestHotelDTO.DestinationId)
		if err != nil {
			sendErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hotels)
	} else {
		sendErrorResponse(w, "Please specify at least one hotel ID(s) or a single destination ID", http.StatusBadRequest)
	}
}

func sendErrorResponse(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	errResp := model.ErrorResponse{
		Message: msg,
	}
	json.NewEncoder(w).Encode(errResp)
}
