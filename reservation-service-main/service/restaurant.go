package service

import (
	"context"
	"log"

	pb "github.com/dilshodforever/reservation-service/genprotos"
	s "github.com/dilshodforever/reservation-service/storage"
)

type RestoranService struct {
	stg s.InitRoot
	pb.UnimplementedRestoranServiceServer
}

func NewRestoranService(stg s.InitRoot) *RestoranService {
	return &RestoranService{stg: stg}
}

func (c *RestoranService) CreateRestoran(ctx context.Context, Restoran *pb.Restoran) (*pb.Void, error) {
	pb, err := c.stg.Restoran().CreateRestoran(Restoran)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *RestoranService) GetAllRestoran(ctx context.Context, pb *pb.Restoran) (*pb.GetAllRestorans, error) {
	Restorans, err := c.stg.Restoran().GetAllRestoran(pb)
	if err != nil {
		log.Print(err)
	}

	return Restorans, err
}

func (c *RestoranService) GetByIdRestoran(ctx context.Context, id *pb.ById) (*pb.Restoran, error) {
	prod, err := c.stg.Restoran().GetByIdRestoran(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *RestoranService) UpdateRestoran(ctx context.Context, Restoran *pb.Restoran) (*pb.Void, error) {
	pb, err := c.stg.Restoran().UpdateRestoran(Restoran)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *RestoranService) DeleteRestoran(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.Restoran().DeleteRestoran(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}
