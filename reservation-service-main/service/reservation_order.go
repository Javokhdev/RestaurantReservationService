package service

import (
	"context"
	"log"
	pb "github.com/dilshodforever/reservation-service/genprotos"
	s "github.com/dilshodforever/reservation-service/storage"

)

type Reservation_orderService struct {
	stg s.InitRoot
	pb.UnimplementedReservationOrderServiceServer
}

func NewReservation_orderService(stg s.InitRoot) *Reservation_orderService {
	return &Reservation_orderService{stg: stg}
}

func (c *Reservation_orderService) CreateReservation(ctx context.Context, Reservation *pb.Reservation) (*pb.Void, error) {
	pb, err := c.stg.Reservation().CreateReservation(Reservation)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *Reservation_orderService) GetAllReservations(ctx context.Context, pb *pb.Reservation) (*pb.GetAllReservations, error) {
	Reservations, err := c.stg.Reservation().GetAllReservation(pb)
	if err != nil {
		log.Print(err)
	}

	return Reservations, err
}

func (c *Reservation_orderService) GetByIdReservation(ctx context.Context, id *pb.ById) (*pb.Reservation, error) {
	prod, err := c.stg.Reservation().GetByIdReservation(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *Reservation_orderService) UpdateReservation(ctx context.Context, Reservation *pb.Reservation) (*pb.Void, error) {
	pb, err := c.stg.Reservation().UpdateReservation(Reservation)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *Reservation_orderService) DeleteReservation(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.Reservation().DeleteReservation(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

