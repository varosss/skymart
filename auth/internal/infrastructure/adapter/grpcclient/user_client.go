package grpcclient

import (
	"clirzy/auth/internal/application/port"
	"clirzy/auth/internal/domain/valueobject"
	pb "clirzy/user/proto"
	"context"

	"google.golang.org/grpc"
)

type UserServiceClient struct {
	client pb.UserServiceClient
}

func NewUserServiceClient(conn *grpc.ClientConn) *UserServiceClient {
	return &UserServiceClient{
		client: pb.NewUserServiceClient(conn),
	}
}

func (c *UserServiceClient) FindByEmail(ctx context.Context, email string) (*port.UserDTO, error) {
	user, err := c.client.GetUserByEmail(ctx, &pb.GetUserByEmailRequest{Email: email})
	if err != nil {
		return nil, err
	}

	return &port.UserDTO{
		ID:           user.Id,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}

func (c *UserServiceClient) GetByID(ctx context.Context, userID valueobject.UserID) (*port.UserDTO, error) {
	user, err := c.client.GetUserByID(ctx, &pb.GetUserByIdRequest{Id: userID.String()})
	if err != nil {
		return nil, err
	}

	return &port.UserDTO{
		ID:           user.Id,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}

func (c *UserServiceClient) RegisterUser(ctx context.Context, email string, password string) (*port.UserDTO, error) {
	resp, err := c.client.RegisterUser(ctx, &pb.RegisterUserRequest{Email: email, Password: password})
	if err != nil {
		return nil, err
	}

	return &port.UserDTO{
		ID: resp.UserId,
	}, nil
}
