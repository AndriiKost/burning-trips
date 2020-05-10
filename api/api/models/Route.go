package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Route struct {
	ID        uint64      `gorm:"primary_key;auto_increment" json:"id"`
	ImageUrl  string      `gorm:"size:255;not null;" json:"imageUrl"`
	Name      string      `gorm:"size:255;not null;unique" json:"name"`
	Content   string      `gorm:"size:255;not null;" json:"content"`
	Author    User        `json:"author"`
	Votes     []RouteVote `json:"votes"`
	Stops     []Stop      `gorm:"many2many:stop_route;" json:"stops"`
	AuthorID  uint32      `gorm:"not null" json:"authorID"`
	CreatedAt time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (route *Route) Prepare() {
	route.ID = 0
	route.Name = html.EscapeString(strings.TrimSpace(route.Name))
	route.Content = html.EscapeString(strings.TrimSpace(route.Content))
	route.ImageUrl = html.EscapeString(strings.TrimSpace(route.ImageUrl))
	route.Author = User{}
	route.Stops = []Stop{}
	route.CreatedAt = time.Now()
	route.UpdatedAt = time.Now()
	route.Votes = []RouteVote{}
}

func (route *Route) Validate() error {

	if route.Name == "" {
		return errors.New("Required Name")
	}
	if route.Content == "" {
		return errors.New("Required Content")
	}
	if route.AuthorID < 1 {
		return errors.New("Required Author")
	}
	if len(route.Stops) < 2 {
		return errors.New("At least 2 stops are required")
	}
	return nil
}

func (route *Route) Find(db *gorm.DB, sid uint64) (*Route, error) {
	var err error
	err = db.Debug().Model(&Route{}).Where("id = ?", sid).Take(&route).Error
	if err != nil {
		return &Route{}, err
	}
	if route.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", route.AuthorID).Take(&route.Author).Error
		if err != nil {
			return &Route{}, err
		}
		err = db.Debug().Model(&[]RouteVote{}).Where("route_id = ?", route.ID).Find(&route.Votes).Error
		if err != nil {
			return &Route{}, err
		}
		err = db.Debug().Model(&route).Association("Stops").Find(&route.Stops).Error
		if err != nil {
			return &Route{}, err
		}
	}
	return route, nil
}

func (route *Route) FindAll(db *gorm.DB) (*[]Route, error) {
	var err error
	var routes []Route
	err = db.Debug().Model(&Route{}).Limit(100).Find(&routes).Error
	if err != nil {
		return &[]Route{}, err
	}
	if len(routes) > 0 {
		for i, _ := range routes {
			err := db.Debug().Model(&User{}).Where("id = ?", routes[i].AuthorID).Take(&routes[i].Author).Error
			if err != nil {
				return &[]Route{}, err
			}
			err = db.Debug().Model(&RouteVote{}).Where("route_id = ?", routes[i].ID).Find(&routes[i].Votes).Error
			if err != nil {
				return &[]Route{}, err
			}
			err = db.Debug().Model(&routes[i]).Related(&routes[i].Stops, "Stops").Error
			if err != nil {
				return &[]Route{}, err
			}
			if len(routes[i].Stops) > 0 {
				for j, _ := range routes[i].Stops {
					err = db.Debug().Model(&User{}).Where("id = ?", routes[i].Stops[j].AuthorID).Take(&routes[i].Stops[j].Author).Error
					if err != nil {
						return &[]Route{}, err
					}
					err = db.Debug().Model(&[]StopVote{}).Where("stop_id = ?", routes[i].Stops[j].ID).Find(&routes[i].Stops[j].Votes).Error
					if err != nil {
						return &[]Route{}, err
					}
				}
			}
		}
	}
	return &routes, nil
}

func (route *Route) Update(db *gorm.DB) (*Route, error) {

	var err error

	err = db.Debug().Model(&Stop{}).Where("id = ?", route.ID).Updates(
		Route{
			Name:      route.Name,
			Content:   route.Content,
			ImageUrl:  route.ImageUrl,
			UpdatedAt: time.Now(),
		}).Error

	if err != nil {
		return &Route{}, err
	}
	if route.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", route.AuthorID).Take(&route.Author).Error
		if err != nil {
			return &Route{}, err
		}
		err = db.Debug().Model(&[]RouteVote{}).Where("route_id = ?", route.ID).Find(&route.Votes).Error
		if err != nil {
			return &Route{}, err
		}
	}
	return route, nil
}

func (route *Route) Create(db *gorm.DB) (*Route, error) {
	var err error
	err = db.Debug().Model(&Route{}).Create(&route).Error
	if err != nil {
		return &Route{}, err
	}
	if route.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", route.AuthorID).Take(&route.Author).Error
		if err != nil {
			return &Route{}, err
		}
		err = db.Debug().Model(&RouteVote{}).Where("route_id = ?", route.ID).Find(&route.Votes).Error
		if err != nil {
			return &Route{}, err
		}
	}
	return route, nil
}

func (route *Route) Delete(db *gorm.DB, rid uint64) (int64, error) {

	db = db.Debug().Model(&Route{}).Where("id = ?", rid).Take(&Route{}).Delete(&Route{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Route not found")
		}
		return 0, db.Error
	}

	db = db.Debug().Model(&RouteVote{}).Where("route_id = ?", rid).Find(&RouteVote{}).Delete(&RouteVote{})
	return db.RowsAffected, nil
}
