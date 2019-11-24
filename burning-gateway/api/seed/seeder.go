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

	// {
	// 	uid: 'sajkdajknsdknjsa',
	// 	name: "Stop Name",
	// 	authorID: "asjdkjasdaklsjdklj12",
	// 	content: "Such a cool place omg lol omg lorem ispum ups dolor set isjask!",
	// 	trending: true,
	// 	address: {
	// 		state: 'SA',
	// 		address: '1 Capitol Square',
	// 		zipcode: 53701,
	// 		country: 'USA',
	// 		city: 'Madison'
	// 	},
	// 	coords: {
	// 		lat: 8912381297398127938721,
	// 		lng: 128938912398781293871293
	// 	},
	// 	votes: [
	// 		{ uid: 'sakldjaskljdklsa', userVotes: 59, userID: 'ahsdjkhaskjdhjkashdjkhas' }
	// 	],
	// 	imageUrl: 'https://images.unsplash.com/photo-1572295727871-7638149ea3d7?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=500&q=60'
	// },

	// {
	// 	uid: 'sdasadasdasdass',
	// 	name: "Madison Stop",
	// 	authorID: "sheila shasl222",
	// 	content: "Super duper awesome place!",
	// 	trending: false,
	// 	address: {
	// 		state: 'SA',
	// 		address: '1 Capitol Square',
	// 		zipcode: 53701,
	// 		country: 'USA',
	// 		city: 'Madison'
	// 	},
	// 	coords: {
	// 		lat: 8912381297398127938721,
	// 		lng: 128938912398781293871293
	// 	},
	// 	votes: [
	// 		{ uid: 'sakldjaskljdklsa', userVotes: 19, userID: '213asdajksh' }
	// 	],
	// 	imageUrl: 'https://images.unsplash.com/photo-1561953131-4b07835d56e8?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1351&q=80'
	// },

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
