package postgres

import (
	pb "github.com/dilshodforever/reservation-service/genprotos"
)

type InitRoot interface {
	Restoran() Restoran
	Reservation() Reservation
	ReservationOrder() ReservationOrder
	Menu() Menu
}
type Restoran interface {
	CreateRestoran(Restoran *pb.Restoran) (*pb.Void, error)
	GetByIdRestoran(id *pb.ById) (*pb.Restoran, error)
	GetAllRestoran(_ *pb.Restoran) (*pb.GetAllRestorans, error)
	UpdateRestoran(Restoran *pb.Restoran) (*pb.Void, error)
	DeleteRestoran(id *pb.ById) (*pb.Void, error)
}

type Reservation interface {
	CreateReservation(Reservation *pb.Reservation) (*pb.Void, error)
	GetByIdReservation(id *pb.ById) (*pb.Reservation, error)
	GetAllReservation(_ *pb.Reservation) (*pb.GetAllReservations, error)
	UpdateReservation(Reservation *pb.Reservation) (*pb.Void, error)
	DeleteReservation(id *pb.ById) (*pb.Void, error)
	Reservations(Reservation *pb.Reservation) (*pb.Void, error)
}

type ReservationOrder interface {
	CreateReservationOrder(ReservationOrder *pb.ReservationOrder) (*pb.Void, error)
	GetByIdReservationOrder(id *pb.ById) (*pb.ReservationOrder, error)
	GetAllReservationOrder(_ *pb.ReservationOrder) (*pb.GetAllReservationOrders, error)
	UpdateReservationOrder(ReservationOrder *pb.ReservationOrder) (*pb.Void, error)
	DeleteReservationOrder(id *pb.ById) (*pb.Void, error)
}

type Menu interface {
	CreateMenu(Menu *pb.Menu) (*pb.Void, error)
	GetByIdMenu(id *pb.ById) (*pb.Menu, error)
	GetAllMenu(_ *pb.Menu) (*pb.GetAllMenus, error)
	UpdateMenu(Menu *pb.Menu) (*pb.Void, error)
	DeleteMenu(id *pb.ById) (*pb.Void, error)
}
