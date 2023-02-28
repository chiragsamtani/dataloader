package model

import "encoding/json"

// HotelDataLoaderSupplierB represents an adapter of the HotelLoaderData that
// represents struct binding to one of the Supplier datasets.
// Fields such as Lat/Long whose types can be variable will be declared as an interface
// and if the types are not compatible with the internal format (float64) then those
// values will be ignored
// DestinationId is a required field and will NOT be ignored
type HotelDataLoaderSupplierB struct {
	HotelID           string             `json:"hotel_id"`
	DestinationID     int                `json:"destination_id"`
	HotelName         string             `json:"hotel_name"`
	Location          LocationSupplierB  `json:"location"`
	Details           string             `json:"details"`
	Amenities         AmenitiesSupplierB `json:"amenities"`
	Images            ImagesSupplierB    `json:"images"`
	BookingConditions []string           `json:"booking_conditions"`
}

type AmenitiesSupplierB struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}

type LocationSupplierB struct {
	Address string `json:"address"`
	Country string `json:"country"`
}

type ImagesSupplierB struct {
	Rooms []ImageSupplierB `json:"rooms"`
	Site  []ImageSupplierB `json:"site"`
}

type ImageSupplierB struct {
	Link    string `json:"link"`
	Caption string `json:"caption"`
}

func (h *HotelDataLoaderSupplierB) ConvertToHotelLoaderData(t interface{}) (HotelLoaderData, error) {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	var result HotelDataLoaderSupplierB
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (h *HotelDataLoaderSupplierB) GetId() string {
	return h.HotelID
}

func (h *HotelDataLoaderSupplierB) GetDestinationId() int {
	return h.DestinationID
}

func (h *HotelDataLoaderSupplierB) GetName() string {
	return h.HotelName
}

func (h *HotelDataLoaderSupplierB) GetLocation() HotelLocation {
	return HotelLocation{
		Address: h.Location.Address,
		Country: h.Location.Country,
	}
}

func (h *HotelDataLoaderSupplierB) GetDescription() string {
	return h.Details
}

func (h *HotelDataLoaderSupplierB) GetImages() HotelImages {
	roomImages := make([]Image, 0)
	for _, images := range h.Images.Rooms {
		roomImages = append(roomImages, Image{
			Link:        images.Link,
			Description: images.Caption,
		})
	}
	siteImages := make([]Image, 0)
	for _, images := range h.Images.Site {
		siteImages = append(siteImages, Image{
			Link:        images.Link,
			Description: images.Caption,
		})
	}
	return HotelImages{
		Rooms: roomImages,
		Site:  siteImages,
	}
}

func (h *HotelDataLoaderSupplierB) GetAmenities() HotelAmenities {
	return HotelAmenities{
		Room:    h.Amenities.Room,
		General: h.Amenities.General,
	}
}

func (h *HotelDataLoaderSupplierB) GetBookingConditions() []string {
	return h.BookingConditions
}
