package repository

import (
	"datamerge/internal/config"
	"datamerge/internal/model"
)

type HotelRepository interface {
	GetHotelsByHotelIds(hotelIds []string) []*model.Hotel
	GetHotelsByDestinationId(destinationId int) []*model.Hotel
	InsertHotel(hotel *model.Hotel)
}

func RepositoryFactory(cfg *config.RootConfig) (HotelRepository, error) {
	switch cfg.GetRepositoryType() {
	case "in-memory":
		return NewInMemoryHotelRepository(), nil
	case "mysql":
		db, err := config.OpenMysqlConn(cfg)
		if err != nil {
			return nil, err
		}
		return NewMysqlHotelRepository(db), nil
	default:
		return nil, &model.InvalidRepositoryType{}
	}
}
