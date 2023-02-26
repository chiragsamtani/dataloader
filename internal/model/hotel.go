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

type HotelImages struct {
	Rooms []struct {
		Link        string `json:"link"`
		Description string `json:"description"`
	} `json:"rooms"`
	Site []struct {
		Link        string `json:"link"`
		Description string `json:"description"`
	} `json:"site"`
	Amenities []struct {
		Link        string `json:"link"`
		Description string `json:"description"`
	} `json:"amenities"`
}
