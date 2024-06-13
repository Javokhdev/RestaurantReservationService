package main

import (
	"log"
	"net"

	pb "github.com/dilshodforever/reservation-service/genprotos"
	"github.com/dilshodforever/reservation-service/service"
	postgres "github.com/dilshodforever/reservation-service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	liss, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterRestoranServiceServer(s, service.NewRestoranService(db))
	pb.RegisterReservationServiceServer(s, service.NewReservationService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
