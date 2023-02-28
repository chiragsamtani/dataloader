package service

import (
	"datamerge/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	supplierA = model.HotelDataLoaderSupplierA{
		ID:            "ibx8",
		DestinationID: 5432,
		Name:          "Hotel Singapura",
		Description:   "Beautiful hotel with luxurious rooms",
		Latitude:      1.45090,
		Longitude:     -12.4490,
		City:          "Singapore",
		Country:       "SG",
		Address:       "1 Singapore Road",
		Facilities:    []string{"Pool"},
	}

	supplierB = model.HotelDataLoaderSupplierB{
		HotelID:       "ibx8",
		DestinationID: 5432,
		HotelName:     "Hotel Singapura",
		Details:       "Beautiful hotel with luxurious rooms",
		Location: model.LocationSupplierB{
			Address: "1 Singapore Road",
			Country: "SG",
		},
		Amenities: model.AmenitiesSupplierB{
			Room:    []string{"Microwave"},
			General: []string{"Fitness Center"},
		},
		Images: model.ImagesSupplierB{
			Rooms: []model.ImageSupplierB{{Link: "link1", Caption: "caption1"}},
		},
	}

	supplierC = model.HotelDataLoaderSupplierC{
		ID:          "ibx8",
		Destination: 5432,
		Name:        "Hotel Singapura",
		Info:        "Beautiful hotel with luxurious rooms",
		Lat:         1.45090,
		Lng:         -12.4490,
		Address:     "1 Singapore Road",
		Amenities:   []string{"Jacuzzi"},
		Images:      model.ImagesSupplierC{Rooms: []model.ImageSupplierC{{URL: "url1", Description: "desc1"}}},
	}

	supplierDataSets = []model.HotelLoaderData{&supplierA, &supplierB, &supplierC}
)

func TestMergeData_WithEmptyExistingValues(t *testing.T) {
	existing := model.Hotel{}
	for _, supplier := range supplierDataSets {
		actual := MergeData(existing, supplier)
		assert.Equal(t, actual.ID, supplier.GetId())
		assert.Equal(t, actual.DestinationID, supplier.GetDestinationId())
	}
}

func TestMergeData_WithLongerNameOnNewData(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Name:          "Hotel SG",
	}
	for _, supplier := range supplierDataSets {
		actual := MergeData(existing, supplier)
		assert.Equal(t, actual.ID, supplier.GetId())
		assert.Equal(t, actual.DestinationID, supplier.GetDestinationId())
		assert.Equal(t, actual.Name, supplier.GetName())
	}
}

func TestMergeData_WithLongerDescriptionOnNewData(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Description:   "Beautiful hotel",
	}
	sanitizedDescription := "Beautiful hotel with luxurious rooms"
	for _, supplier := range supplierDataSets {
		actual := MergeData(existing, supplier)
		assert.Equal(t, actual.ID, supplier.GetId())
		assert.Equal(t, actual.DestinationID, supplier.GetDestinationId())
		assert.Equal(t, sanitizedDescription, supplier.GetDescription())
	}
}

func TestMergeData_WithZeroedLatLong(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Location: model.HotelLocation{
			Lat: 0.0,
			Lng: 0.0,
		},
	}
	expectedResult := model.HotelLocation{
		Lat: 1.45090,
		Lng: -12.4490,
	}
	for _, supplier := range []model.HotelLoaderData{&supplierA, &supplierC} {
		actual := MergeData(existing, supplier)
		assert.Equal(t, actual.ID, supplier.GetId())
		assert.Equal(t, actual.DestinationID, supplier.GetDestinationId())
		assert.Equal(t, actual.Location.Lat, expectedResult.Lat)
		assert.Equal(t, actual.Location.Lng, expectedResult.Lng)
	}
}

func TestMergeData_WithExistingLatLong(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Location: model.HotelLocation{
			Lat: 1.45090001,
			Lng: -12.009401,
		},
	}
	expectedResult := model.HotelLocation{
		Lat: 1.45090001,
		Lng: -12.009401,
	}
	for _, supplier := range supplierDataSets {
		actual := MergeData(existing, supplier)
		assert.Equal(t, actual.ID, supplier.GetId())
		assert.Equal(t, actual.DestinationID, supplier.GetDestinationId())
		assert.Equal(t, actual.Location.Lat, expectedResult.Lat)
		assert.Equal(t, actual.Location.Lng, expectedResult.Lng)
	}
}

func TestMergeData_WithEmptyExistingCountryName(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
	}
	expectedResult := model.HotelLocation{
		Country: "SG",
	}
	for _, supplier := range []model.HotelLoaderData{&supplierA, &supplierB} {
		actual := MergeData(existing, supplier)
		assert.Equal(t, actual.ID, supplier.GetId())
		assert.Equal(t, actual.DestinationID, supplier.GetDestinationId())
		assert.Equal(t, actual.Location.Country, expectedResult.Country)
	}
}

func TestMergeData_WithExistingCountryName(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Location: model.HotelLocation{
			Country: "Singapore",
		},
	}
	expectedResult := model.HotelLocation{
		Country: "SG",
	}
	for _, supplier := range []model.HotelLoaderData{&supplierA, &supplierB} {
		actual := MergeData(existing, supplier)
		assert.Equal(t, actual.ID, supplier.GetId())
		assert.Equal(t, actual.DestinationID, supplier.GetDestinationId())
		assert.Equal(t, actual.Location.Country, expectedResult.Country)
	}
}

func TestMergeData_WithExistingCityName(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Location: model.HotelLocation{
			City: "Singapore City",
		},
	}
	expectedResult := model.HotelLocation{
		City: "Singapore City",
	}

	actual := MergeData(existing, &supplierA)
	assert.Equal(t, actual.ID, supplierA.GetId())
	assert.Equal(t, actual.DestinationID, supplierA.GetDestinationId())
	assert.Equal(t, actual.Location.City, expectedResult.City)

}

