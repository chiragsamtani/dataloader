package model

import "encoding/json"

type HotelDataLoaderSupplierC struct {
	ID          string   `json:"id"`
	Destination int      `json:"destination"`
	Name        string   `json:"name"`
	Lat         float64  `json:"lat"`
	Lng         float64  `json:"lng"`
	Address     string   `json:"address"`
	Info        string   `json:"info"`
	Amenities   []string `json:"amenities"`
	Images      struct {
		Rooms []struct {
			URL         string `json:"url"`
			Description string `json:"description"`
		} `json:"rooms"`
		Amenities []struct {
			URL         string `json:"url"`
			Description string `json:"description"`
		} `json:"amenities"`
	} `json:"images"`
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
	return HotelLocation{
		Lat:     h.Lat,
		Lng:     h.Lng,
		Address: h.Address,
	}
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
