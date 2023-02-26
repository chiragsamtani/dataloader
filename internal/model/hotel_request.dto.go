package model

// HotelRequestDTO is the parameter that the user can specify
// when searching for hotels
// Supported parameters are:
// List of HotelIds
// Single DestinationId
type HotelRequestDTO struct {
	HotelId       []string `json:"hotel_ids"`
	DestinationId string   `json:"destination_id"`
}
