package service

import (
	"context"
	"log"
	pb "github.com/dilshodforever/restaurant-auth/genprotos"
	s "github.com/dilshodforever/restaurant-auth/storage"

)

type UserService struct {
	stg s.InitRoot
}

func NewUserService(stg s.InitRoot) *UserService {
	return &UserService{stg: stg}
}

func (c *UserService) CreateUser(ctx context.Context, User *pb.User) (*pb.Void, error) {
	pb, err := c.stg.User().CreateUser(User)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *UserService) GetAllUser(ctx context.Context, pb *pb.User) (*pb.GetAllUsers, error) {
	Users, err := c.stg.User().GetAllUser(pb)
	if err != nil {
		log.Print(err)
	}

	return Users, err
}

func (c *UserService) GetByIdUser(ctx context.Context, id *pb.ById) (*pb.User, error) {
	prod, err := c.stg.User().GetByIdUser(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *UserService) UpdateUser(ctx context.Context, User *pb.User) (*pb.Void, error) {
	pb, err := c.stg.User().UpdateUser(User)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *UserService) DeleteUser(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.User().DeleteUser(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *UserService) LoginUser(ctx context.Context, username *pb.User) (*pb.User, error) {
	prod, err := c.stg.User().LoginUser(username)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}
