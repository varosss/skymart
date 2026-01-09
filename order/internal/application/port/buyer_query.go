package port

type BuyerDTO struct {
	ID       string
	IsActive bool
}

type BuyerQuery interface {
	GetByID(buyerID string) (BuyerDTO, error)
}
