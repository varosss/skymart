package grpc

import (
	aport "clirzy/product/internal/application/port"
	"clirzy/product/internal/domain/valueobject"
	pb "clirzy/product/proto"
	"context"
)

type ProductServiceServer struct {
	pb.UnimplementedProductServiceServer

	products aport.ProductQuery
}

func NewProductServiceServer(products aport.ProductQuery) *ProductServiceServer {
	return &ProductServiceServer{
		products: products,
	}
}

func (s *ProductServiceServer) GetProductsByIDs(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	productIDs := make([]valueobject.ProductID, len(req.Ids))
	for i, id := range req.Ids {
		parsedProductID, err := valueobject.ParseProductID(id)
		if err != nil {
			return nil, err
		}

		productIDs[i] = parsedProductID
	}

	products, err := s.products.GetProducts(ctx, productIDs)
	if err != nil {
		return nil, err
	}

	respProducts := make([]*pb.Product, len(products))

	for i, product := range products {
		respProducts[i] = &pb.Product{
			Id:       product.ID,
			SellerId: product.SellerID,
			Price:    product.Price,
			Currency: product.Currency,
		}
	}

	return &pb.GetProductsResponse{
		Products: respProducts,
	}, nil
}
