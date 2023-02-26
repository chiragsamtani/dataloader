package handler

import (
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mock our HotelService dependency to the handler
type HotelServiceMock struct {
	mock.Mock
}

// we sadly have to stub all the methods - if we have multiple
// methods, that needs to be stubbed out as well
func (h *HotelServiceMock) GetHotels() (interface{}, error) {
	// mock will record that the method was called and may
	// optionally take in some parameter it was called with
	args := h.Called()
	return args.Get(0), args.Error(1)
}

func generateMock() *HotelServiceMock {
	hotelSvcMock := new(HotelServiceMock)
	hotelSvcMock.On("GetHotels").Return("test", nil)
	return hotelSvcMock
}

func TestHotelHandlerGetAllHotels_withValidMethod(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/hotels", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewHotelHandler(generateMock())

	// function under test
	handler.GetAllHotels(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHotelHandlerGetAllHotels_withInvalidMethod(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("POST", "/hotels", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewHotelHandler(generateMock())

	// function under test
	handler.GetAllHotels(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}
