package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Vote interface {
	UserId() uint32
	VoteCount() uint64
}

type StopVote struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint32    `gorm:"primary_key" json:"userId"`
	StopID    uint64    `gorm:"primary_key" json:"stopId"`
	Count     uint32    `gorm:"default:0" json:"count"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (stopVote *StopVote) Prepare() {
	stopVote.ID = 0
	stopVote.CreatedAt = time.Now()
	stopVote.UpdatedAt = time.Now()
}

func (stopVote *StopVote) Validate() error {
	fmt.Println("Validate ", stopVote)
	if stopVote.UserID < 0 {
		fmt.Println("User ID is required")
		return errors.New("User ID is required")
	}
	if stopVote.StopID < 1 {
		fmt.Println("Stop ID is required")
		return errors.New("Stop ID is required")
	}
	if stopVote.Count < 1 {
		fmt.Println("Vote count can't be lower then 1")
		return errors.New("Vote count can't be lower then 1")
	}
	return nil
}

func (stopVote *StopVote) Find(db *gorm.DB, vid uint64) (*StopVote, error) {
	var err error
	err = db.Debug().Model(&StopVote{}).Where("id = ?", vid).Take(&stopVote).Error
	if err != nil {
		return &StopVote{}, err
	}
	return stopVote, nil
}

func (stopVote *StopVote) FindAll(db *gorm.DB) (*[]StopVote, error) {
	var err error
	stopVotes := []StopVote{}
	err = db.Debug().Model(&StopVote{}).Limit(10000).Find(&stopVotes).Error
	if err != nil {
		return &[]StopVote{}, err
	}
	return &stopVotes, nil
}

func (stopVote *StopVote) FindStopVotes(db *gorm.DB, stopId uint64) (*StopVote, error) {
	var err error
	err = db.Debug().Model(&StopVote{}).Where("stop_id = ?", stopId).Take(&stopVote).Error
	if err != nil {
		return &StopVote{}, err
	}
	return stopVote, nil
}

func (stopVote *StopVote) FindUserStopVotes(db *gorm.DB, stopId uint64, userId uint32) (*StopVote, error) {
	var err error
	err = db.Debug().Model(&StopVote{}).Where("stop_id = ? AND user_id", stopId, userId).Take(&stopVote).Error
	if err != nil {
		return &StopVote{}, err
	}
	return stopVote, nil
}

func (stopVote *StopVote) Update(db *gorm.DB) (*StopVote, error) {
	var err error

	err = db.Debug().Model(&StopVote{}).Create(&stopVote).Error
	if err != nil {
		return &StopVote{}, err
	}
	return stopVote, nil
}

func (stopVote *StopVote) DeleteStopVotes(db *gorm.DB, stopId uint64) (int64, error) {

	db = db.Debug().Model(&StopVote{}).Where("stop_id = ?", stopId).Take(&StopVote{}).Delete(&StopVote{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Stop not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
