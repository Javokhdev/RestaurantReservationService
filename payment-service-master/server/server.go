package main

import (
	"log"
	"net"
	pb "github.com/dilshodforever/payment-service/genprotos"
	"github.com/dilshodforever/payment-service/service"
	postgres "github.com/dilshodforever/payment-service/storage/postgres"
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
	pb.RegisterPaymentServiceServer(s, service.NewPaymentService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
}
