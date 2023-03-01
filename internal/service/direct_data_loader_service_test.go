package service

import (
	"datamerge/internal/model"
	"datamerge/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var logger = logrus.New()

const (
	ValidHotelId       = "iJhz"
	ValidDestinationId = 5432
)

var (
	supplierADataset = `[{
					"Id": "iJhz",
					"DestinationId": 5432,
					"Name": "Beach Villas Singapore",
					"Latitude": 1.264751,
					"Longitude": 103.824006,
					"Address": " 8 Sentosa Gateway, Beach Villas ",
					"City": "Singapore",
					"Country": "SG",
					"PostalCode": "098269",
					"Description": "  This 5 star hotel is located on the coastline of Singapore.",
					"Facilities": ["Pool", "BusinessCenter", "WiFi ", "DryCleaning", " Breakfast"]
				  }]`
	supplierBDataset = `[{
					"hotel_id": "iJhz",
					"destination_id": 5432,
					"hotel_name": "InterContinental",
					"location": {
					  "address": "1 Nanson Rd, Singapore 238909",
					  "country": "Singapore"
					},
					"details": "InterContinental Singapore Robertson Quay is luxury's preferred address offering stylishly cosmopolitan riverside living for discerning travelers to Singapore. Prominently situated along the Singapore River, the 225-room inspiring luxury hotel is easily accessible to the Marina Bay Financial District, Central Business District, Orchard Road and Singapore Changi International Airport, all located a short drive away. The hotel features the latest in Club InterContinental design and service experience, and five dining options including Publico, an Italian landmark dining and entertainment destination by the waterfront.",
					"amenities": {
					  "general": ["outdoor pool", "business center", "childcare", "parking", "bar", "dry cleaning", "wifi", "breakfast", "concierge"],
					  "room": ["aircon", "minibar", "tv", "bathtub", "hair dryer"]
					},
					"images": {
					  "rooms": [
						{ "link": "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i93_m.jpg", "caption": "Double room" },
						{ "link": "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i94_m.jpg", "caption": "Bathroom" }
					  ],
					  "site": [
						{ "link": "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i1_m.jpg", "caption": "Restaurant" },
						{ "link": "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i2_m.jpg", "caption": "Hotel Exterior" },
						{ "link": "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i5_m.jpg", "caption": "Entrance" },
						{ "link": "https://d2ey9sqrvkqdfs.cloudfront.net/Sjym/i24_m.jpg", "caption": "Bar" }
					  ]
					},
					"booking_conditions": []
	}]`
	supplierCDataset = `[{
					"id": "iJhz",
					"destination": 5432,
					"name": "Beach Villas Singapore",
					"lat": 1.264751,
					"lng": 103.824006,
					"address": "8 Sentosa Gateway, Beach Villas, 098269",
					"info": "Located at the western tip of Resorts World Sentosa, guests at the Beach Villas are guaranteed privacy while they enjoy spectacular views of glittering waters. Guests will find themselves in paradise with this series of exquisite tropical sanctuaries, making it the perfect setting for an idyllic retreat. Within each villa, guests will discover living areas and bedrooms that open out to mini gardens, private timber sundecks and verandahs elegantly framing either lush greenery or an expanse of sea. Guests are assured of a superior slumber with goose feather pillows and luxe mattresses paired with 400 thread count Egyptian cotton bed linen, tastefully paired with a full complement of luxurious in-room amenities and bathrooms boasting rain showers and free-standing tubs coupled with an exclusive array of ESPA amenities and toiletries. Guests also get to enjoy complimentary day access to the facilities at Asia’s flagship spa – the world-renowned ESPA.",
					"amenities": ["Aircon", "Tv", "Coffee machine", "Kettle", "Hair dryer", "Iron", "Tub"],
					"images": {
					  "rooms": [
						{ "url": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg", "description": "Double room" },
						{ "url": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/4.jpg", "description": "Bathroom" }
					  ],
					  "amenities": [
						{ "url": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/0.jpg", "description": "RWS" },
						{ "url": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/6.jpg", "description": "Sentosa Gateway" }
					  ]
					}
			}]`
)

func TestDirectDataLoaderService_LoadDataWithBadurl(t *testing.T) {
	repo := repository.NewInMemoryHotelRepository()
	loader := NewDirectDataLoaderService("supplierA:badurl", repo, logger)
	err := loader.LoadData()
	assert.Error(t, err)
	assert.IsType(t, err, &model.HttpError{})
}

func TestDirectDataLoaderService_UrlReturnNonJsonData(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html></htmL>"))
	}))
	defer mockHttpServer.Close()
	repo := repository.NewInMemoryHotelRepository()
	loader := NewDirectDataLoaderService("supplierA:"+mockHttpServer.URL, repo, logger)
	err := loader.LoadData()
	assert.Error(t, err)
	assert.IsType(t, err, &model.JsonError{})
}

