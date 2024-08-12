package service

import (
	"context"
	"errors"
	"log"

	pb "auth/genprotos"
	s "auth/storage"
)

type UserService struct {
	stg s.StorageI
	pb.UnimplementedUserServiceServer
}

func NewUserService(stg s.StorageI) *UserService {
	return &UserService{stg: stg}
}

var ErrUserAlreadyExists = errors.New("user already exists")

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	existingUser, err := s.stg.User().GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}
	_, err = s.stg.User().RegisterUser(req)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterUserResponse{}, nil
}

func (c *UserService) LoginUser(ctx context.Context, login *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := c.stg.User().LoginUser(login)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return user, nil
}

func (c *UserService) GetByIdUser(ctx context.Context, id *pb.GetByIdUserRequest) (*pb.GetByIdUserResponse, error) {
	user, err := c.stg.User().GetByIdUser(id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return user, nil
}

func (c *UserService) UpdateUser(ctx context.Context, user *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	res, err := c.stg.User().UpdateUser(user)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return res, nil
}

func (c *UserService) DeleteUser(ctx context.Context, id *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	void, err := c.stg.User().DeleteUser(id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return void, nil
}

func (c *UserService) ChangePassword(ctx context.Context, changePass *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	void, err := c.stg.User().ChangePassword(changePass)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return void, nil
}

func (c *UserService) ForgotPassword(ctx context.Context, forgotPass *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	void, err := c.stg.User().ForgotPassword(forgotPass)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return void, nil
}

func (c *UserService) ResetPassword(ctx context.Context,resetPass *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error){
	void, err := c.stg.User().ResetPassword(resetPass)
	if err != nil{
		log.Println(err)
		return nil, err
	}
	return void, err
}