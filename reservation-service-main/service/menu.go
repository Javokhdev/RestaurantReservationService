package service

import (
	"context"
	"log"
	pb "github.com/dilshodforever/reservation-service/genprotos"
	s "github.com/dilshodforever/reservation-service/storage"

)

type MenuService struct {
	stg s.InitRoot
	pb.UnimplementedMenuServiceServer
}

func NewMenuService(stg s.InitRoot) *MenuService {
	return &MenuService{stg: stg}
}

func (c *MenuService) CreateMenu(ctx context.Context, Menu *pb.Menu) (*pb.Void, error) {
	pb, err := c.stg.Menu().CreateMenu(Menu)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *MenuService) GetAllMenus(ctx context.Context, pb *pb.Menu) (*pb.GetAllMenus, error) {
	Menus, err := c.stg.Menu().GetAllMenu(pb)
	if err != nil {
		log.Print(err)
	}

	return Menus, err
}

func (c *MenuService) GetByIdMenu(ctx context.Context, id *pb.ById) (*pb.Menu, error) {
	prod, err := c.stg.Menu().GetByIdMenu(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *MenuService) UpdateMenu(ctx context.Context, Menu *pb.Menu) (*pb.Void, error) {
	pb, err := c.stg.Menu().UpdateMenu(Menu)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *MenuService) DeleteMenu(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.Menu().DeleteMenu(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

