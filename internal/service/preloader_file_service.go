package service

import (
	"datamerge/internal/repository"
	"encoding/json"
	"net/http"
	"strings"
)

type PreloadFileService struct {
	configUrls string
	repo       repository.HotelRepository
}

func NewPreloadFileService(configs string, repo repository.HotelRepository) *PreloadFileService {
	return &PreloadFileService{repo: repo}
}

func readJsonFileFromUrl(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var jsonData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (f *PreloadFileService) PreloadData() {
	urls := strings.Split(f.configUrls, ",")
	for _, url := range urls {
		readJsonFileFromUrl(url)
	}
}
