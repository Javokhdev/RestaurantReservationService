package postgres

import (
	pb "github.com/dilshodforever/payment-service/genprotos"
)

type InitRoot interface {
	Payment() Payment
}
type Payment interface {
	CreatePayment(Payment *pb.Payment) (*pb.Void, error)
	GetByIdPayment(id *pb.ById) (*pb.Payment, error)
	GetAllPayments(_ *pb.Payment) (*pb.GetAllPayments, error)
	UpdatePayment(Payment *pb.Payment) (*pb.Void, error)
	DeletePayment(id *pb.ById) (*pb.Void, error)
}

