package model

type HotelRequestDTO struct {
	HotelId       []string `json:"hotel_ids"`
	DestinationId string   `json:"destination_id"`
}
