package service

type IHotelService interface {
	GetHotels() (interface{}, error)
}

type HotelService struct {
}

// NewHotelService is a factory function that returns a pointer to a concrete HotelService type that
// implements the IHotelService interface
func NewHotelService() *HotelService {
	return &HotelService{}
}

func (s *HotelService) GetHotels() (interface{}, error) {
	// Business logic to get all hotels
	return "test", nil
}
