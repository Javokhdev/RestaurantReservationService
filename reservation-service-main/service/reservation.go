package service

import (
	"context"
	"log"
	pb "github.com/dilshodforever/reservation-service/genprotos"
	s "github.com/dilshodforever/reservation-service/storage"

)

type ReservationService struct {
	stg s.InitRoot
	pb.UnimplementedReservationServiceServer
}

func NewReservationService(stg s.InitRoot) *ReservationService {
	return &ReservationService{stg: stg}
}

func (c *ReservationService) CreateReservation(ctx context.Context, Reservation *pb.Reservation) (*pb.Void, error) {
	pb, err := c.stg.Reservation().CreateReservation(Reservation)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *ReservationService) GetAllReservation(ctx context.Context, pb *pb.Reservation) (*pb.GetAllReservations, error) {
	Reservations, err := c.stg.Reservation().GetAllReservation(pb)
	if err != nil {
		log.Print(err)
	}

	return Reservations, err
}

func (c *ReservationService) GetByIdReservation(ctx context.Context, id *pb.ById) (*pb.Reservation, error) {
	prod, err := c.stg.Reservation().GetByIdReservation(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *ReservationService) UpdateReservation(ctx context.Context, Reservation *pb.Reservation) (*pb.Void, error) {
	pb, err := c.stg.Reservation().UpdateReservation(Reservation)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *ReservationService) DeleteReservation(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.Reservation().DeleteReservation(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}



func (c *ReservationService) Reservations(ctx context.Context, Reservation *pb.Reservation) (*pb.Void, error) {
	pb, err := c.stg.Reservation().Reservations(Reservation)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}