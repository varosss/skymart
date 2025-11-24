package userclient

import (
	"context"

	userpb "clirzy/services/user/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *grpc.ClientConn
	remote userpb.UserServiceClient
}

func New(address string) (*Client, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		remote: userpb.NewUserServiceClient(conn),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) CreateUser(ctx context.Context, username, password string) (*userpb.CreateUsersResponse, error) {
	return c.remote.CreateUsers(ctx, &userpb.CreateUsersRequest{
		Users: []*userpb.CreateUser{{Username: username, Password: password}},
	})
}

func (c *Client) GetUserByUsername(ctx context.Context, username string) (*userpb.UserResponse, error) {
	return c.remote.GetUserByUsername(ctx, &userpb.GetUserByUsernameRequest{
		Username: username,
	})
}

func (c *Client) GetUserById(ctx context.Context, id uint) (*userpb.UserResponse, error) {
	return c.remote.GetUserById(ctx, &userpb.GetUserByIdRequest{
		Id: int64(id),
	})
}
