package grpc

import (
	aport "clirzy/seller/internal/application/port"
	"clirzy/seller/internal/domain/valueobject"
	pb "clirzy/seller/proto"
	"context"
)

type SellerServiceServer struct {
	pb.UnimplementedSellerServiceServer

	sellers aport.SellerQuery
}

func NewSellerServiceServer(buyers aport.SellerQuery) *SellerServiceServer {
	return &SellerServiceServer{
		sellers: buyers,
	}
}

func (s *SellerServiceServer) GetSellerByID(ctx context.Context, req *pb.GetSellerByIdRequest) (*pb.Seller, error) {
	sellerID, err := valueobject.ParseSellerID(req.Id)
	if err != nil {
		return nil, err
	}

	seller, err := s.sellers.GetByID(ctx, sellerID)
	if err != nil {
		return nil, err
	}

	return &pb.Seller{
		Id:       seller.ID,
		UserId:   seller.UserID,
		IsActive: seller.IsActive,
	}, nil
}

func (s *SellerServiceServer) GetSellerByUserID(ctx context.Context, req *pb.GetSellerByUserIdRequest) (*pb.Seller, error) {
	userID, err := valueobject.ParseUserID(req.UserId)
	if err != nil {
		return nil, err
	}

	seller, err := s.sellers.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &pb.Seller{
		Id:       seller.ID,
		UserId:   seller.UserID,
		IsActive: seller.IsActive,
	}, nil
}
