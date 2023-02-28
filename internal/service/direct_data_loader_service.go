package service

import (
	"datamerge/internal/model"
	"datamerge/internal/repository"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	ConfigSubstringLimitSeparator = 2
)

type DirectDataLoaderService struct {
	configs                string
	repo                   repository.HotelRepository
	hotelLoaderDataFactory model.HotelLoaderDataFactory
	logger                 *logrus.Logger
}

func NewDirectDataLoaderService(configs string, repo repository.HotelRepository, logger *logrus.Logger) *DirectDataLoaderService {
	return &DirectDataLoaderService{configs: configs, repo: repo, logger: logger}
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
		configSplit := strings.SplitN(config, ":", ConfigSubstringLimitSeparator)
		if len(configSplit) != 2 {
			panic("supplier config is broken, please check env variable SUPPLIER_CONFIG")
		}
		supplierIdentifier := configSplit[0]
		url := configSplit[1]
		results, err := readJsonFileFromUrl(url)
		if err != nil {
			d.logger.WithFields(logrus.Fields{
				"supplier": supplierIdentifier,
				"url":      url,
			}).Warn(err)
			continue
		}
		supplierModel := d.hotelLoaderDataFactory.CreateSupplier(supplierIdentifier)
		var hotels []model.HotelLoaderData
		for _, result := range results {
			hotel, err := supplierModel.ConvertToHotelLoaderData(result)
			if err != nil {
				d.logger.WithFields(logrus.Fields{
					"supplier": supplierIdentifier,
					"url":      url,
				}).Warn(err)
				continue
			}
			hotels = append(hotels, hotel)
		}

		for _, hotel := range hotels {
			var existingData model.Hotel
			hotelData := d.repo.GetHotelsByHotelId([]string{hotel.GetId()})
			if hotelData != nil {
				existingData = *hotelData[0]
			}
			d.repo.InsertHotel(MergeData(existingData, hotel))
		}
	}
}
