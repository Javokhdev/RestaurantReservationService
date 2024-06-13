package postgres

import (
	pb "github.com/dilshodforever/restaurant-auth/genprotos"
)

type InitRoot interface {
	User() User
}
type User interface {
	CreateUser(user *pb.User) (*pb.Void, error)
	GetByIdUser(id *pb.ById) (*pb.User, error)
	GetAllUser(_ *pb.User) (*pb.GetAllUsers, error)
	UpdateUser(user *pb.User) (*pb.Void, error)
	DeleteUser(id *pb.ById) (*pb.Void, error)
	LoginUser(user *pb.User ) (*pb.User, error)
}
