package main

import (
	"fmt"
	"log"

	"github.com/dilshodforever/restaurant-auth/api"
	"github.com/dilshodforever/restaurant-auth/api/handler"
	pb "github.com/dilshodforever/restaurant-auth/service"
	"github.com/dilshodforever/restaurant-auth/storage/postgres"
)

func main() {
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	us := pb.NewUserService(db)

	h := handler.NewHandler(us)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8081")
	err = r.Run(":8081")
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
	}
}
