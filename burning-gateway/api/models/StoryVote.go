package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type StoryVote struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint32    `gorm:"primary_key" json:"userId"`
	StoryID   uint64    `gorm:"primary_key" json:"storyId"`
	Count     uint32    `gorm:"default:0" json:"count"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (storyVote *StoryVote) Prepare() {
	storyVote.ID = 0
	storyVote.CreatedAt = time.Now()
	storyVote.UpdatedAt = time.Now()
}

func (storyVote *StoryVote) Validate() error {
	fmt.Println("Validate ", storyVote)
	if storyVote.UserID < 0 {
		fmt.Println("User ID is required")
		return errors.New("User ID is required")
	}
	if storyVote.StoryID < 1 {
		fmt.Println("Story ID is required")
		return errors.New("Story ID is required")
	}
	if storyVote.Count < 1 {
		fmt.Println("Vote count can't be lower then 1")
		return errors.New("Vote count can't be lower then 1")
	}
	return nil
}

func (storyVote *StoryVote) Find(db *gorm.DB, vid uint64) (*StoryVote, error) {
	var err error
	err = db.Debug().Model(&StoryVote{}).Where("id = ?", vid).Take(&storyVote).Error
	if err != nil {
		return &StoryVote{}, err
	}
	return storyVote, nil
}

func (storyVote *StoryVote) FindAll(db *gorm.DB) (*[]StoryVote, error) {
	var err error
	storyVotes := []StoryVote{}
	err = db.Debug().Model(&StoryVote{}).Limit(10000).Find(&storyVotes).Error
	if err != nil {
		return &[]StoryVote{}, err
	}
	return &storyVotes, nil
}

func (storyVote *StoryVote) FindStoryVotes(db *gorm.DB, storyId uint64) (*StoryVote, error) {
	var err error
	err = db.Debug().Model(&StoryVote{}).Where("story_id = ?", storyId).Take(&storyVote).Error
	if err != nil {
		return &StoryVote{}, err
	}
	return storyVote, nil
}

func (storyVote *StoryVote) FindUserStoryVotes(db *gorm.DB, storyId uint64, userId uint32) (*StoryVote, error) {
	var err error
	err = db.Debug().Model(&StoryVote{}).Where("story_id = ? AND user_id", storyId, userId).Take(&storyVote).Error
	if err != nil {
		return &StoryVote{}, err
	}
	return storyVote, nil
}

func (storyVote *StoryVote) SaveStoryVote(db *gorm.DB) (*StoryVote, error) {
	var err error

	if storyVote.ID != 0 {
		err = db.Debug().Model(&StoryVote{}).Update(&storyVote).Error
	} else {
		storyVote.Prepare()
		err = db.Debug().Model(&StoryVote{}).Create(&storyVote).Error
	}

	if err != nil {
		return &StoryVote{}, err
	}
	return storyVote, nil
}

func (storyVote *StoryVote) DeleteStoryVotes(db *gorm.DB, storyId uint64) (int64, error) {

	db = db.Debug().Model(&StoryVote{}).Where("story_id = ?", storyId).Take(&StoryVote{}).Delete(&StoryVote{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Story not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
