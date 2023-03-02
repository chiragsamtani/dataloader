package repository

import (
	"datamerge/internal/model"
	"sync"
)

// InMemoryHotelRepository uses one hashmaps to store keys that
// are indexed by hotelId only. Filtering by destinationIds will
// be using a map<int, map<string, hotel>>
type InMemoryHotelRepository struct {
	kvStore            map[string]*model.Hotel
	destinationIdStore map[int]map[string]*model.Hotel
	mu                 sync.Mutex
}

func NewInMemoryHotelRepository() *InMemoryHotelRepository {
	m := make(map[string]*model.Hotel, 0)
	destinationIdStore := make(map[int]map[string]*model.Hotel, 0)
	return &InMemoryHotelRepository{
		kvStore:            m,
		destinationIdStore: destinationIdStore,
	}
}

// GetHotelsByHotelIds returns all the hotels given a list of hotelIds
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
	var hotelResult []*model.Hotel
	allHotelsByDestinationId, _ := i.destinationIdStore[destinationId]
	for _, val := range allHotelsByDestinationId {
		hotelResult = append(hotelResult, val)
	}
	return hotelResult
}

// InsertHotel inserts a hotel with the keys being both the hotelId and destinationId
// this function is thread-safe
func (i *InMemoryHotelRepository) InsertHotel(hotel *model.Hotel) {
	i.mu.Lock()
	defer i.mu.Unlock()
	hotelIdKey := hotel.ID
	i.kvStore[hotelIdKey] = hotel
	// we will have to check if map of map exist to guard against
	// null value assignments i.e assigning to a null or empty map of maps
	mapForDestinationId, present := i.destinationIdStore[hotel.DestinationID]
	if present {
		// update the object of hotelIds for this particular destinationId
		mapForDestinationId[hotelIdKey] = hotel
	} else {
		// otherwise, insert a new map for this destinationId, with hotelId: hotel mapping
		i.destinationIdStore[hotel.DestinationID] = map[string]*model.Hotel{hotelIdKey: hotel}
	}
}
