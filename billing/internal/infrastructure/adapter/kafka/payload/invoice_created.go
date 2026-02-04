package payload

type InvoiceCreatedItem struct {
	ProductID string `json:"product_id"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	Qty       int64  `json:"qty"`
}

type InvoiceCreatedPayload struct {
	InvoiceID string               `json:"invoice_id"`
	BuyerID   string               `json:"buyer_id"`
	Total     int64                `json:"total"`
	Currency  string               `json:"currency"`
	Items     []InvoiceCreatedItem `json:"items"`
}
