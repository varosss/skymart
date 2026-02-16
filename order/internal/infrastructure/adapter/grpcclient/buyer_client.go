package grpcclient

import (
	pb "clirzy/buyer/proto"
	aport "clirzy/order/internal/application/port"
	"context"

	"google.golang.org/grpc"
)

type BuyerServiceClient struct {
	client pb.BuyerServiceClient
}

func NewBuyerServiceClient(conn *grpc.ClientConn) *BuyerServiceClient {
	return &BuyerServiceClient{
		client: pb.NewBuyerServiceClient(conn),
	}
}

func (c *BuyerServiceClient) GetByID(ctx context.Context, buyerID string) (*aport.BuyerDTO, error) {
	buyer, err := c.client.GetBuyerByID(ctx, &pb.GetBuyerByIdRequest{Id: buyerID})
	if err != nil {
		return nil, err
	}

	return &aport.BuyerDTO{
		ID:       buyer.Id,
		IsActive: buyer.IsActive,
	}, nil
}
