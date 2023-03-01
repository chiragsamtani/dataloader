package repository

import (
	"datamerge/internal/model"
	"gorm.io/gorm"
)

type MysqlHotelRepository struct {
	db *gorm.DB
}

func NewMysqlHotelRepository(db *gorm.DB) *MysqlHotelRepository {
	mySqlRepo := &MysqlHotelRepository{
		db: db,
	}
	mySqlRepo.db.AutoMigrate(model.Hotel{})
	return mySqlRepo
}

func (m *MysqlHotelRepository) GetHotelsByHotelIds(hotelIds []string) []*model.Hotel {
	var result []*model.Hotel
	m.db.Find(result, hotelIds)
	return result
}

func (m *MysqlHotelRepository) GetHotelsByDestinationId(destinationId int) []*model.Hotel {
	return []*model.Hotel{}
}

func (m *MysqlHotelRepository) InsertHotel(hotel *model.Hotel) {

}
