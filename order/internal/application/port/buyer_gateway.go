package port

type BuyerDTO struct {
	ID       string
	IsActive bool
}

type BuyerGateway interface {
	GetByID(buyerID string) (BuyerDTO, error)
}
