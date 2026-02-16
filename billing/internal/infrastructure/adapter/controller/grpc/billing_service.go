package grpc

import (
	aport "clirzy/billing/internal/application/port"
	"clirzy/billing/internal/domain/valueobject"
	pb "clirzy/billing/proto"
	"context"
)

type BillingServiceServer struct {
	pb.UnimplementedBillingServiceServer

	invoices aport.InvoiceQuery
}

func NewBillingServiceServer(invoices aport.InvoiceQuery) *BillingServiceServer {
	return &BillingServiceServer{
		invoices: invoices,
	}
}

func (s *BillingServiceServer) GetInvoiceById(ctx context.Context, req *pb.GetInvoiceByIdRequest) (*pb.Invoice, error) {
	invoiceID, err := valueobject.ParseInvoiceID(req.Id)
	if err != nil {
		return nil, err
	}

	invoice, err := s.invoices.GetByID(ctx, invoiceID)
	if err != nil {
		return nil, err
	}

	return &pb.Invoice{
		Id:       invoice.ID,
		Amount:   invoice.Amount,
		Currency: invoice.Currency,
		Status:   invoice.Status,
	}, nil
}