func TestMergeData_WithEmptyExistingCityName(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
	}
	actual := MergeData(existing, &supplierA)
	assert.Equal(t, actual.ID, supplierA.GetId())
	assert.Equal(t, actual.DestinationID, supplierA.GetDestinationId())
	assert.Equal(t, actual.Location.City, supplierA.City)
}

func TestMergeData_WithExistingAddress(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Location: model.HotelLocation{
			Address: "Jln 1",
		},
	}
	expectedResult := model.HotelLocation{
		Address: "1 Singapore Road",
	}
	for _, supplier := range []model.HotelLoaderData{&supplierA, &supplierB} {
		actual := MergeData(existing, supplier)
		assert.Equal(t, actual.ID, supplier.GetId())
		assert.Equal(t, actual.DestinationID, supplier.GetDestinationId())
		assert.Equal(t, actual.Location.Address, expectedResult.Address)
	}
}

func TestMergeData_WithExistingLongerDescription(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Description:   "Description one",
	}
	newData := model.HotelDataLoaderSupplierA{
		ID:            "ibx8",
		DestinationID: 5432,
		Description:   "Desc one",
	}
	expectedResult := "Description One"
	actual := MergeData(existing, &newData)
	assert.Equal(t, actual.ID, newData.ID)
	assert.Equal(t, actual.DestinationID, newData.DestinationID)
	assert.Equal(t, actual.Description, expectedResult)
}

func TestMergeData_WithExistingAmenities(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Amenities: model.HotelAmenities{
			Room:    []string{"Tv"},
			General: []string{"Pool"},
		},
	}
	newData := model.HotelDataLoaderSupplierA{
		ID:            "ibx8",
		DestinationID: 5432,
		Facilities:    []string{"Fitness Center"},
	}
	expectedResult := model.HotelAmenities{
		Room:    []string{"Tv"},
		General: []string{"Pool", "Fitness Center"},
	}
	actual := MergeData(existing, &newData)
	assert.Equal(t, actual.ID, newData.ID)
	assert.Equal(t, actual.DestinationID, newData.DestinationID)
	assert.Equal(t, actual.Amenities, expectedResult)
}

func TestMergeData_WithNoExistingAmenitiesForGeneral(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Amenities: model.HotelAmenities{
			Room: []string{"Tv"},
		},
	}
	expectedResult := model.HotelAmenities{
		Room:    []string{"Tv"},
		General: []string{"Pool"},
	}
	actual := MergeData(existing, &supplierA)
	assert.Equal(t, actual.ID, supplierA.ID)
	assert.Equal(t, actual.DestinationID, supplierA.DestinationID)
	assert.Equal(t, actual.Amenities, expectedResult)
}

func TestMergeData_WithNoExistingAmenitiesForRooms(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Amenities: model.HotelAmenities{
			General: []string{"Pool"},
		},
	}
	expectedResult := model.HotelAmenities{
		Room:    []string{"Jacuzzi"},
		General: []string{"Pool"},
	}
	actual := MergeData(existing, &supplierC)
	assert.Equal(t, actual.ID, supplierC.ID)
	assert.Equal(t, actual.DestinationID, supplierC.Destination)
	assert.Equal(t, actual.Amenities, expectedResult)
}

func TestMergeData_WithExistingAmenitiesForRoomsAndGeneral(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Amenities: model.HotelAmenities{
			General: []string{"Pool"},
			Room:    []string{"Tv"},
		},
	}
	expectedResult := model.HotelAmenities{
		Room:    []string{"Tv", "Microwave"},
		General: []string{"Pool", "Fitness Center"},
	}
	actual := MergeData(existing, &supplierB)
	assert.Equal(t, actual.ID, supplierB.HotelID)
	assert.Equal(t, actual.DestinationID, supplierB.DestinationID)
	assert.Equal(t, actual.Amenities, expectedResult)
}

func TestMergeData_WithNoImages(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
	}
	expectedResult := model.HotelImages{
		Rooms: []model.Image{{Link: "link1", Description: "caption1"}},
	}
	actual := MergeData(existing, &supplierB)
	assert.Equal(t, actual.ID, supplierB.HotelID)
	assert.Equal(t, actual.DestinationID, supplierB.DestinationID)
	assert.Equal(t, actual.Images, expectedResult)
}

func TestMergeData_WithExistingImages(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
		Images:        model.HotelImages{Rooms: []model.Image{{"imageurl1", "imagedesc1"}}},
	}
	expectedResult := model.HotelImages{
		Rooms: []model.Image{{Link: "imageurl1", Description: "imagedesc1"},
			{Link: "link1", Description: "caption1"}},
	}
	actual := MergeData(existing, &supplierB)
	assert.Equal(t, actual.ID, supplierB.HotelID)
	assert.Equal(t, actual.DestinationID, supplierB.DestinationID)
	assert.Equal(t, actual.Images, expectedResult)
}

func TestMergeData_NoImagesWithMultipleMerges(t *testing.T) {
	existing := model.Hotel{
		ID:            "ibx8",
		DestinationID: 5432,
	}
	expectedResult := model.HotelImages{
		Rooms: []model.Image{{Link: "link1", Description: "caption1"},
			{Link: "url1", Description: "desc1"}},
	}
	actual := MergeData(existing, &supplierB)
	actual = MergeData(*actual, &supplierC)
	assert.Equal(t, actual.ID, supplierB.HotelID)
	assert.Equal(t, actual.DestinationID, supplierB.DestinationID)
	assert.Equal(t, actual.Images, expectedResult)
}
