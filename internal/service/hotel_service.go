package service

import (
	"datamerge/internal/model"
	"datamerge/internal/repository"
)

type IHotelService interface {
	SearchHotelsByHotelId(hotelId []string) (interface{}, error)
	SearchHotelsByDestinationId(destinationId int) (interface{}, error)
}

type HotelService struct {
	repository repository.HotelRepository
}

// NewHotelService is a factory function that returns a pointer to a concrete HotelService type that
// implements the IHotelService interface
func NewHotelService(hotelRepository repository.HotelRepository) *HotelService {
	return &HotelService{
		repository: hotelRepository,
	}
}

// SearchHotelsByHotelId will search for every hotel that matches the list
// of hotel ids.
// The response object will be a list of hotels, if there are no entries
// an empty list will be returned. An interface is used as the return value
// for modularity and ease of testing
func (s *HotelService) SearchHotelsByHotelId(hotelId []string) (interface{}, error) {
	hotels := s.repository.GetHotelsByHotelId(hotelId)
	if hotels == nil || len(hotels) == 0 {
		return []*model.Hotel{}, nil
	}
	return hotels, nil
}

// SearchHotelsByDestinationId will search for every hotel that matches the destinationId
// The response object will be a list of hotels. If no hotels exist under that destinationId
// an empty list will be returned. Similarly, an interface is returned
func (s *HotelService) SearchHotelsByDestinationId(destinationId int) (interface{}, error) {
	hotel := s.repository.GetHotelsByDestinationId(destinationId)
	if hotel == nil {
		return []*model.Hotel{}, nil
	}
	return hotel, nil
}
