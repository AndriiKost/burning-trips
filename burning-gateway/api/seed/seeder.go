package seed

import (
	"log"

	"github.com/andriikost/burning-gateway/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Username: "Andrii Kost",
		Email:    "admin@gmail.com",
		Password: "password",
	},
	models.User{
		Username: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
}

var stops = []models.Stop{
	models.Stop{
		Name:       "Madison State Capitol",
		Content:    "The closest capitol design to the Washington DC one.",
		ImageUrl:   "https://images.unsplash.com/photo-1496144300411-8dd31ce145ba?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=900&q=60",
		Address:    "2 E Main St, Madison, WI 53703",
		Latitude:   0,
		Longtitude: 0,
	},
	models.Stop{
		Name:       "Madison State Street",
		Content:    "Street where you can find all kinds of entertainment.",
		ImageUrl:   "https://architectureunderdevelopment.files.wordpress.com/2011/06/middle-of-state-street.jpg",
		Address:    "State Street, Madisom, WI 53703",
		Latitude:   0,
		Longtitude: 0,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Stop{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Stop{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Stop{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		stops[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Stop{}).Create(&stops[i]).Error
		if err != nil {
			log.Fatalf("cannot seed stops table: %v", err)
		}
	}
}
