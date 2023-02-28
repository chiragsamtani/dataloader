package model

import "encoding/json"

type HotelDataLoaderSupplierC struct {
	ID          string          `json:"id"`
	Destination int             `json:"destination"`
	Name        string          `json:"name"`
	Lat         interface{}     `json:"lat"`
	Lng         interface{}     `json:"lng"`
	Address     string          `json:"address"`
	Info        string          `json:"info"`
	Amenities   []string        `json:"amenities"`
	Images      ImagesSupplierC `json:"images"`
}

type ImagesSupplierC struct {
	Rooms     []ImageSupplierC `json:"rooms"`
	Amenities []ImageSupplierC `json:"amenities"`
}

type ImageSupplierC struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

func (h *HotelDataLoaderSupplierC) ConvertToHotelLoaderData(t interface{}) (HotelLoaderData, error) {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	var result HotelDataLoaderSupplierC
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (h *HotelDataLoaderSupplierC) GetId() string {
	return h.ID
}

func (h *HotelDataLoaderSupplierC) GetDestinationId() int {
	return h.Destination
}

func (h *HotelDataLoaderSupplierC) GetName() string {
	return h.Name
}

func (h *HotelDataLoaderSupplierC) GetLocation() HotelLocation {
	hotelLocation := HotelLocation{
		Address: h.Address,
	}
	lat, ok := h.Lat.(float64)
	if !ok {
		return hotelLocation
	}
	long, ok := h.Lng.(float64)
	if !ok {
		return hotelLocation
	}
	hotelLocation.Lat = lat
	hotelLocation.Lng = long
	return hotelLocation
}

func (h *HotelDataLoaderSupplierC) GetDescription() string {
	return h.Info
}

func (h *HotelDataLoaderSupplierC) GetImages() HotelImages {
	roomImages := make([]Image, 0)
	for _, images := range h.Images.Rooms {
		roomImages = append(roomImages, Image{
			Link:        images.URL,
			Description: images.Description,
		})
	}
	amenitiesImage := make([]Image, 0)
	for _, images := range h.Images.Amenities {
		amenitiesImage = append(amenitiesImage, Image{
			Link:        images.URL,
			Description: images.Description,
		})
	}
	return HotelImages{
		Rooms:     roomImages,
		Amenities: amenitiesImage,
	}
}

func (h *HotelDataLoaderSupplierC) GetAmenities() HotelAmenities {
	return HotelAmenities{
		Room: h.Amenities,
	}
}

func (h *HotelDataLoaderSupplierC) GetBookingConditions() []string {
	return []string{}
}
