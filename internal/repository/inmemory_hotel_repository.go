package repository

import (
	"datamerge/internal/model"
	"sync"
)

// InMemoryHotelRepository uses two hashmaps to store keys that
// are indexed by hotelId and destinationId, we use two hashmaps
// as one destinationIds can map to many hotels. We can use a single
// hashmap, but accepting the tradeoff that more memory will be used
// in return for faster lookups
type InMemoryHotelRepository struct {
	hotelIdKvStore       map[string]*model.Hotel
	destinationIdKvStore map[int][]*model.Hotel
	mu                   sync.Mutex
}

func NewInMemoryHotelRepository() *InMemoryHotelRepository {
	m := make(map[string]*model.Hotel, 0)
	md := make(map[int][]*model.Hotel, 0)
	return &InMemoryHotelRepository{
		hotelIdKvStore:       m,
		destinationIdKvStore: md,
	}
}

// GetHotelsByHotelId returns all the hotels given a list of hotelIds
// this function is thread-safe
func (i *InMemoryHotelRepository) GetHotelsByHotelId(hotelIds []string) []*model.Hotel {
	i.mu.Lock()
	defer i.mu.Unlock()
	var result []*model.Hotel
	for _, val := range hotelIds {
		hotel, present := i.hotelIdKvStore[val]
		if present {
			result = append(result, hotel)
		}
	}
	return result
}

// GetHotelsByDestinationId returns all the hotels given single destinationId
// this function is thread-safe
func (i *InMemoryHotelRepository) GetHotelsByDestinationId(destinationId int) []*model.Hotel {
	i.mu.Lock()
	defer i.mu.Unlock()
	hotel, present := i.destinationIdKvStore[destinationId]
	if present {
		return hotel
	}
	return nil
}

// InsertHotel inserts a hotel with the keys being both the hotelId and destinationId
// this function is thread-safe
func (i *InMemoryHotelRepository) InsertHotel(hotel *model.Hotel) {
	i.mu.Lock()
	hotelIdKey := hotel.ID
	destinationIdKey := hotel.DestinationID
	i.hotelIdKvStore[hotelIdKey] = hotel
	i.destinationIdKvStore[destinationIdKey] = append(i.destinationIdKvStore[destinationIdKey], hotel)
	defer i.mu.Unlock()
}
