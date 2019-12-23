package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Stop struct {
	ID         uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Address    string     `gorm:"size:255;not null;" json:"address"`
	ImageUrl   string     `gorm:"size:255;not null;" json:"imageUrl"`
	Name       string     `gorm:"size:255;not null;unique" json:"name"`
	Content    string     `gorm:"size:255;not null;" json:"content"`
	Latitude   float32    `json:"latitude"`
	Longtitude float32    `json:"longtitude"`
	Author     User       `json:"author"`
	Votes      []StopVote `json:"votes"`
	AuthorID   uint32     `gorm:"not null" json:"authorID"`
	CreatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (stop *Stop) Prepare() {
	stop.ID = 0
	stop.Name = html.EscapeString(strings.TrimSpace(stop.Name))
	stop.Content = html.EscapeString(strings.TrimSpace(stop.Content))
	stop.ImageUrl = html.EscapeString(strings.TrimSpace(stop.ImageUrl))
	stop.Author = User{}
	stop.Address = html.EscapeString(strings.TrimSpace(stop.Address))
	stop.CreatedAt = time.Now()
	stop.UpdatedAt = time.Now()
}

func (stop *Stop) Validate() error {

	if stop.Name == "" {
		return errors.New("Required Name")
	}
	if stop.Content == "" {
		return errors.New("Required Content")
	}
	if stop.AuthorID < 1 {
		return errors.New("Required Author")
	}
	return nil
}

func (stop *Stop) Find(db *gorm.DB, sid uint64) (*Stop, error) {
	var err error
	err = db.Debug().Model(&Stop{}).Where("id = ?", sid).Take(&stop).Error
	if err != nil {
		return &Stop{}, err
	}
	if stop.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", stop.AuthorID).Take(&stop.Author).Error
		if err != nil {
			return &Stop{}, err
		}
	}
	return stop, nil
}

func (stop *Stop) FindAll(db *gorm.DB) (*[]Stop, error) {
	var err error
	stops := []Stop{}
	err = db.Debug().Model(&Stop{}).Limit(100).Find(&stops).Error
	if err != nil {
		return &[]Stop{}, err
	}
	if len(stops) > 0 {
		for i, _ := range stops {
			err := db.Debug().Model(&User{}).Where("id = ?", stops[i].AuthorID).Take(&stops[i].Author).Error
			if err != nil {
				return &[]Stop{}, err
			}
		}
	}
	return &stops, nil
}

func (stop *Stop) Update(db *gorm.DB) (*Stop, error) {

	var err error

	err = db.Debug().Model(&Stop{}).Where("id = ?", stop.ID).Updates(
		Stop{
			Name:      stop.Name,
			Content:   stop.Content,
			Address:   stop.Address,
			ImageUrl:  stop.ImageUrl,
			UpdatedAt: time.Now(),
		}).Error

	if err != nil {
		return &Stop{}, err
	}
	if stop.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", stop.AuthorID).Take(&stop.Author).Error
		if err != nil {
			return &Stop{}, err
		}
	}
	return stop, nil
}

func (stop *Stop) Save(db *gorm.DB) (*Stop, error) {
	var err error
	err = db.Debug().Model(&Stop{}).Create(&stop).Error
	if err != nil {
		return &Stop{}, err
	}
	if stop.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", stop.AuthorID).Take(&stop.Author).Error
		if err != nil {
			return &Stop{}, err
		}
	}
	return stop, nil
}

func (stop *Stop) Delete(db *gorm.DB, sid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Stop{}).Where("id = ? and author_id = ?", sid, uid).Take(&Stop{}).Delete(&Stop{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Stop not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
