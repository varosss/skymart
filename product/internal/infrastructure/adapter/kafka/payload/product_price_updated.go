package payload

type ProductPriceUpdatedPayload struct {
	ProductID string `json:"product_id"`
	SellerID  string `json:"seller_id"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
}
