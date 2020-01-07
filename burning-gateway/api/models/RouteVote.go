package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type RouteVote struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	RouteID   uint64    `gorm:"primary_key" json:"routeId"`
	UserID    uint32    `gorm:"primary_key" json:"userId"`
	Count     uint32    `gorm:"default:0" json:"count"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (routeVote *RouteVote) Prepare() {
	routeVote.ID = 0
	routeVote.CreatedAt = time.Now()
	routeVote.UpdatedAt = time.Now()
}

func (routeVote *RouteVote) Validate() error {
	fmt.Println("Validate ", routeVote)
	if routeVote.UserID < 0 {
		fmt.Println("User ID is required")
		return errors.New("User ID is required")
	}
	if routeVote.RouteID < 1 {
		fmt.Println("Route ID is required")
		return errors.New("Route ID is required")
	}
	if routeVote.Count < 1 {
		fmt.Println("Vote count can't be lower then 1")
		return errors.New("Vote count can't be lower then 1")
	}
	return nil
}

func (routeVote *RouteVote) Find(db *gorm.DB, vid uint64) (*RouteVote, error) {
	var err error
	err = db.Debug().Model(&RouteVote{}).Where("id = ?", vid).Take(&routeVote).Error
	if err != nil {
		return &RouteVote{}, err
	}
	return routeVote, nil
}

func (routeVote *RouteVote) FindAll(db *gorm.DB) (*[]RouteVote, error) {
	var err error
	routeVotes := []RouteVote{}
	err = db.Debug().Model(&RouteVote{}).Limit(10000).Find(&routeVotes).Error
	if err != nil {
		return &[]RouteVote{}, err
	}
	return &routeVotes, nil
}

func (routeVote *RouteVote) FindRouteVotes(db *gorm.DB, routeId uint64) (*RouteVote, error) {
	var err error
	err = db.Debug().Model(&RouteVote{}).Where("route_id = ?", routeId).Take(&routeVote).Error
	if err != nil {
		return &RouteVote{}, err
	}
	return routeVote, nil
}

func (routeVote *RouteVote) FindUserRouteVotes(db *gorm.DB, routeId uint64, userId uint32) (*RouteVote, error) {
	var err error
	err = db.Debug().Model(&RouteVote{}).Where("route_id = ? AND user_id", routeId, userId).Take(&routeVote).Error
	if err != nil {
		return &RouteVote{}, err
	}
	return routeVote, nil
}

func (routeVote *RouteVote) SaveRouteVote(db *gorm.DB) (*RouteVote, error) {
	var err error

	if routeVote.ID != 0 {
		err = db.Debug().Model(&RouteVote{}).Update(&routeVote).Error
	} else {
		routeVote.Prepare()
		err = db.Debug().Model(&RouteVote{}).Create(&routeVote).Error
	}

	if err != nil {
		return &RouteVote{}, err
	}
	return routeVote, nil
}

func (routeVote *RouteVote) DeleteRouteVotes(db *gorm.DB, routeId uint64) (int64, error) {

	db = db.Debug().Model(&RouteVote{}).Where("route_id = ?", routeId).Take(&RouteVote{}).Delete(&RouteVote{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Route not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
