package model

type Hotel struct {
	ID                string         `json:"id"`
	DestinationID     int            `json:"destination_id"`
	Name              string         `json:"name"`
	Location          HotelLocation  `json:"location"`
	Description       string         `json:"description"`
	Amenities         HotelAmenities `json:"amenities"`
	Images            HotelImages    `json:"images"`
	BookingConditions []string       `json:"booking_conditions"`
}

type HotelLocation struct {
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Address string  `json:"address"`
	City    string  `json:"city"`
	Country string  `json:"country"`
}

type HotelAmenities struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}

type Image struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}

type HotelImages struct {
	Rooms     []Image `json:"rooms"`
	Site      []Image `json:"site"`
	Amenities []Image `json:"amenities"`
}
