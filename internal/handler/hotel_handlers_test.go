package handler

import (
	"bytes"
	"datamerge/internal/model"
	"encoding/json"
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

// since we are stubbing this out for our positive case, we
// can state that it will never return an error
func (h *HotelServiceMock) SearchHotels(dto model.HotelRequestDTO) (interface{}, error) {
	// mock will record that the method was called and may
	// optionally take in some parameter it was called with
	args := h.Called(dto)
	return args.Get(0), nil
}

func generateMock() *HotelServiceMock {
	hotelSvcMock := new(HotelServiceMock)
	hotelSvcMock.On("SearchHotels").Return("test", nil)
	return hotelSvcMock
}

func TestHotelHandlerSearchHotels_withInvalidMethod(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/hotels", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewHotelHandler(generateMock())

	// function under test
	handler.SearchHotels(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestHotelHandlerSearchHotels_withInvalidContentType(t *testing.T) {
	req, err := http.NewRequest("POST", "/hotels", nil)
	req.Header.Set("Content-Type", "")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewHotelHandler(generateMock())

	// function under test
	handler.SearchHotels(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestHotelHandlerSearchHotels_withInvalidRequestBody(t *testing.T) {
	emptyJson := []byte(`{}`)
	req, err := http.NewRequest("POST", "/hotels", bytes.NewBuffer(emptyJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewHotelHandler(generateMock())

	// function under test
	handler.SearchHotels(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestHotelHandlerSearchHotels_withInvalidHotelIdsType(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]string{
		"hotel_ids": "2lx0",
	})
	req, err := http.NewRequest("POST", "/hotels", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewHotelHandler(generateMock())

	// function under test
	handler.SearchHotels(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestHotelHandlerSearchHotels_PositiveCase(t *testing.T) {
	reqBody, _ := json.Marshal(map[string][]string{
		"hotel_ids": {"2lx0"},
	})
	req, err := http.NewRequest("POST", "/hotels", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	mockSvc := generateMock()
	// use mock.Anything to signify that the argument under the function being tested
	// should not be taken into consideration
	mockSvc.On("SearchHotels", mock.Anything).Return("success")
	handler := NewHotelHandler(mockSvc)

	// function under test
	handler.SearchHotels(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
