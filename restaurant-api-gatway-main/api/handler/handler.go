package handler

import pb "github.com/dilshodforever/restaurant-submodule/genprotos"

type Handler struct {
	Menu             pb.MenuServiceClient
	ReservationOrder pb.ReservationOrderServiceClient
	Reservation      pb.ReservationServiceClient
	Restoran         pb.RestoranServiceClient
	Payment          pb.PaymentServiceClient
}

func NewHandler(me pb.MenuServiceClient, reo pb.ReservationOrderServiceClient,
	re pb.ReservationServiceClient, el pb.RestoranServiceClient, py pb.PaymentServiceClient) *Handler {
	return &Handler{me, reo, re, el, py}
}
