package service

import (
	"datamerge/internal/model"
	"datamerge/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var StringCaseOption = cases.Title(language.English)

func MergeData(existingData model.Hotel, data model.HotelLoaderData) *model.Hotel {
	return &model.Hotel{
		ID:                data.GetId(),
		DestinationID:     data.GetDestinationId(),
		Name:              mergeName(existingData.Name, data.GetName()),
		Location:          mergeLocation(existingData.Location, data.GetLocation()),
		Description:       mergeDescription(existingData.Description, data.GetDescription()),
		Amenities:         mergeAmenities(existingData.Amenities, data.GetAmenities()),
		Images:            mergeImages(existingData.Images, data.GetImages()),
		BookingConditions: mergeBookingConditions(existingData.BookingConditions, data.GetBookingConditions()),
	}
}

func mergeName(exist, new string) string {
	return utils.MergeStringFieldByLength(exist, new, &StringCaseOption)
}

func mergeDescription(exist, new string) string {
	return utils.MergeStringFieldByLength(exist, new, &StringCaseOption)
}

func mergeBookingConditions(exist, new []string) []string {
	// no need to modify existing booking conditions string with
	// uppercase/lowercase or titles
	return utils.MergeStringArrayField(exist, new, nil)
}

func mergeLocation(exist, new model.HotelLocation) model.HotelLocation {
	return model.HotelLocation{
		Country: utils.MergeCountry(exist.Country, new.Country),
		City:    utils.MergeStringFieldByLength(exist.City, new.City, &StringCaseOption),
		Address: utils.MergeStringFieldByLength(exist.Address, new.Address, &StringCaseOption),
		Lat:     utils.MergingCoordinateFields(exist.Lat, new.Lat),
		Lng:     utils.MergingCoordinateFields(exist.Lng, new.Lng),
	}
}

func mergeAmenities(exist, new model.HotelAmenities) model.HotelAmenities {
	return model.HotelAmenities{
		Room:    utils.MergeStringArrayField(exist.Room, new.Room, &StringCaseOption),
		General: utils.MergeStringArrayField(exist.General, new.General, &StringCaseOption),
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
