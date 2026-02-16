package grpcclient

import (
	pb "clirzy/auth/proto"
	"context"

	"google.golang.org/grpc"
)

type AuthServiceClient struct {
	client pb.AuthServiceClient
}

func NewAuthServiceClient(conn *grpc.ClientConn) *AuthServiceClient {
	return &AuthServiceClient{
		client: pb.NewAuthServiceClient(conn),
	}
}

func (c *AuthServiceClient) ValidateAccessToken(ctx context.Context, accessToken string) (userID string, roles []string, err error) {
	resp, err := c.client.ValidateAccessToken(ctx, &pb.ValidateAccessTokenRequest{AccessToken: accessToken})
	if err != nil {
		return "", nil, err
	}

	return resp.UserId, resp.Roles, nil
}
