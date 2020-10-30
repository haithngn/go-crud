package db

import (
	"fmt"
	"github.com/haithngn/go-crud/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start(config *configuration.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database on %s", config.Name, config.Host)
		return nil, err
	}

	return db, nil
}
