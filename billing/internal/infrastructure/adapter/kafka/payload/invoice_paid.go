package payload

type InvoicePaidPayload struct {
	InvoiceID string `json:"invoice_id"`
	BuyerID   string `json:"buyer_id"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
}
