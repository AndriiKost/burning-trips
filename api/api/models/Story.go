package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Story struct {
	ID       uint64      `gorm:"primary_key;auto_increment" json:"id"`
	ImageUrl string      `gorm:"size:255;not null;" json:"imageUrl"`
	Title    string      `gorm:"size:255;not null;" json:"title"`
	Content  string      `gorm:"size:max;not null;" json:"content"`
	Author   User        `json:"author"`
	Votes    []StoryVote `json:"votes"`
	// Stops     []Stop      `json:"stops"`
	AuthorID  uint32    `gorm:"not null" json:"authorID"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TODO: Store Story Content as  []Section{}
type Section struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	ImageUrl  string `gorm:"size:255" json:"imageUrl"`
	Title     string `gorm:"size:255" json:"title"`
	Paragraph string `gorm:"size:max;non null;" json:"content"`
}

func (story *Story) Prepare() {
	story.ID = 0
	story.ImageUrl = html.EscapeString(strings.TrimSpace(story.ImageUrl))
	story.Title = html.EscapeString(strings.TrimSpace(story.Title))
	story.Content = strings.TrimSpace(story.Content)
	story.Author = User{}
	// story.Stops = []Stop{}
	story.Votes = []StoryVote{}
	story.CreatedAt = time.Now()
	story.UpdatedAt = time.Now()
}

func (story *Story) Validate() error {

	if story.Title == "" {
		return errors.New("Required Title")
	}
	if story.Content == "" {
		return errors.New("Required Content")
	}
	if story.AuthorID < 1 {
		return errors.New("Required Author")
	}
	return nil
}

func (story *Story) Find(db *gorm.DB, sid uint64) (*Story, error) {
	var err error
	err = db.Debug().Model(&Story{}).Where("id = ?", sid).Take(&story).Error
	if err != nil {
		return &Story{}, err
	}
	if story.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", story.AuthorID).Take(&story.Author).Error
		if err != nil {
			return &Story{}, err
		}
		err = db.Debug().Model(&[]StoryVote{}).Where("story_id = ?", story.ID).Find(&story.Votes).Error
		if err != nil {
			return &Story{}, err
		}
	}
	return story, nil
}

func (story *Story) FindAll(db *gorm.DB) (*[]Story, error) {
	var err error
	stories := []Story{}
	err = db.Model(&Story{}).Limit(100).Order("updated_at desc").Find(&stories).Error
	if err != nil {
		return &[]Story{}, err
	}
	if len(stories) > 0 {
		users := []User{}
		storyVotes := []StoryVote{}

		err := db.Debug().Model(&User{}).Find(&users).Error
		if err != nil {
			return &[]Story{}, err
		}

		err = db.Debug().Model(&StoryVote{}).Find(&storyVotes).Error
		if err != nil {
			return &[]Story{}, err
		}

		for i, _ := range stories {
			for u, _ := range users {
				if stories[i].AuthorID == users[u].ID {
					stories[i].Author = users[u]
				}
			}

			for j, _ := range storyVotes {
				if stories[i].ID == storyVotes[j].StoryID {
					stories[i].Votes = append(stories[i].Votes, storyVotes[j])
				}
			}
		}
	}

	return &stories, nil
}

func (story *Story) Update(db *gorm.DB) (*Story, error) {

	var err error

	err = db.Debug().Model(&Story{}).Where("id = ?", story.ID).Updates(
		Story{
			Title:     story.Title,
			Content:   story.Content,
			ImageUrl:  story.ImageUrl,
			UpdatedAt: time.Now(),
		}).Error

	if err != nil {
		return &Story{}, err
	}
	if story.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", story.AuthorID).Take(&story.Author).Error
		if err != nil {
			return &Story{}, err
		}
		err = db.Debug().Model(&[]StoryVote{}).Where("story_id = ?", story.ID).Find(&story.Votes).Error
		if err != nil {
			return &Story{}, err
		}
	}
	return story, nil
}

func (story *Story) Create(db *gorm.DB) (*Story, error) {
	var err error
	err = db.Debug().Model(&Story{}).Create(&story).Error
	if err != nil {
		return &Story{}, err
	}
	if story.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", story.AuthorID).Take(&story.Author).Error
		if err != nil {
			return &Story{}, err
		}
		err = db.Debug().Model(&StoryVote{}).Where("story_id = ?", story.ID).Find(&story.Votes).Error
		if err != nil {
			return &Story{}, err
		}
	}
	return story, nil
}

func (story *Story) Delete(db *gorm.DB, sid uint64) (int64, error) {

	db = db.Debug().Model(&Story{}).Where("id = ?", sid).Take(&Story{}).Delete(&Story{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Story not found")
		}
		return 0, db.Error
	}

	db = db.Debug().Model(&StoryVote{}).Where("story_id = ?", sid).Find(&StoryVote{}).Delete(&StoryVote{})
	return db.RowsAffected, nil
}
