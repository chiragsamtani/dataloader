package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenMysqlConn(d *RootConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		d.GetDatabaseUsername(),
		d.GetDatabasePassword(),
		d.GetDatabaseHost(),
		d.GetDatabaseName())

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		AllowGlobalUpdate: false,
	})
	if err != nil {
		return nil, err
	}
	return db, nil

}
