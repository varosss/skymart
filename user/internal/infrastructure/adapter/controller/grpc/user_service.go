package grpc

import (
	aport "clirzy/user/internal/application/port"
	"clirzy/user/internal/application/usecase"
	"clirzy/user/internal/domain"
	pb "clirzy/user/proto"
	"context"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer

	users          aport.UserQuery
	registerUserUC *usecase.RegisterUserUseCase
}

func NewUserServiceServer(registerUserUC *usecase.RegisterUserUseCase) *UserServiceServer {
	return &UserServiceServer{
		registerUserUC: registerUserUC,
	}
}

func (s *UserServiceServer) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.User, error) {
	user, err := s.users.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	return &pb.User{
		Id:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}

func (s *UserServiceServer) GetUserByID(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error) {
	user, err := s.users.GetByID(ctx, req.Id)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	return &pb.User{
		Id:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}

func (s *UserServiceServer) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user, err := s.registerUserUC.Execute(ctx, usecase.RegisterUserCommand{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &pb.RegisterUserResponse{
		UserId: user.ID().String(),
	}, nil
}
