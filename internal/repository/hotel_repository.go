package repository

import "datamerge/internal/model"

type HotelRepository interface {
	GetHotelsByHotelId(hotelIds []string) []*model.Hotel
	GetHotelsByDestinationId(destinationId int) []*model.Hotel
	PreloadHotelData(url string)
	InsertHotel(hotel *model.Hotel)
}
