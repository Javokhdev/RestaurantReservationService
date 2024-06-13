package handler

import pb "github.com/dilshodforever/restaurant-auth/service"

type Handler struct {
	User   *pb.UserService
}

func NewHandler(us *pb.UserService) *Handler {
	return &Handler{us}
}
