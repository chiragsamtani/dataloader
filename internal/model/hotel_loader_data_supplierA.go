package model

import (
	"encoding/json"
)

// HotelDataLoaderSupplierA represents an adapter of the HotelLoaderData that
// represents struct binding to one of the Supplier datasets.
// Fields such as Lat/Long whose types can be variable will be declared as an interface
// and if the types are not compatible with the internal format (float64) then those
// values will be ignored
// DestinationId is a required field and will NOT be ignored
type HotelDataLoaderSupplierA struct {
	ID            string      `json:"Id"`
	DestinationID int         `json:"DestinationId"`
	Name          string      `json:"Name"`
	Latitude      interface{} `json:"Latitude"`
	Longitude     interface{} `json:"Longitude"`
	Address       string      `json:"Address"`
	City          string      `json:"City"`
	Country       string      `json:"Country"`
	PostalCode    string      `json:"PostalCode"`
	Description   string      `json:"Description"`
	Facilities    []string    `json:"Facilities"`
}

func (h *HotelDataLoaderSupplierA) ConvertToHotelLoaderData(t interface{}) (HotelLoaderData, error) {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	var result HotelDataLoaderSupplierA
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (h *HotelDataLoaderSupplierA) GetId() string {
	return h.ID
}

func (h *HotelDataLoaderSupplierA) GetDestinationId() int {
	return h.DestinationID
}

func (h *HotelDataLoaderSupplierA) GetName() string {
	return h.Name
}

func (h *HotelDataLoaderSupplierA) GetLocation() HotelLocation {
	lat, ok := h.Latitude.(float64)
	if !ok {
		return HotelLocation{}
	}
	long, ok := h.Latitude.(float64)
	if !ok {
		return HotelLocation{}
	}
	return HotelLocation{
		Lat:     lat,
		Lng:     long,
		Address: h.Address,
		City:    h.City,
		Country: h.Country,
	}
}

func (h *HotelDataLoaderSupplierA) GetDescription() string {
	return h.Description
}

func (h *HotelDataLoaderSupplierA) GetImages() HotelImages {
	return HotelImages{}
}

func (h *HotelDataLoaderSupplierA) GetAmenities() HotelAmenities {
	return HotelAmenities{General: h.Facilities}
}

func (h *HotelDataLoaderSupplierA) GetBookingConditions() []string {
	return []string{}
}
