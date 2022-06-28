package database

import (
	"fmt"

	"github.com/shayamvlmna/cab-booking-app/pkg/models"
	"gorm.io/gorm"
)

func OpenTripDb() (*gorm.DB, error) {
	Db, err := openDB()
	if err != nil {
		return nil, err
	}
	trip := &models.Trip{}
	err = Db.AutoMigrate(&trip)
	if err != nil {
		return nil, err
	}
	fmt.Println("trip db opened")
	return Db, nil
}

func GetTrips(id uint64) *[]models.Trip {

	db, _ := OpenTripDb()
	trips := []models.Trip{}
	db.Where("user_id=?", id).Find(&trips)

	return &trips
}
