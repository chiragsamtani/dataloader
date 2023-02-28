package service

import (
	"datamerge/internal/model"
	"datamerge/internal/repository"
	"datamerge/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type DirectDataLoaderService struct {
	configs                string
	repo                   repository.HotelRepository
	hotelLoaderDataFactory model.HotelLoaderDataFactory
}

func NewDirectDataLoaderService(configs string, repo repository.HotelRepository) *DirectDataLoaderService {
	return &DirectDataLoaderService{configs: configs, repo: repo}
}

func readJsonFileFromUrl(url string) ([]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (d *DirectDataLoaderService) LoadData() {
	configs := strings.Split(d.configs, ",")
	for _, config := range configs {
		configSplit := strings.SplitN(config, ":", 2)
		supplierIdentifier := configSplit[0]
		url := configSplit[1]
		results, err := readJsonFileFromUrl(url)
		if err != nil {
			return
		}
		supplierModel := d.hotelLoaderDataFactory.CreateSupplier(supplierIdentifier)
		var hotels []model.HotelLoaderData
		for _, result := range results {
			hotel, err := supplierModel.ConvertToHotelLoaderData(result)
			if err != nil {
				fmt.Println(err)
				continue
			}
			hotels = append(hotels, hotel)
		}

		for _, hotel := range hotels {
			d.MergeData(hotel)
		}
	}
}

func (d *DirectDataLoaderService) MergeData(data model.HotelLoaderData) {
	var existingData model.Hotel
	hotelData := d.repo.GetHotelsByHotelId([]string{data.GetId()})
	if hotelData != nil {
		existingData = *hotelData[0]
	}
	hotelDTO := model.Hotel{
		ID:                data.GetId(),
		DestinationID:     data.GetDestinationId(),
		Name:              utils.MergeStringFieldByLength(existingData.Name, data.GetName()),
		Location:          mergeLocation(existingData.Location, data.GetLocation()),
		Description:       utils.MergeStringFieldByLength(existingData.Description, data.GetDescription()),
		Amenities:         mergeAmenities(existingData.Amenities, data.GetAmenities()),
		Images:            mergeImages(existingData.Images, data.GetImages()),
		BookingConditions: utils.MergeStringArrayField(existingData.BookingConditions, data.GetBookingConditions()),
	}

	d.repo.InsertHotel(&hotelDTO)
	return
}

func mergeLocation(exist, new model.HotelLocation) model.HotelLocation {
	return model.HotelLocation{
		Country: utils.MergeCountry(exist.Country, new.Country),
		City:    utils.MergeStringFieldByLength(exist.City, new.City),
		Address: utils.MergeStringFieldByLength(exist.Address, new.Address),
		Lat:     utils.MergingCoordinateFields(exist.Lat, new.Lat),
		Lng:     utils.MergingCoordinateFields(exist.Lng, new.Lng),
	}
}

func mergeAmenities(exist, new model.HotelAmenities) model.HotelAmenities {
	return model.HotelAmenities{
		Room:    utils.MergeStringArrayField(exist.Room, new.Room),
		General: utils.MergeStringArrayField(exist.General, new.General),
	}
}

func mergeImages(exist, new model.HotelImages) model.HotelImages {
	existingRoomImages := exist.Rooms
	existingAmenitiesImages := exist.Amenities
	existingSiteImages := exist.Site
	for _, roomImages := range new.Rooms {
		existingRoomImages = append(existingRoomImages, roomImages)
	}
	for _, amenitiesImages := range new.Amenities {
		existingAmenitiesImages = append(existingAmenitiesImages, amenitiesImages)
	}
	for _, siteImages := range new.Site {
		existingSiteImages = append(existingSiteImages, siteImages)
	}
	return model.HotelImages{
		Rooms:     existingRoomImages,
		Amenities: existingAmenitiesImages,
		Site:      existingSiteImages,
	}
}
