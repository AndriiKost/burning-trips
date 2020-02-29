package models

import (
	"errors"
	"html"
	"sort"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Stop struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Address   string     `gorm:"size:255;not null;" json:"address"`
	ImageUrl  string     `gorm:"size:255;not null;" json:"imageUrl"`
	Name      string     `gorm:"size:255;not null;unique" json:"name"`
	Content   string     `gorm:"size:255;not null;" json:"content"`
	Latitude  float32    `json:"latitude"`
	Longitude float32    `json:"longitude"`
	Author    User       `json:"author"`
	Votes     []StopVote `json:"votes"`
	Routes    []Route    `json:"stops"`
	AuthorID  uint32     `gorm:"not null" json:"authorID"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type SearchQuery struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func (stop *Stop) CountVotes() uint32 {
	var count uint32
	count = 0
	for i, _ := range stop.Votes {
		count = count + stop.Votes[i].Count
	}
	return count
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
	stop.Votes = []StopVote{}
	stop.Routes = []Route{}
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
		err = db.Debug().Model(&[]StopVote{}).Where("stop_id = ?", stop.ID).Find(&stop.Votes).Error
		if err != nil {
			return &Stop{}, err
		}
	}
	return stop, nil
}

func (stop *Stop) FindAll(db *gorm.DB) (*[]Stop, error) {
	var err error
	stops := []Stop{}
	err = db.Debug().Model(&Stop{}).Limit(10).Find(&stops).Error
	if err != nil {
		return &[]Stop{}, err
	}
	if len(stops) > 0 {
		for i, _ := range stops {
			err := db.Debug().Model(&User{}).Where("id = ?", stops[i].AuthorID).Take(&stops[i].Author).Error
			if err != nil {
				return &[]Stop{}, err
			}
			err = db.Debug().Model(&StopVote{}).Where("stop_id = ?", stops[i].ID).Find(&stops[i].Votes).Error
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
		err = db.Debug().Model(&[]StopVote{}).Where("stop_id = ?", stop.ID).Find(&stop.Votes).Error
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
		err = db.Debug().Model(&StopVote{}).Where("stop_id = ?", stop.ID).Find(&stop.Votes).Error
		if err != nil {
			return &Stop{}, err
		}
	}
	return stop, nil
}

func (stop *Stop) Delete(db *gorm.DB, sid uint64) (int64, error) {

	db = db.Debug().Model(&Stop{}).Where("id = ?", sid).Take(&Stop{}).Delete(&Stop{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Stop not found")
		}
		return 0, db.Error
	}

	db = db.Debug().Model(&StopVote{}).Where("stop_id = ?", sid).Find(&StopVote{}).Delete(&StopVote{})
	return db.RowsAffected, nil
}

func (stop *Stop) SearchStops(db *gorm.DB, query SearchQuery) ([]Stop, error) {
	var stops []Stop

	min_lat, max_lat := getLatMinMax(query)
	min_lng, max_lng := getLngMinMax(query)

	err := db.Model(&Stop{}).Where("latitude BETWEEN ? AND ?", min_lat, max_lat).Where("longitude BETWEEN ? AND ?", min_lng, max_lng).Find(&stops).Error

	if err != nil {
		return stops, err
	}

	if len(stops) > 0 {
		users := []User{}
		stopVotes := []StopVote{}

		err := db.Debug().Model(&User{}).Find(&users).Error
		if err != nil {
			return stops, err
		}

		err = db.Debug().Model(&StopVote{}).Find(&stopVotes).Error
		if err != nil {
			return stops, err
		}

		for i, _ := range stops {
			for u, _ := range users {
				if stops[i].AuthorID == users[u].ID {
					stops[i].Author = users[u]
				}
			}

			for j, _ := range stopVotes {
				if stops[i].ID == stopVotes[j].StopID {
					stops[i].Votes = append(stops[i].Votes, stopVotes[j])
				}
			}
		}
	}

	sort.Slice(stops, func(i, j int) bool {
		votesA := stops[i].CountVotes()
		votesB := stops[j].CountVotes()
		return votesA > votesB
	})

	return stops, nil
}

func getLatMinMax(query SearchQuery) (float32, float32) {
	min_lat := query.Latitude - 0.1
	max_lat := query.Latitude + 0.1
	return min_lat, max_lat
}
func getLngMinMax(query SearchQuery) (float32, float32) {
	min_lng := query.Longitude - 0.1
	max_lng := query.Longitude + 0.1
	return min_lng, max_lng
}
