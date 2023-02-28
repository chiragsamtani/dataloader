package service

import (
	"datamerge/internal/repository"
	"fmt"
	"testing"
)

const (
	TEST_URL_SUPPLIER_A = "http://www.mocky.io/v2/5ebbea002e000054009f3ffc"
	TEST_URL_SUPPLIER_B = "http://www.mocky.io/v2/5ebbea102e000029009f3fff"
)

func TestInMemoryHotelRepository_GetHotelsWithModification(t *testing.T) {
	repo := repository.NewInMemoryHotelRepository()
	s := NewDirectDataLoaderService("supplierA:"+TEST_URL_SUPPLIER_A, repo)
	s.LoadData()
	s = NewDirectDataLoaderService("supplierB:"+TEST_URL_SUPPLIER_B, repo)
	s.LoadData()
	fmt.Println(repo.GetHotelsByHotelId([]string{"iJhz"})[0].Amenities.General)
}
