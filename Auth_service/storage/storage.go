package storage

import pb "auth/genprotos"

type StorageI interface {
	User() UserI
}

type UserI interface {
	RegisterUser(user *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error)
	LoginUser(user *pb.LoginUserRequest) (*pb.LoginUserResponse, error)
	GetByIdUser(id *pb.GetByIdUserRequest) (*pb.GetByIdUserResponse, error)
	UpdateUser(req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUser(id *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	ChangePassword(password *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error)
	ForgotPassword(forgotPass *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error)
	GetUserByEmail(email string) (*pb.UpdateUserResponse, error)
	ResetPassword(resetPass *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error)
}