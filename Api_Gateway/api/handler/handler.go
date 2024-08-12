package handler

import (
	u "api/genprotos/auth"

	"google.golang.org/grpc"
)

type UserService struct {
	Auth u.UserServiceClient
}

func NewUserService(authConn *grpc.ClientConn) *UserService {
	return &UserService{
		Auth: u.NewUserServiceClient(authConn),
	}
}
