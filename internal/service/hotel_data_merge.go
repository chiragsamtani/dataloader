package service

import (
	"datamerge/internal/model"
	"datamerge/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	TitleFirstLetter = cases.Title(language.English)
	LowerCaseString  = cases.Lower(language.English)
)

// MergeData will create a new Hotel object after merging existingData with new data
// if exisitingData does not exist, then the new data will be returned as the new Hotel
// object after each Hotel fields are sanitized
func MergeData(existingData model.Hotel, data model.HotelLoaderData) *model.Hotel {
	s := &model.Hotel{
		ID:                data.GetId(),
		DestinationID:     data.GetDestinationId(),
		Name:              mergeName(existingData.Name, data.GetName()),
		Location:          mergeLocation(existingData.Location, data.GetLocation()),
		Description:       mergeDescription(existingData.Description, data.GetDescription()),
		Amenities:         mergeAmenities(existingData.Amenities, data.GetAmenities()),
		Images:            mergeImages(existingData.Images, data.GetImages()),
		BookingConditions: mergeBookingConditions(existingData.BookingConditions, data.GetBookingConditions()),
	}
	return s
}

// mergeName will choose the hotelName between the existing data and new data
// the decision of which hotel name to choose is made by the bigger hotel name
// length
func mergeName(exist, new string) string {
	return utils.MergeStringFieldByLength(exist, new, &TitleFirstLetter)
}

// mergeDescription will choose the hotel description between the existing data and new data
// the decision of which hotel description to choose is made by the bigger hotel description
// length
func mergeDescription(exist, new string) string {
	return utils.MergeStringFieldByLength(exist, new, &TitleFirstLetter)
}

// mergeBookingConditions will merge booking conditions between existing data and new data
// the merge will be a union between the existing and new data with no filter applied
func mergeBookingConditions(exist, new []string) []string {
	// no need to modify existing booking conditions string with
	// uppercase/lowercase or titles
	return utils.MergeStringArrayField(exist, new, nil)
}

// mergeLocation will merge locations between existing data and new data
// Country: will choose ISO-3601 compliant country codes and choose non-empty strings in that order
// City: will choose longer length city between existing and new data
// Address: will choose longer address between existing and new data
// Lat: will choose first non-zero data
// Lng: will choose first non-zero data
func mergeLocation(exist, new model.HotelLocation) model.HotelLocation {
	return model.HotelLocation{
		Country: utils.MergeCountry(exist.Country, new.Country),
		City:    utils.MergeStringFieldByLength(exist.City, new.City, &TitleFirstLetter),
		Address: utils.MergeStringFieldByLength(exist.Address, new.Address, &TitleFirstLetter),
		Lat:     utils.MergingCoordinateFields(exist.Lat, new.Lat),
		Lng:     utils.MergingCoordinateFields(exist.Lng, new.Lng),
	}
}

// mergeAmenities will merge amenities data between existing data and new data
// the data will be a union of existing and new data, filtering out any duplicate
// or similar data
func mergeAmenities(exist, new model.HotelAmenities) model.HotelAmenities {
	return model.HotelAmenities{
		Room:    utils.MergeStringArrayWithNoDuplicates(exist.Room, new.Room, &LowerCaseString),
		General: utils.MergeStringArrayWithNoDuplicates(exist.General, new.General, &LowerCaseString),
	}
}

// mergeImages will merge image data between existing data and new data
// the data will be a union of existing and new data images, the URLs are
// first checked to make sure we don't have any duplicate data
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
