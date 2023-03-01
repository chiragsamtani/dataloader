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

// DirectDataLoaderService will load json data from the configUrls directly
// the object type assigned from each url are given through the configs and
// returned from the CreateSupplier factory method
// Raw JSON returned from the url is converted to the supplier object type
// using the ConvertToHotelLoaderData defined by each supplier class
// newHotelData will then be merged with existing data by querying them
// from the repository using their hotelId (acts as the PK in this case)
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
		return nil, &model.HttpError{}
	}
	var result []interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, &model.JsonError{}
	}
	return result, nil
}

func (d *DirectDataLoaderService) LoadData() error {
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
			return err
		}
		supplierModel := d.hotelLoaderDataFactory.CreateSupplier(supplierIdentifier)
		var newHotelData []model.HotelLoaderData
		for _, result := range results {
			explicitSupplierTypeHotel, err := supplierModel.ConvertToHotelLoaderData(result)
			if err != nil {
				d.logger.WithFields(logrus.Fields{
					"supplier": supplierIdentifier,
					"url":      url,
				}).Warn(err)
				continue
			}
			newHotelData = append(newHotelData, explicitSupplierTypeHotel)
		}

		for _, hotel := range newHotelData {
			var existingData model.Hotel
			hotelData := d.repo.GetHotelsByHotelIds([]string{hotel.GetId()})
			if hotelData != nil {
				existingData = *hotelData[0]
			}
			d.repo.InsertHotel(MergeData(existingData, hotel))
		}
	}
	return nil
}
