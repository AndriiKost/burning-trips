package api

import (
	"fmt"
	"log"
	"os"

	"github.com/andriikost/burning-gateway/api/controllers"
	"github.com/andriikost/burning-gateway/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	seed.Load(server.DB)

	server.Run(":8081")

}
