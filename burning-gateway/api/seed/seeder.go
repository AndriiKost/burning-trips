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

var stories = []models.Story{
	models.Story{
		Title: "Top 5 Places To Visit in Prague",
		Content: `
		<p>The largest city and the capital in the Czech Republic, Prague is known as one of the most historic and picturesque cities in Europe. This is where a thousand years of history is still breathing perfectly, where unique architectural styles meet in ancient medieval streets to create a picture that is simply unforgettable. Prague is not only one of the most visited destinations in the continent of Europe but the whole world. Let’s focus on top 5 places that will help you to capture the essence of Prague!</p>
		
		<h2>Prague Astronomical Clock</h2>
		
		<p>We have to admit, clocks are not something you would probably imagine to be among the premier attractions within one of the most beautiful cities on the planet, but Prague’s Astronomical Clock is here to surprise you!</p>
		
		<p>It started counting days and hours away from all the way back in 1410 making it the oldest clock in the world that is still operating. The fascinating and majestic Prague Astronomical Clock is definitely one of the top attractions out there in Prague that you can discover during your visit.</p>
		
		<p>To see it in action, find it situated in the charming and cobblestoned Old Town, on the southern wall of the Old Town Hall where the time still ticks away. Be sure to wait for it to strike an hour to watch all 12 apostles come to life with four figures on the sides with each representing things like greed and death.</p>
		
		<h2>Prague Castle Complex</h2>
		
		<p>An unmissable attraction of stunning proportions will catch your eye instantly. Perched on the hill, it is a one of a kind spectacle with nearly 2 million visitors making their way to the largest ancient castle on the planet annually. And you will notice why in a heartbeat. 
		Rooted in rich history and enchanted in almost every architectural style that there is, the medieval masterpiece will unveil fascinating halls, ancient medieval towers, charming gardens, soulful churches, and magnificent palaces.</p>
		
		<p>For a cherry on top, try to be there for a guard changing ceremony that takes place every noon as it will make the experience even more fantastic.</p>
		
		
		<h2>Charles Bridge</h2>
		
		<p>Sometimes hosting massive tourist crowds, don’t let it scare you away from the chance to have another marvelous and atmospheric medieval experience.</p>
		
		<p>Stretching to more than 600 meters, the historic Charles Bridge is definitely everyone’s favorite place to take a slow and relaxing stroll in Prague. It connects Prague Castle and the Old Town gracefully while delivering some of the most picturesque views of the city with Vitava River flowing below.</p>
		
		<p>Walk in the footsteps of history that takes us back all the way to 1357 when the construction of this iconic bridge started with baroque statues and picturesque views of the capital to accompany you along the way.</p>
		
		
		<h2>Petrin Tower</h2>
		
		<p>Take a break from medieval times and history and no better way to do it than with Petrin Tower. Located on the banks of the Vltava River in one of the greenest spaces of Prague, it resembles the iconic Eiffel Tower and reaches the heights of 63.5 meters that will deliver sweeping panoramic views of the city.</p>
		
		<p>Test your legs and fitness while trying to conquer the 299 steps of the tower with stunning rewards waiting at the top. You will also find a gift shop, cafeteria and small exhibition near the tower.
		And if the number of stairs seems like too much, you can always use the elevator and funicular service!</p>
		
		
		<h2>St. Vitus Cathedral</h2>
		
		<p>Striking from outside, and even more mesmerizing from the inside, St. Vitus Cathedral is a must in Prague. 
		Containing tombs of many Bohemian Kings and Holy Roman Emperors, it is not only the largest and most important church in the country but a prime example of Gothic architecture blending with a dash of Baroque and Renaissance.</p>
		
		<p>Founded in 930, it was truly completed just recently, in 1929, with extraordinary sightseeing experience for us to admire.
		You will find incredible art and inspiring mosaics, with architectural details and design that will simply leave you speechless.</p>`,
		ImageUrl: "https://vuejsexamples.com/wysiwyg-editor-for-vuetify-component-simplifies-integration-tiptap-with-vuetify/",
	},
	models.Story{
		Title:    "Some new Story 2",
		Content:  "<p>Some html here 2 blah</p>",
		ImageUrl: "https://vuejsexamples.com/wysiwyg-editor-for-vuetify-component-simplifies-integration-tiptap-with-vuetify/",
	},
}

var storyVotes = []models.StoryVote{
	models.StoryVote{
		Count: 15,
	},
	models.StoryVote{
		Count: 10,
	},
	models.StoryVote{
		Count: 9,
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

	err := db.Debug().DropTableIfExists(&models.StopVote{}, &models.RouteVote{}, &models.StoryVote{}, &models.Stop{}, &models.Route{}, &models.Story{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Stop{}, &models.StopVote{}, &models.Route{}, &models.RouteVote{}, &models.Story{}, &models.StoryVote{}).Error
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

	err = db.Debug().Model(&models.Story{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
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

	err = db.Debug().Model(&models.StoryVote{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key user_id -> users(id) error: %v", err)
	}

	err = db.Debug().Model(&models.StoryVote{}).AddForeignKey("story_id", "stories(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign story_id stories(id) -> key error: %v", err)
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

	for i, _ := range stories {
		stories[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Story{}).Create(&stories[i]).Error
		if err != nil {
			log.Fatalf("cannot seed stories tavle %v", err)
		}

		for j, _ := range storyVotes {
			storyVotes[j].UserID = users[i].ID
			storyVotes[j].StoryID = routes[i].ID
		}

		err = db.Debug().Model(&models.StoryVote{}).Create(&storyVotes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed story votes table: %v", err)
		}
	}
}
