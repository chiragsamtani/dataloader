package repository

import (
	"datamerge/internal/model"
	"sync"
)

// InMemoryHotelRepository uses one hashmaps to store keys that
// are indexed by hotelId only. Filtering by destinationIds will
// be doing a full table scan currently
type InMemoryHotelRepository struct {
	kvStore map[string]*model.Hotel
	mu      sync.Mutex
}

func NewInMemoryHotelRepository() *InMemoryHotelRepository {
	m := make(map[string]*model.Hotel, 0)
	return &InMemoryHotelRepository{
		kvStore: m,
	}
}

// GetHotelsByHotelId returns all the hotels given a list of hotelIds
// this function is thread-safe
func (i *InMemoryHotelRepository) GetHotelsByHotelIds(hotelIds []string) []*model.Hotel {
	i.mu.Lock()
	defer i.mu.Unlock()
	var result []*model.Hotel
	for _, val := range hotelIds {
		hotel, present := i.kvStore[val]
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
	var hotels []*model.Hotel
	for _, val := range i.kvStore {
		if val.DestinationID == destinationId {
			hotels = append(hotels, val)
		}
	}
	return hotels
}

// InsertHotel inserts a hotel with the keys being both the hotelId and destinationId
// this function is thread-safe
func (i *InMemoryHotelRepository) InsertHotel(hotel *model.Hotel) {
	i.mu.Lock()
	hotelIdKey := hotel.ID
	i.kvStore[hotelIdKey] = hotel
	i.mu.Unlock()
}
