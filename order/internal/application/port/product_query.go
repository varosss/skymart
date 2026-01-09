package port

type ProductDTO struct {
	ID          string
	SellerID    string
	Amount      int64
	Currency    string
	IsPublished bool
}

type ProductQuery interface {
	GetProducts(productIDs []string) ([]ProductDTO, error)
}
