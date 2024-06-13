package service

import (
	"context"
	"log"
	pb "github.com/dilshodforever/payment-service/genprotos"
	s "github.com/dilshodforever/payment-service/storage"

)

type PaymentService struct {
	stg s.InitRoot
	pb.UnimplementedPaymentServiceServer
}

func NewPaymentService(stg s.InitRoot) *PaymentService {
	return &PaymentService{stg: stg}
}

func (c *PaymentService) CreatePayment(ctx context.Context, Payment *pb.Payment) (*pb.Void, error) {
	pb, err := c.stg.Payment().CreatePayment(Payment)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *PaymentService) GetAllPayments(ctx context.Context, pb *pb.Payment) (*pb.GetAllPayments, error) {
	Payments, err := c.stg.Payment().GetAllPayments(pb)
	if err != nil {
		log.Print(err)
	}

	return Payments, err
}

func (c *PaymentService) GetByIdPayments(ctx context.Context, id *pb.ById) (*pb.Payment, error) {
	pay, err := c.stg.Payment().GetByIdPayment(id)
	if err != nil {
		log.Print(err)
	}

	return pay, err
}

func (c *PaymentService) UpdatePayment(ctx context.Context, Payment *pb.Payment) (*pb.Void, error) {
	pb, err := c.stg.Payment().UpdatePayment(Payment)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *PaymentService) DeletePayment(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.Payment().DeletePayment(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

