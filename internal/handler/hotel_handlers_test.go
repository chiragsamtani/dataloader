package handler

import (
	"bytes"
	"encoding/json"
	"errors"
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
func (h *HotelServiceMock) SearchHotelsByHotelId(hotelId []string) (interface{}, error) {
	// mock will record that the method was called and may
	// optionally take in some parameter it was called with
	args := h.Called(hotelId)
	return args.Get(0), args.Error(1)
}

func (h *HotelServiceMock) SearchHotelsByDestinationId(destinationId int) (interface{}, error) {
	// mock will record that the method was called and may
	// optionally take in some parameter it was called with
	args := h.Called(destinationId)
	return args.Get(0), args.Error(1)
}

func generateMock() *HotelServiceMock {
	hotelSvcMock := new(HotelServiceMock)
	hotelSvcMock.On("SearchHotelsByHotelId").Return("test", nil)
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

func TestHotelHandlerSearchHotels_PositiveCaseWithValidHotelId(t *testing.T) {
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
	mockSvc.On("SearchHotelsByHotelId", mock.Anything).Return("success", nil)
	handler := NewHotelHandler(mockSvc)

	// function under test
	handler.SearchHotels(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHotelHandlerSearchHotels_PositiveCaseWithValidDestinationId(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]int{
		"destination_id": 5432,
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
	mockSvc.On("SearchHotelsByDestinationId", mock.Anything).Return("success", nil)
	handler := NewHotelHandler(mockSvc)

	// function under test
	handler.SearchHotels(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHotelHandlerSearchHotels_WithValidHotelIdAndServiceError(t *testing.T) {
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
	mockSvc.On("SearchHotelsByHotelId", mock.Anything).Return(nil, errors.New("mock server error"))
	handler := NewHotelHandler(mockSvc)

	// function under test
	handler.SearchHotels(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestHotelHandlerSearchHotels_WithValidDestinationIdAndServiceError(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]int{
		"destination_id": 5432,
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
	mockSvc.On("SearchHotelsByDestinationId", mock.Anything).Return(nil, errors.New("mock server error"))
	handler := NewHotelHandler(mockSvc)

	// function under test
	handler.SearchHotels(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
