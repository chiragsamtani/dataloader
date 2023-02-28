package model

import "fmt"

type HotelLoaderData interface {
	GetId() string
	GetDestinationId() int
	GetName() string
	GetLocation() HotelLocation
	GetDescription() string
	GetAmenities() HotelAmenities
	GetImages() HotelImages
	GetBookingConditions() []string
	ConvertToHotelLoaderData(t interface{}) (HotelLoaderData, error)
}

type HotelLoaderDataFactory struct{}

func (sf HotelLoaderDataFactory) CreateSupplier(supplierType string) HotelLoaderData {
	switch supplierType {
	case "supplierA":
		return &HotelDataLoaderSupplierA{}
	case "supplierB":
		return &HotelDataLoaderSupplierB{}
	case "supplierC":
		return &HotelDataLoaderSupplierC{}
	default:
		panic(fmt.Sprintf("Invalid supplier type: %s", supplierType))
	}
}
