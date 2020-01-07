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
	models.User{
		Username: "Some User",
		Email:    "someuser@gmail.com",
		Password: "someuser",
	},
	models.User{
		Username: "Some Other User",
		Email:    "someotheruser@gmail.com",
		Password: "someotheruser",
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
	models.Stop{
		Name:       "Other Stop",
		Content:    "Other Stop where people hang out.",
		ImageUrl:   "",
		Address:    "",
		Latitude:   0,
		Longtitude: 0,
	},
	models.Stop{
		Name:       "Some Other Stop",
		Content:    "Some Other Stop where you can find huge amount of fun.",
		ImageUrl:   "",
		Address:    "",
		Latitude:   0,
		Longtitude: 0,
	},
}

var stopVotes = []models.StopVote{
	models.StopVote{
		Count: 15,
	},
	models.StopVote{
		Count: 10,
	},
	models.StopVote{
		Count: 5,
	},
	models.StopVote{
		Count: 12,
	},
}

var routes = []models.Route{
	models.Route{
		ImageUrl: "",
		Name:     "Best Route of all",
		Content:  "Some cool spots are on this route",
	},
	models.Route{
		ImageUrl: "",
		Name:     "Even Better then Best Route of all",
		Content:  "Even Better spots are on this route",
	},
}

var routeVotes = []models.RouteVote{
	models.RouteVote{
		Count: 10,
	},
	models.RouteVote{
		Count: 25,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.StopVote{}, &models.RouteVote{}, &models.Stop{}, &models.Route{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Stop{}, &models.StopVote{}, &models.Route{}, &models.RouteVote{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Stop{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Route{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Route{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.StopVote{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key user_id -> users(id) error: %v", err)
	}

	err = db.Debug().Model(&models.StopVote{}).AddForeignKey("stop_id", "stops(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign stop_id stops(id) -> key error: %v", err)
	}

	err = db.Debug().Model(&models.RouteVote{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key user_id -> users(id) error: %v", err)
	}

	err = db.Debug().Model(&models.RouteVote{}).AddForeignKey("route_id", "routes(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign route_id routes(id) -> key error: %v", err)
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

		for j, _ := range stopVotes {
			stopVotes[j].UserID = users[i].ID
			stopVotes[j].StopID = stops[i].ID
		}

		err = db.Debug().Model(&models.StopVote{}).Create(&stopVotes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed stop votes table: %v", err)
		}
	}

	for i, _ := range routes {
		routes[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Stop{}).Create(&routes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed routes table: %v", err)
		}

		for j, _ := range routeVotes {
			routeVotes[j].UserID = users[i].ID
			routeVotes[j].RouteID = routes[i].ID
		}

		err = db.Debug().Model(&models.RouteVote{}).Create(&routeVotes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed route votes table: %v", err)
		}
	}
}
