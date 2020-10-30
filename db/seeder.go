package db

import (
	"github.com/haithngn/go-crud/model"
	"gorm.io/gorm"
	"log"
)

type Seeder struct {
}

func (receiver Seeder) Populate(db *gorm.DB) {
	err := db.AutoMigrate(&model.Question{})
	if err != nil {
		log.Fatal(err)
	}
}
