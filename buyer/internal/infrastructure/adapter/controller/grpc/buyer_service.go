package grpc

import (
	aport "clirzy/buyer/internal/application/port"
	"clirzy/buyer/internal/domain/valueobject"
	pb "clirzy/buyer/proto"
	"context"
)

type BuyerServiceServer struct {
	pb.UnimplementedBuyerServiceServer

	buyers aport.BuyerQuery
}

func NewBuyerServiceServer(buyers aport.BuyerQuery) *BuyerServiceServer {
	return &BuyerServiceServer{
		buyers: buyers,
	}
}

func (s *BuyerServiceServer) GetBuyerByID(ctx context.Context, req *pb.GetBuyerByIdRequest) (*pb.Buyer, error) {
	buyerID, err := valueobject.ParseBuyerID(req.Id)
	if err != nil {
		return nil, err
	}

	buyer, err := s.buyers.GetByID(ctx, buyerID)
	if err != nil {
		return nil, err
	}

	return &pb.Buyer{
		Id:       buyer.ID,
		IsActive: buyer.IsActive,
	}, nil
}
