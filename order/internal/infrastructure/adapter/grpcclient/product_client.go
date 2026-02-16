package grpcclient

import (
	aport "clirzy/order/internal/application/port"
	"clirzy/order/internal/domain/valueobject"
	pb "clirzy/product/proto"
	"context"

	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	client pb.ProductServiceClient
}

func NewProductServiceClient(conn *grpc.ClientConn) *ProductServiceClient {
	return &ProductServiceClient{
		client: pb.NewProductServiceClient(conn),
	}
}

func (c *ProductServiceClient) GetProducts(ctx context.Context, productIDs []string) ([]*aport.ProductDTO, error) {
	resp, err := c.client.GetProducts(ctx, &pb.GetProductsRequest{Ids: productIDs})
	if err != nil {
		return nil, err
	}

	products := make([]*aport.ProductDTO, len(resp.Products))
	for i, product := range resp.Products {
		productID, err := valueobject.ParseProductID(product.Id)
		if err != nil {
			return nil, err
		}

		sellerID, err := valueobject.ParseSellerID(product.SellerId)
		if err != nil {
			return nil, err
		}

		price, err := valueobject.NewMoney(product.Price, product.Currency)
		if err != nil {
			return nil, err
		}

		products[i] = &aport.ProductDTO{
			ID:          productID,
			SellerID:    sellerID,
			Price:       price,
			IsPublished: product.IsPublished,
		}
	}

	return products, nil
}
