package grpcclient

import (
	pb "clirzy/billing/proto"
	aport "clirzy/payment/internal/application/port"
	"clirzy/payment/internal/domain/valueobject"
	"context"

	"google.golang.org/grpc"
)

type BillingServiceClient struct {
	client pb.BillingServiceClient
}

func NewBillingServiceClient(conn *grpc.ClientConn) *BillingServiceClient {
	return &BillingServiceClient{
		client: pb.NewBillingServiceClient(conn),
	}
}

func (c *BillingServiceClient) GetInvoiceByID(ctx context.Context, invoiceID valueobject.InvoiceID) (*aport.InvoiceDTO, error) {
	invoice, err := c.client.GetInvoiceById(ctx, &pb.GetInvoiceByIdRequest{Id: invoiceID.String()})
	if err != nil {
		return nil, err
	}

	return &aport.InvoiceDTO{
		ID:       invoice.Id,
		Amount:   invoice.Amount,
		Currency: invoice.Currency,
		Status:   invoice.Status,
	}, nil
}
