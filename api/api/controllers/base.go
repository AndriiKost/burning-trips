package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/etherlabsio/healthcheck"
	"github.com/etherlabsio/healthcheck/checkers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/andriikost/burning-gateway/api/models"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"github.com/rs/cors"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Stop{}) //database migration

	server.Router = mux.NewRouter()

	server.Router.Handle("/healthcheck", healthcheck.Handler(

		// WithTimeout allows you to set a max overall timeout.
		healthcheck.WithTimeout(5*time.Second),

		// Checkers fail the status in case of any error.
		healthcheck.WithChecker(
			"heartbeat", checkers.Heartbeat("/Users/andrii/Projects/burning-trips/api/heartbeat"),
		),

		// healthcheck.WithChecker(
		// 	"database", healthcheck.CheckerFunc(
		// 		func(ctx context.Context) error {
		// 			return server.DB.PingContext(ctx)
		// 		},
		// 	),
		// ),

		// Observers do not fail the status in case of error.
		healthcheck.WithObserver(
			"diskspace", checkers.DiskSpace("/var/log", 90),
		),
	))

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://192.168.1.219:8080"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowCredentials: true,
		Debug:            true,
		AllowedHeaders:   []string{"Authorization", "Content-Type", "X-Requested-With"},
	})

	handler := c.Handler(server.Router)

	fmt.Println("Listening to port 8081")
	log.Fatal(http.ListenAndServe(addr, handler))
}