func TestDirectDataLoaderService_WithValidSupplierADataset(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(supplierADataset))
	}))
	defer mockHttpServer.Close()
	repo := repository.NewInMemoryHotelRepository()
	loader := NewDirectDataLoaderService("supplierA:"+mockHttpServer.URL, repo, logger)
	err := loader.LoadData()
	assert.Nil(t, err)
	persistedData := repo.GetHotelsByHotelIds([]string{"iJhz"})
	assert.Equal(t, len(persistedData), 1)
	assert.Equal(t, persistedData[0].ID, "iJhz")
	assert.Equal(t, persistedData[0].DestinationID, ValidDestinationId)
}

func TestDirectDataLoaderService_WithValidSupplierBDataset(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(supplierBDataset))
	}))
	defer mockHttpServer.Close()
	repo := repository.NewInMemoryHotelRepository()
	loader := NewDirectDataLoaderService("supplierB:"+mockHttpServer.URL, repo, logger)
	err := loader.LoadData()
	assert.Nil(t, err)
	persistedData := repo.GetHotelsByHotelIds([]string{ValidHotelId})
	assert.Equal(t, len(persistedData), 1)
	assert.Equal(t, persistedData[0].ID, ValidHotelId)
	assert.Equal(t, persistedData[0].DestinationID, ValidDestinationId)
}

func TestDirectDataLoaderService_WithValidSupplierCDataset(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(supplierCDataset))
	}))
	defer mockHttpServer.Close()
	repo := repository.NewInMemoryHotelRepository()
	loader := NewDirectDataLoaderService("supplierC:"+mockHttpServer.URL, repo, logger)
	err := loader.LoadData()
	assert.Nil(t, err)
	persistedData := repo.GetHotelsByHotelIds([]string{ValidHotelId})
	assert.Equal(t, len(persistedData), 1)
	assert.Equal(t, persistedData[0].ID, ValidHotelId)
	assert.Equal(t, persistedData[0].DestinationID, ValidDestinationId)
}

func TestDirectDataLoaderService_WithInvalidSupplierADataset(t *testing.T) {
	// use supplierB dataset for supplierA, this data should be skipped
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(supplierBDataset))
	}))
	defer mockHttpServer.Close()
	repo := repository.NewInMemoryHotelRepository()
	loader := NewDirectDataLoaderService("supplierA:"+mockHttpServer.URL, repo, logger)
	err := loader.LoadData()
	assert.Nil(t, err)
	persistedData := repo.GetHotelsByHotelIds([]string{ValidHotelId})
	assert.Equal(t, len(persistedData), 0)
}

func TestDirectDataLoaderService_WithInvalidSupplierBDataset(t *testing.T) {
	// use supplierA dataset for supplierB, this data should be skipped
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(supplierADataset))
	}))
	defer mockHttpServer.Close()
	repo := repository.NewInMemoryHotelRepository()
	loader := NewDirectDataLoaderService("supplierB:"+mockHttpServer.URL, repo, logger)
	err := loader.LoadData()
	assert.Nil(t, err)
	persistedData := repo.GetHotelsByHotelIds([]string{ValidHotelId})
	assert.Equal(t, len(persistedData), 0)
}

func TestDirectDataLoaderService_WithInvalidSupplierCDataset(t *testing.T) {
	// use supplierB dataset for supplierC, this data should be skipped
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(supplierBDataset))
	}))
	defer mockHttpServer.Close()
	repo := repository.NewInMemoryHotelRepository()
	loader := NewDirectDataLoaderService("supplierC:"+mockHttpServer.URL, repo, logger)
	err := loader.LoadData()
	assert.Nil(t, err)
	persistedData := repo.GetHotelsByHotelIds([]string{ValidHotelId})
	assert.Equal(t, len(persistedData), 0)
}
