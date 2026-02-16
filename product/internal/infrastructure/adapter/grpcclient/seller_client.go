package grpcclient

import (
	aport "clirzy/product/internal/application/port"
	"clirzy/product/internal/domain/valueobject"
	pb "clirzy/seller/proto"
	"context"

	"google.golang.org/grpc"
)

type SellerServiceClient struct {
	client pb.SellerServiceClient
}

func NewSellerServiceClient(conn *grpc.ClientConn) *SellerServiceClient {
	return &SellerServiceClient{
		client: pb.NewSellerServiceClient(conn),
	}
}

func (c *SellerServiceClient) GetByID(ctx context.Context, sellerID valueobject.SellerID) (*aport.SellerDTO, error) {
	seller, err := c.client.GetSellerByID(ctx, &pb.GetSellerByIdRequest{Id: sellerID.String()})
	if err != nil {
		return nil, err
	}

	parsedSellerID, err := valueobject.ParseSellerID(seller.Id)
	if err != nil {
		return nil, err
	}

	return &aport.SellerDTO{
		ID:       parsedSellerID,
		IsActive: seller.IsActive,
	}, nil
}

func (c *SellerServiceClient) GetByUserID(ctx context.Context, userID valueobject.UserID) (*aport.SellerDTO, error) {
	seller, err := c.client.GetSellerByUserID(ctx, &pb.GetSellerByUserIdRequest{UserId: userID.String()})
	if err != nil {
		return nil, err
	}

	parsedSellerID, err := valueobject.ParseSellerID(seller.Id)
	if err != nil {
		return nil, err
	}

	return &aport.SellerDTO{
		ID:       parsedSellerID,
		IsActive: seller.IsActive,
	}, nil
}
