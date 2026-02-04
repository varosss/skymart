package payload

type OrderCreatedItem struct {
	ProductID string `json:"product_id"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	Qty       int64  `json:"qty"`
}

type OrderCreatedPayload struct {
	OrderID  string             `json:"order_id"`
	BuyerID  string             `json:"buyer_id"`
	Total    int64              `json:"total"`
	Currency string             `json:"currency"`
	Items    []OrderCreatedItem `json:"items"`
}
