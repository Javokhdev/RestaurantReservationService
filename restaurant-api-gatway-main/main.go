package main

import (
	"fmt"
	"log"

	"github.com/dilshodforever/restaurant-submodule/api"
	"github.com/dilshodforever/restaurant-submodule/api/handler"
	pb "github.com/dilshodforever/restaurant-submodule/genprotos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	UserConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()

	UpaymentConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8086"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UpaymentConn.Close()
	us := pb.NewMenuServiceClient(UserConn)
	ps := pb.NewReservationOrderServiceClient(UserConn)
	ca := pb.NewReservationServiceClient(UserConn)
	el := pb.NewRestoranServiceClient(UserConn)
	py := pb.NewPaymentServiceClient(UpaymentConn)

	h := handler.NewHandler(us, ps, ca, el, py)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run()
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
	}
}
