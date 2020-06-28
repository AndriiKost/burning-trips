package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Landmark struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:255;not null;" json:"name"`
	Description string    `gorm:"size:255;not null;" json:"description"`
	Image       string    `gorm:"size:255;not null;" json:"image"`
	Latitude    float32   `json:"latitude"`
	Longitude   float32   `json:"longitude"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (landmark *Landmark) Find(db *gorm.DB, sid uint64) (*Landmark, error) {
	var err error
	err = db.Debug().Model(&Landmark{}).Where("id = ?", sid).Take(&landmark).Error
	if err != nil {
		return &Landmark{}, err
	}
	return landmark, nil
}

func (landmark *Landmark) FindAll(db *gorm.DB) (*[]Landmark, error) {
	var err error
	landmarks := []Landmark{}
	err = db.Debug().Model(&Landmark{}).Find(&landmarks).Error
	if err != nil {
		return &[]Landmark{}, err
	}
	return &landmarks, nil
}

func (landmark *Landmark) SearchLandmarks(db *gorm.DB, query SearchQuery) ([]Landmark, error) {
	var landmarks []Landmark

	min_lat, max_lat := getLatMinMax(query)
	min_lng, max_lng := getLngMinMax(query)

	err := db.Model(&Landmark{}).Where("latitude BETWEEN ? AND ?", min_lat, max_lat).Where("longitude BETWEEN ? AND ?", min_lng, max_lng).Find(&landmarks).Error

	if err != nil {
		return landmarks, err
	}

	return landmarks, nil
}

// func getLatMinMax(query SearchQuery) (float32, float32) {
// 	min_lat := query.Latitude - 0.1
// 	max_lat := query.Latitude + 0.1
// 	return min_lat, max_lat
// }
// func getLngMinMax(query SearchQuery) (float32, float32) {
// 	min_lng := query.Longitude - 0.1
// 	max_lng := query.Longitude + 0.1
// 	return min_lng, max_lng
// }
