package repository

import (
	"datamerge/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testSingleDestinationId   = 12
	testMultipleDestinationId = 14
	testHotelId1              = "slx4"
	testHotelId2              = "pz82"
	testHotelId3              = "x3cc"
)

// Testing data
var (
	hotelData1 = model.Hotel{
		ID:            testHotelId1,
		DestinationID: testSingleDestinationId,
		Name:          "Radisson",
	}
	hotelData2 = model.Hotel{
		ID:            testHotelId2,
		DestinationID: testMultipleDestinationId,
		Name:          "Holiday Inn",
	}
	hotelData3 = model.Hotel{
		ID:            testHotelId3,
		DestinationID: testMultipleDestinationId,
		Name:          "Holiday Inn",
	}
)

func prefilledTestingRepository() *InMemoryHotelRepository {
	repo := NewInMemoryHotelRepository()
	repo.InsertHotel(&hotelData1)
	repo.InsertHotel(&hotelData2)
	repo.InsertHotel(&hotelData3)
	return repo
}

func TestInMemoryHotelRepository_InsertOneElement(t *testing.T) {
	repo := NewInMemoryHotelRepository()
	repo.InsertHotel(&hotelData1)
	assert.Equal(t, *repo.hotelIdKvStore[testHotelId1], hotelData1)
	assert.Equal(t, repo.destinationIdKvStore[hotelData1.DestinationID], []*model.Hotel{&hotelData1})
	assert.Nil(t, repo.hotelIdKvStore[testHotelId2])
}

func TestInMemoryHotelRepository_InsertMultipleElementsWithSingleDestinationId(t *testing.T) {
	repo := NewInMemoryHotelRepository()
	repo.InsertHotel(&hotelData2)
	repo.InsertHotel(&hotelData3)
	assert.Equal(t, *repo.hotelIdKvStore[hotelData2.ID], hotelData2)
	assert.Equal(t, *repo.hotelIdKvStore[hotelData3.ID], hotelData3)
	assert.Equal(t, repo.destinationIdKvStore[testMultipleDestinationId], []*model.Hotel{&hotelData2, &hotelData3})
}

func TestInMemoryHotelRepository_GetHotelWithSingleHotelId(t *testing.T) {
	repo := prefilledTestingRepository()
	actual := repo.GetHotelsByHotelId([]string{testHotelId1})
	assert.Equal(t, len(actual), 1)
	assert.Equal(t, *actual[0], hotelData1)
}

func TestInMemoryHotelRepository_GetHotelWithMultipleHotelId(t *testing.T) {
	repo := prefilledTestingRepository()
	actual := repo.GetHotelsByHotelId([]string{testHotelId1, testHotelId2})
	assert.Equal(t, len(actual), 2)
	// the result needs to be ordered according to the keys supplied in the parameter
	assert.Equal(t, *actual[0], hotelData1)
	assert.Equal(t, *actual[1], hotelData2)
}

func TestInMemoryHotelRepository_GetHotelsWithNoExisitingHotelIdKey(t *testing.T) {
	repo := prefilledTestingRepository()
	actual := repo.GetHotelsByHotelId([]string{"0000"})
	assert.Equal(t, len(actual), 0)
	assert.Empty(t, actual)
}

func TestInMemoryHotelRepository_GetHotelsWithDestinationIdMapToOneHotel(t *testing.T) {
	repo := prefilledTestingRepository()
	actual := repo.GetHotelsByDestinationId(testSingleDestinationId)
	assert.Equal(t, len(actual), 1)
	assert.Equal(t, *actual[0], hotelData1)
}

func TestInMemoryHotelRepository_GetHotelsWithDestinationIdMapToMultipleHotel(t *testing.T) {
	repo := prefilledTestingRepository()
	actual := repo.GetHotelsByDestinationId(testMultipleDestinationId)
	assert.Equal(t, len(actual), 2)
	assert.Contains(t, actual, &hotelData2)
	assert.Contains(t, actual, &hotelData3)
}

func TestInMemoryHotelRepository_GetHotelsWithNoExisitingDestinationIdKey(t *testing.T) {
	repo := prefilledTestingRepository()
	actual := repo.GetHotelsByDestinationId(7000)
	assert.Equal(t, len(actual), 0)
	assert.Empty(t, actual)
}
