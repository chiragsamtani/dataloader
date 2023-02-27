package repository

import "datamerge/internal/model"

type HotelRepository interface {
	GetHotelsByHotelId(hotelIds []string) []*model.Hotel
	GetHotelsByDestinationId(destinationId int) []*model.Hotel
	InsertHotel(hotel *model.Hotel)
}
